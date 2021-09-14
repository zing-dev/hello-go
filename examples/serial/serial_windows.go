package main

import (
	"github.com/tarm/serial"
	"log"
	"time"
)

func conn() *serial.Port {
	log.Println("start connect")
	var port, err = serial.OpenPort(&serial.Config{
		Name:        "COM4",
		Baud:        9600,
		ReadTimeout: time.Millisecond * 200,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("open success")
	return port
}

func main() {
	test()
}

func test() {
	port := conn()
	log.Println("start send")
	n, err := port.Write([]byte{0x55, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x66})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stat send: ", n)
	buf := make([]byte, 16)
	n, err = port.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stat recv: ", n)
}

func test2() {
	port := conn()
	go func() {
		buf := make([]byte, 64)
		for {
			n, err := port.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("receive: ", n, buf[:n])
		}
	}()
	//time.Sleep(time.Minute)
	for {
		for k, v := range [][]byte{
			{0x55, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x66},
			//{0x55, 0x01, 0x12, 0x00, 0x00, 0x01, 0x02, 0x6b},
			//{0x55, 0x01, 0x11, 0x00, 0x00, 0x01, 0x02, 0x6a},
		} {
			n, err := port.Write(v)
			if err != nil {
				log.Fatal(err)
			}
			switch k {
			case 0:
				log.Println("stat send: ", n)
				log.Println(v)
			case 1:
				log.Println("open send: ", n)
			case 2:
				log.Println("clos send: ", n)
			}
			time.Sleep(time.Second * 5)
		}
	}
}
