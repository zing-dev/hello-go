package net

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

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
		panic(fmt.Sprint("DialTimeout: ", err))
	}
	err = conn.(*net.TCPConn).SetWriteBuffer(1024)
	if err != nil {
		panic(fmt.Sprint("SetWriteBuffer: ", err))
	}
	c.conn = conn
	log.Println(conn.LocalAddr())

	go func() {
		log.Println("client write")
		p := NewPackage(List, nil)
		n, err := p.WriteTo(c.conn)
		if err != nil {
			log.Println(n, err)
		}
	}()
	log.Println("client read")
	p := &Package{}
	for {
		n, err := p.ReadFrom(conn)
		if err != nil {
			log.Println(n, err)
		}
		switch p.Cmd {
		case One:
			user := User{}
			err = json.Unmarshal(p.Data, &user)
			log.Println(user)
		case List:
			var users []User
			err = json.Unmarshal(p.Data, &users)
			log.Println(len(users))
		}
	}
}
