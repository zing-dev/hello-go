package main

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"time"
)

type Client struct {
	conn net.Conn
}

type OpenMessageRequest struct {
	Relay []bool
}
type HeartBeat struct {
}

func NewClient() *Client {
	conn, err := net.Dial("tcp", ":17000")
	if err != nil {
		return nil
	}
	log.Println("success")
	return &Client{conn: conn}
}

func (c *Client) Heart() {
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			var data []byte
			d, _ := json.Marshal(HeartBeat{})
			data = append(data, IntToBytes(len(d))...)
			data = append(data, byte(250))
			data = append(data, d...)
			log.Println(data)
			write, _ := c.conn.Write(data)
			log.Println(write)
		}
	}
}

func (c *Client) Open() error {
	var data []byte
	bytes, _ := json.Marshal(OpenMessageRequest{Relay: []bool{
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
	}})
	data = append(data, IntToBytes(len(bytes))...)
	data = append(data, byte(2))
	data = append(data, bytes...)
	write, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	log.Println(write)
	return errors.New("open ok")
}

func (c *Client) Close() error {
	var data []byte
	bytes, _ := json.Marshal(OpenMessageRequest{Relay: []bool{
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
		true, true, true, true, true, true, true, true,
	}})
	data = append(data, IntToBytes(len(bytes))...)
	data = append(data, byte(3))
	data = append(data, bytes...)
	write, err := c.conn.Write(data)
	if err != nil {
		return err
	}
	log.Println(write)
	return errors.New("close ok")
}
