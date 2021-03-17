package main

import (
	"fmt"
	"github.com/albenik/go-serial/v2"
	"github.com/albenik/go-serial/v2/enumerator"
	"log"
	"testing"
	"time"
)

func TestGetPortsList(t *testing.T) {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		log.Fatal("No serial ports found!")
	}
	for _, port := range ports {
		fmt.Printf("Found port: %v\n", port)
	}
}

func TestGetDetailedPortsList(t *testing.T) {
	ports, err := enumerator.GetDetailedPortsList()
	if err != nil {
		log.Fatal(err)
	}
	if len(ports) == 0 {
		fmt.Println("No serial ports found!")
		return
	}
	for _, port := range ports {
		fmt.Printf("Found port: %s\n", port.Name)
		if port.IsUSB {
			fmt.Printf("   USB ID     %s:%s\n", port.VID, port.PID)
			fmt.Printf("   USB serial %s\n", port.SerialNumber)
		}
	}
}

func TestOpen(t *testing.T) {
	dev, err := serial.Open("/dev/pts/4",
		serial.WithBaudrate(9600),
		serial.WithParity(serial.NoParity),
		serial.WithDataBits(8),
		serial.WithStopBits(serial.OneStopBit),
		serial.WithReadTimeout(2),
		serial.WithWriteTimeout(2),
	)
	if err != nil {
		t.Fatal(err)
	}

	log.Println("start send")
	err = dev.SetFirstByteReadTimeout(100)
	if err != nil {
		log.Fatal(err)
	}
	//n, err := dev.Write([]byte{0x55, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x66})
	n, err := dev.Write([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stat send: ", n)
	buf := make([]byte, 16)
	n, err = dev.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("stat recv: ", n, string(buf))
	dev.Close()
}

func TestName(t *testing.T) {
	dev, err := serial.Open("COM3",
		serial.WithBaudrate(9600),
		serial.WithParity(serial.NoParity),
		serial.WithDataBits(8),
		serial.WithStopBits(serial.OneStopBit),
		serial.WithReadTimeout(200),
		serial.WithWriteTimeout(200),
	)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		buf := make([]byte, 64)
		for {
			n, err := dev.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			if n != 0 {
				log.Println("receive: ", n, buf[:n])
			}
		}
	}()
	for {
		for k, v := range [][]byte{
			{0x55, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x66}, //status
			{0x55, 0x01, 0x12, 0x00, 0x00, 0x01, 0x02, 0x6b}, //open
			{0x55, 0x01, 0x11, 0x00, 0x00, 0x01, 0x02, 0x6a}, //close
			{0x55, 0x01, 0x33, 0xFF, 0xFF, 0xFF, 0xFF, 0x85}, //open all
			{0x55, 0x01, 0x10, 0x00, 0x00, 0x00, 0x00, 0x66}, //status
			{0x55, 0x01, 0x33, 0x00, 0x00, 0x00, 0x00, 0x89}, //close all
		} {
			n, err := dev.Write(v)
			if err != nil {
				log.Fatal(err)
			}
			switch k {
			case 0, 4:
				log.Println("stat send: ", n)
				log.Println(v)
			case 1:
				log.Println("open send: ", n)
			case 2:
				log.Println("clos send: ", n)
			}
			time.Sleep(time.Second * 3)
		}
	}
}
