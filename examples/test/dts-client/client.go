package dts

import (
	"encoding/binary"
	"encoding/json"
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk/model"
	"github.com/golang/protobuf/proto"
	"io"
	"log"
	"net"
	"time"
)

type Client struct {
}

func (c Client) Run() {
	conn, err := net.DialTimeout("tcp", "192.168.0.215:17083", time.Second)
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		n, err := conn.Write([]byte{0, 0, 0, 0, 4})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n)
		// [0 0 0 4 4 24 1 32 1]

		//[]byte{0, 0, 0, 8, 4, 8, 1, 16, 1, 24, 1, 32, 1}
		//[]byte{0, 0, 0, 4, 4, 24, 1, 32, 1}

		n, err = conn.Write([]byte{0, 0, 0, 2, 4, 32, 1})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n)
	}()

	go func(conn net.Conn) {
		for {
			n, err := conn.Write([]byte{0, 0, 0, 0, 250})
			if err != nil {
				log.Fatal(err)
			}
			log.Println("heart", n)
			time.Sleep(time.Second * 10)
		}
	}(conn)
	for {
		sizeData := make([]byte, 4)
		n, err := io.ReadFull(conn, sizeData)
		if err != nil {
			log.Fatal(err)
		}
		size := binary.BigEndian.Uint32(sizeData)
		log.Println("size", size)
		cmdData := make([]byte, 1)
		n, err = io.ReadFull(conn, cmdData)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("cmd", cmdData[0])
		data := make([]byte, size)
		n, err = io.ReadFull(conn, data)
		if err != nil {
			log.Fatal(err)
		}
		switch model.MsgID(cmdData[0]) {
		case model.MsgID_ZoneTempNotifyID:
			reply := model.ZoneTempNotify{}
			err := proto.Unmarshal(data, &reply)
			data, _ = json.Marshal(reply.Zones)
			log.Println(n, len(data), len(reply.Zones), err)
		default:
			log.Println(cmdData[0])
		}
	}
}
