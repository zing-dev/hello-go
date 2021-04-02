package bytes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

var buffer = bytes.NewBufferString("hello world")

func TestNewBufferString(t *testing.T) {
	buf := bytes.NewBufferString("hello world")
	log.Println(buf.Bytes())
	log.Println(buf.Cap())
	log.Println(buf.Len())
	log.Println(buf.String())
}

func TestNewBuffer(t *testing.T) {
	buf := bytes.NewBuffer([]byte("hello world"))
	log.Println(buf.String())
}

func TestTruncate(t *testing.T) {
	buffer.Truncate(2)
	log.Println(buffer.String())
}

func TestGrow(t *testing.T) {
	log.Println(buffer.Len())
	log.Println(buffer.Cap())
	buffer.Grow(buffer.Cap())
	log.Println(buffer.Len())
	log.Println(buffer.Cap())
}

func TestWrite(t *testing.T) {
	buffer.Reset()
	buffer.Write([]byte("hello golang"))
	log.Println(string(buffer.Bytes()))
}

func TestWriteString(t *testing.T) {
	buffer.Reset()
	buffer.WriteString("golang")
	log.Println(string(buffer.Bytes()))
}

func TestReadFrom(t *testing.T) {
	buf := bytes.NewBufferString("hello golang")
	buffer.Reset()
	n, err := buffer.ReadFrom(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buffer.String())
}

func TestWriteTo(t *testing.T) {
	buf := bytes.Buffer{}
	n, err := buffer.WriteTo(&buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buf.String())
}

func TestWriteByte(t *testing.T) {
	buffer.Reset()
	buffer.WriteByte('a')
	buffer.WriteByte('b')
	buffer.WriteByte('c')
	log.Println(buffer.String())
}

func TestWriteRune(t *testing.T) {
	buffer.Reset()
	buffer.WriteRune('哈')
	buffer.WriteRune('喽')
	log.Println(buffer.String())
}

func TestRead(t *testing.T) {
	buf := make([]byte, buffer.Len())
	n, err := buffer.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	log.Println(buf)
	log.Println(buffer.Len())
}

func TestNext(t *testing.T) {
	log.Println(buffer.Next(1))
	log.Println(buffer.Next(2))
}

func TestReadByte(t *testing.T) {
	b, err := buffer.ReadByte()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(b), b)
}

func TestReadRune(t *testing.T) {
	r, size, err := buffer.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r, size)

	buffer.Reset()
	buffer.WriteRune('哈')
	r, size, err = buffer.ReadRune()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r, size)
}

func TestUnreadRune(t *testing.T) {
	log.Println(buffer.UnreadByte())
	buffer.Reset()
	log.Println(buffer.UnreadByte())
	buffer.Next(1)
	log.Println(buffer.UnreadByte())
	buffer.WriteByte('a')
	log.Println(buffer.UnreadByte())
	_, _ = buffer.ReadByte()
	log.Println(buffer.UnreadByte())
}

func TestReadBytes(t *testing.T) {
	line, err := buffer.ReadBytes(' ')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(line, string(line))
}

func TestReadString(t *testing.T) {
	line, err := buffer.ReadString(',')
	if err != nil {
		log.Fatal(err)
	}
	log.Println(line)
}

func TestPackageOne(t *testing.T) {
	user := User{Id: 1, Name: "zing"}
	p := NewPackage(One, user)
	err := p.Unpack(p.Pack())
	if err != nil {
		log.Println(err)
		return
	}
	u := new(User)
	_ = json.Unmarshal(p.Data, u)
	log.Println(u)
	log.Println(p.Cmd, p.Size, string(p.Data), p.Data)

	p = NewPackage(One, nil)
	err = p.Unpack(p.Pack())
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(p.Cmd, p.Size, string(p.Data), p.Data)
}

func TestPackageList(t *testing.T) {
	users := make([]User, 100)
	for i := 0; i < 100; i++ {
		users[i] = User{Id: i + 1, Name: fmt.Sprintf("name-%d", i+1)}
	}
	//users = make([]User, 0)
	p := NewPackage(List, users)
	err := p.Unpack(p.Pack())
	if err != nil {
		log.Println(err)
		return
	}
	u := new([]User)
	_ = json.Unmarshal(p.Data, u)
	log.Println(u)
}

func TestPackageListRead(t *testing.T) {
	users := make([]User, 100000)
	for i := 0; i < 100000; i++ {
		users[i] = User{Id: i + 1, Name: fmt.Sprintf("name-%d", i+1)}
	}
	p := NewPackage(List, users)
	r := bytes.NewReader(p.Pack())
	log.Println("size: ", r.Size())
	for {
		data := make([]byte, 1024)
		n, err := r.Read(data)
		if err != io.EOF /*&& n == 1024*/ {
			continue
		}
		err = p.Unpack(data[:n])
		var u []User
		_ = json.Unmarshal(p.Data, &u)
		log.Println("len", len(u))
		break
	}
}

func TestPackageClient(t *testing.T) {
	client := &Client{}
	client.Run()
	p := NewPackage(One, nil)
	go func(conn net.Conn) {
		n, err := p.ReadFrom(client.conn)
		if err != nil {
			log.Println(n, err)
		}
		user := User{}
		err = json.Unmarshal(p.Data, &user)
		log.Println(user)
	}(client.conn)
	n, err := p.WriteTo(client.conn)
	if err != nil {
		log.Println(n, err)
	}
	time.Sleep(time.Minute * 10)
}

func TestPackageServer(t *testing.T) {
	server := Server{}
	server.Run()
}
