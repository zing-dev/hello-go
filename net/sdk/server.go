package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

var server *SdkServer
var once sync.Once
var group sync.WaitGroup

type MsgID int32

const (
	ConnectID    MsgID = iota //0
	DisconnectID              //1
	Open                      //2
	Close                     //3
	Reset                     //4
	HeartBeatID  MsgID = 250
)

type Message struct {
	Success bool
	Err     string
	Data    interface{}
}

type SdkServer struct {
	listen             *net.TCPListener
	handle             func(byte, int32, []byte, net.Conn)
	sessId             int32            //自增id
	heartBeat          *time.Ticker     //心跳包的发送
	heartBeatCloseFlag chan interface{} //关闭心跳
	clients            *sync.Map        //连接的客户端
}

type SdkClient struct {
	ID            int
	Conn          *net.TCPConn
	LastHeartBeat int64 //最后一次心跳
}

//单例实例化sdk server
func NewSdkServer() *SdkServer {
	once.Do(func() {
		server = &SdkServer{}
		server.init()
	})
	return server
}

func (ss *SdkServer) init() {
	ss.clients = new(sync.Map)
	//addr, err := net.ResolveTCPAddr("tcp4", ":17082")
	addr, err := net.ResolveTCPAddr("tcp4", ":9000")
	if err != nil {
		log.Fatal("ResolveTCPAddr!", err)
	}
	//监听端口
	listen, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal("监听端口失败!", err)
	}
	log.Println("服务器启动成功===>>")
	ss.listen = listen
	ss.heartBeat = time.NewTicker(time.Second * 5)
	ss.SetHandle()
	group.Add(1)
	go ss.accept()
	group.Add(1)
	go ss.timeoutHandle()
}

//心跳
func (ss *SdkServer) timeoutHandle() {
	for {
		select {
		case <-ss.heartBeat.C:
			ss.clients.Range(func(key, value interface{}) bool {
				if time.Now().UnixNano()/1000000-value.(*SdkClient).LastHeartBeat > 10000 {
					v, ok := ss.clients.Load(key.(int32))
					if ok {
						_ = v.(*SdkClient).Conn.Close()
					}
				}
				return true
			})

		case <-ss.heartBeatCloseFlag:
			group.Done()
			return
		}
	}
}

func (ss *SdkServer) accept() {
	for {
		conn, err := ss.listen.AcceptTCP()
		if err != nil {
			log.Println("监听连接失败!", err)
			continue
		}
		client := &(SdkClient{Conn: conn, LastHeartBeat: time.Now().UnixNano() / 1000000})
		ss.clients.Store(ss.sessId, client)
		_ = conn.SetWriteBuffer(5000)
		_ = conn.SetReadBuffer(5000)
		go ss.handles(ss.sessId, conn)
		ss.sessId++
	}
}

func (ss *SdkServer) tcpHandle(cmdId byte, sessId int32, data []byte, conn net.Conn) {
	switch MsgID(cmdId) {
	case HeartBeatID:
		v, ok := ss.clients.Load(sessId)
		if ok {
			c := v.(*SdkClient)
			c.LastHeartBeat = time.Now().UnixNano() / 1000000
			ss.clients.Store(sessId, v)
		}

	}
	if ss.handle != nil {
		ss.handle(cmdId, sessId, data, conn)
	}

}

func (ss *SdkServer) handles(sessId int32, conn net.Conn) {
	defer func(sessId int32) {
		ss.tcpHandle(byte(DisconnectID), sessId, nil, nil)
		_ = conn.Close()
		ss.clients.Delete(sessId)
	}(sessId)

	//todo
	//ss.tcpHandle(byte(ConnectID), sessId, nil, nil)
	buf := make([]byte, 1024)
	var cache bytes.Buffer
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break
		}

		cache.Write(buf[:n])
		for {
			if ss.unpack(sessId, &cache, conn) {
				break
			}
			//if !unpack(&cache) {
			//	ss.handle(cmd, sessId, buf[:pkgSize+5], conn)
			//}
		}

	}
}

// true 处理完成 false 循环继续处理
func (ss *SdkServer) unpack(sessId int32, cache *bytes.Buffer, conn net.Conn) bool {
	if cache.Len() < 5 {
		return true
	}
	buf := cache.Bytes()
	pkgSize := byteToInt(buf[:4])
	//长度不够
	if pkgSize > len(buf)-5 {
		return true
	}
	cmd := buf[4]
	fmt.Println(len(buf), pkgSize+5)
	ss.handle(cmd, sessId, buf[:pkgSize+5], conn)
	cache.Reset()
	cache.Write(buf[5+pkgSize:])
	return false
}

func (sc *SdkClient) handle(cmdId byte, sessId int32, data []byte, conn net.Conn) {
	log.Println(cmdId, sessId, string(data))
	switch MsgID(cmdId) {
	case ConnectID:
		log.Println("client Connect")
	case DisconnectID:
		log.Println("session Disconnect: ", sessId)
	case Open:
		log.Println("session Open: ", sessId)
		send(cmdId, conn)
	case Close:
		log.Println("session Close: ", sessId)
		send(cmdId, conn)
	case Reset:
		log.Println("session Reset: ", sessId)
		send(cmdId, conn)
	}
}

func unpack(cache *bytes.Buffer) bool {
	if cache.Len() < 5 {
		return true
	}
	buf := cache.Bytes()
	pkgSize := byteToInt(buf[:4])
	if pkgSize > len(buf)-5 {
		return true
	}
	cmd := buf[4]
	//fmt.Println(len(buf), pkgSize+5)
	//cmd ,buf[:pkgSize+5]
	log.Println("==> ", cmd, string(buf[:pkgSize+5]), buf[:pkgSize+5])
	//log.Println("==> ", cmd, string(buf[4:pkgSize+4]), buf[:pkgSize+5])
	data := map[string]interface{}{}
	_ = json.Unmarshal(buf[:pkgSize+5], &data)
	log.Println("data >>> ", data)
	cache.Reset()
	cache.Write(buf[5+pkgSize:])
	return false
}

func (ss *SdkServer) SetHandle() {
	//tcpHandle func(byte, int32, []byte, net.Conn)
	//ss.handle = tcpHandle
	ss.handle = func(cmdId byte, sessId int32, data []byte, conn net.Conn) {
		log.Println(cmdId, sessId, string(data))
		switch MsgID(cmdId) {
		case ConnectID:
			log.Println("client Connect")
		case DisconnectID:
			log.Println("session Disconnect: ", sessId)
		case Open:
			log.Println("session Open: ", sessId)
			send(cmdId, conn)
		case Close:
			log.Println("session Close: ", sessId)
			send(cmdId, conn)
		case Reset:
			log.Println("session Reset: ", sessId)
			send(cmdId, conn)
		}
	}
}
func send(cmdId byte, conn net.Conn) {
	c := Message{
		Success: true,
		Err:     "OK ==> " + strconv.Itoa(int(cmdId)),
	}
	b, _ := pack(&c, cmdId)
	_, _ = conn.Write(b)
}
func (ss *SdkServer) SendAll(msgObj interface{}) error {
	b, err := encode(msgObj)
	if err != nil {
		log.Println(err)
		return err
	}
	ss.clients.Range(func(key, value interface{}) bool {
		_, _ = value.(*SdkClient).Conn.Write(b)
		return true
	})
	return nil
}

func (ss *SdkServer) Close() {
	ss.clients.Range(func(key, value interface{}) bool {
		_ = value.(*SdkClient).Conn.Close()
		return true
	})
}

func byteToInt(b []byte) int {
	mask := 0xff
	temp := 0
	n := 0
	for i := 0; i < len(b); i++ {
		n <<= 8
		temp = int(b[i]) & mask
		n |= temp
	}
	return n
}

//整形转换成字节
func intToBytes(n int64, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 5, 6, 7, 8:
		tmp := int64(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		_ = binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

func encode(msgObj interface{}) ([]byte, error) {
	data, err := json.Marshal(msgObj)
	if err != nil {
		return nil, err
	}
	cache := make([]byte, len(data)+5)
	length, _ := intToBytes(int64(len(data)), 4)
	copy(cache, length)
	switch msgObj.(type) {
	case *Message:
		cache[4] = byte(ConnectID)
	}
	copy(cache[5:], data)
	return cache, err
}

func pack(msgObj interface{}, cmdId byte) ([]byte, error) {
	data, err := json.Marshal(msgObj)
	if err != nil {
		return nil, err
	}
	cache := make([]byte, len(data)+5)
	length, _ := intToBytes(int64(len(data)), 4)
	copy(cache, length)
	cache[4] = byte(cmdId)

	switch msgObj.(type) {
	case *Message:
		cache[4] = byte(cmdId)
	}
	copy(cache[5:], data)
	return cache, err
}
