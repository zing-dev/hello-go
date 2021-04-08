package bytes

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

const (
	One byte = iota + 1
	List
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Package struct {
	locker sync.Mutex
	Cache  bytes.Buffer
	Cmd    byte
	Size   int
	Data   []byte
}

func NewPackage(cmd byte, data interface{}) *Package {
	d := make([]byte, 0)
	if data != nil {
		d, _ = json.Marshal(data)
	}
	return &Package{
		Cmd:  cmd,
		Size: len(d),
		Data: d,
	}
}

func (p *Package) Pack() []byte {
	b := bytes.Buffer{}
	size := make([]byte, 4)
	binary.BigEndian.PutUint32(size, uint32(p.Size))
	b.Write(size)
	b.WriteByte(p.Cmd)
	b.Write(p.Data)
	return b.Bytes()
}

//0-3 size
//4 cmd
//5: data
func (p *Package) resolve(data []byte) error {
	if len(data) < 5 {
		return errors.New("data's length so small")
	}
	p.Size = int(binary.BigEndian.Uint32(data[:4]))
	p.Cmd = data[4]
	return nil
}

func (p *Package) read(data []byte) error {
	if len(data) < 5 {
		return errors.New("data's length so small")
	}
	p.Size = int(binary.BigEndian.Uint32(data[:4]))
	p.Cmd = data[4]
	p.Data = data[5:]
	if p.Size != len(p.Data) {
		return errors.New("size not equal data's length")
	}
	return nil
}

func (p *Package) WriteTo(w io.Writer) (int64, error) {
	n, err := w.Write(p.Pack())
	log.Println("WriteTo: ", p.Cmd, p.Size, n)
	return int64(n), err
}

func (p *Package) ReadFrom(r io.Reader) (int64, error) {
	header := make([]byte, 5)
	_, err := io.ReadFull(r, header)
	if err != nil {
		return 0, err
	}
	p.Size = int(binary.BigEndian.Uint32(header[:4]))
	p.Cmd = header[4]
	data := make([]byte, p.Size)
	n, err := io.ReadFull(r, data)
	p.Data = data
	return int64(n), err
}

func (p *Package) Unpack(data []byte) error {
	p.locker.Lock()
	defer p.locker.Unlock()
	log.Println("Unpack start")
	if len(p.Data) == 0 {
		if len(data) < 5 {
			return errors.New("data's length so small")
		}
		p.Size = int(binary.BigEndian.Uint32(data[:4]))
		p.Cmd = data[4]
		p.Data = data[5:]
	}
	if len(p.Data) == p.Size {
		//p.Cache.Reset()
		return nil
	} else {
		p.Data = append(p.Data, data...)
	}
	//p.Cache.Write(data)
	//if err := p.read(p.Cache.Bytes()); err != nil {
	//	return err
	//}
	//log.Println("Unpack end")
	//p.Cache.Reset()
	return errors.New("errr")
}

type Client struct {
	conn net.Conn
}

func (c *Client) Write(data []byte) {
	_, _ = c.conn.Write(data)
}

func (c *Client) Read() []byte {
	data, err := io.ReadAll(c.conn)
	if err != nil && err != io.EOF {
		panic(err)
	}
	return data
}

func (c *Client) Run() {
	conn, err := net.DialTimeout("tcp", ":1122", time.Second)
	if err != nil {
		panic(err)
	}
	err = conn.(*net.TCPConn).SetWriteBuffer(1024)
	if err != nil {
		panic(err)
	}
	c.conn = conn
	log.Println(conn.LocalAddr())
}

type Server struct {
}

func (s *Server) Run() {
	listen, err := net.Listen("tcp", ":1122")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			log.Println("new conn from ", conn.RemoteAddr().String())
			p := Package{}
			for {
				_, err := p.ReadFrom(conn)
				if err != nil {
					log.Println("ReadFrom: ", err)
				} else {
					switch p.Cmd {
					case One:
						log.Println("server cmd: One")
						p := NewPackage(p.Cmd, User{Id: 1, Name: "zing"})
						_, err := p.WriteTo(conn)
						if err != nil {
							log.Println("One WriteTo: ", err)
						}
					case List:
						users := make([]User, 100000)
						for i := 0; i < 100000; i++ {
							users[i] = User{Id: i + 1, Name: fmt.Sprintf("name-%d", i+1)}
						}
						p := NewPackage(List, users)
						_, err := p.WriteTo(conn)
						if err != nil {
							log.Println("List WriteTo: ", err)
						}
					}
				}
			}
		}(conn)
	}
}
