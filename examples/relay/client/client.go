package main

import (
	"errors"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

const (
	disconnected uint32 = iota
	connecting
	reconnecting
	connected
)

type Client struct {
	conn   net.Conn
	status uint32
	stop   chan struct{}
	sync.RWMutex
}

func NewClient() *Client {
	c := &Client{}
	c.setConnected(connecting)
	conn, err := net.Dial("tcp", "192.168.0.111:17000")
	if err != nil {
		return nil
	}
	c.setConnected(connected)
	c.stop = make(chan struct{})
	c.conn = conn
	log.Println("success")
	return c
}

func (c *Client) Heart() {
	t := time.NewTicker(time.Second * 3)
	defer t.Stop()
	for {
		select {
		case <-t.C:
			_, _ = c.conn.Write(NewProtocol(HeartBeatId).Pack())
		case <-c.stop:
			if c.stop != nil {
				close(c.stop)
			}
			c.setConnected(disconnected)
			log.Println("心跳关闭...")
			return
		}
	}
}

func (c *Client) RelayOpenAll() error {
	write, err := c.conn.Write(NewProtocol(OpenId).Pack())
	if err != nil {
		return err
	}
	log.Println(write)
	return errors.New("open ok..")
}

func (c *Client) RelayCloseAll() error {
	write, err := c.conn.Write(NewProtocol(CloseId).Pack())
	if err != nil {
		return err
	}
	log.Println(write)
	return errors.New("close ok..")
}

func (c *Client) IsConnected() bool {
	c.RLock()
	defer c.RUnlock()
	status := atomic.LoadUint32(&c.status)
	switch {
	case status == connected:
		return true
	case status > connecting:
		return true
	case status == connecting:
		return true
	default:
		return false
	}
}

func (c *Client) setConnected(status uint32) {
	c.Lock()
	defer c.Unlock()
	atomic.StoreUint32(&c.status, status)
}

func (c *Client) Close() {
	if c.IsConnected() {
		if c.stop != nil {
			c.stop <- struct{}{}
		}
		_ = c.conn.Close()
	}
}
