package main

import (
	"bytes"
	"io"
	"net"
	"time"
)

func run() {

	//addr, _ := net.ResolveTCPAddr("tcp4", ":17082")
	addr, _ := net.ResolveTCPAddr("tcp4", ":9000")
	conn, _ := net.DialTCP("tcp", nil, addr)

	group.Add(1)
	go func() {
		data := make([]byte, 1024)
		var cache bytes.Buffer

		for {
			n, err := conn.Read(data)
			cache.Write(data[:n])
			if n > 0 && err != io.EOF {
				//log.Println(string(data), n)
				//log.Println(data[4], byteToInt(data[:4]), data[data[4]+5:])
				//log.Println(string(data[4]), string(byteToInt(data[:4])), string(data[data[4]+5:]))
				for {
					if unpack(&cache) {
						break
					}
				}
			}
		}
	}()

	group.Add(1)
	go func() {
		var data []byte
		data, _ = pack(nil, byte(Open))
		_, _ = conn.Write(data)
		time.Sleep(time.Second)

		data, _ = pack(nil, byte(Close))
		_, _ = conn.Write(data)
		time.Sleep(time.Second)

		data, _ = pack(nil, byte(Reset))
		_, _ = conn.Write(data)
		time.Sleep(time.Second)
	}()
	group.Wait()

}
