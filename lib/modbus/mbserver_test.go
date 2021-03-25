// +build windows

package modbus_test

import (
	"encoding/json"
	"github.com/goburrow/modbus"
	"github.com/goburrow/serial"
	"github.com/tbrandon/mbserver"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	//s, err := tty()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//go runServer("/dev/pts/7")
	go runServer(`\\.\COM14`)
	//go runClient("/dev/pts/8")
	//go runClient(`\\.\COM10`)
	time.Sleep(time.Minute * 10)
}

//func tty() (string, error) {
//	pty, err := term.OpenPTY()
//	if err != nil {
//		return "", err
//	}
//	return pty.PTSName()
//}

//func TestTTY(t *testing.T) {
//	t.Log(tty())
//}

func runServer(address string) {
	server := mbserver.NewServer()
	log.Println("server: ", address)
	err := server.ListenRTU(&serial.Config{
		Address:  address,
		BaudRate: 14400,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  time.Second * 100,
	})
	if err != nil {
		log.Println("TCPClient...")
		err = server.ListenTCP("127.0.0.1:3333")
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Println("RTUServer...")
	}
	defer server.Close()
	for k := range server.HoldingRegisters {
		//server.HoldingRegisters[k] = uint16(rand.Intn(30)*10 + 500)
		server.HoldingRegisters[k] = uint16(k)
	}

	for {
		/*for k := range server.HoldingRegisters {
			//server.HoldingRegisters[k] = uint16(rand.Intn(30)*10 + 500)
			server.HoldingRegisters[k] = uint16(k)
		}
		log.Println("server: update...")*/
		//server.Coils
		time.Sleep(time.Second * 10)
	}
}

func TestServer(t *testing.T) {
	runServer(`\\.\COM10`)
}

func runClient(address string) {
	log.Println("client: ", address)
	handler = modbus.NewRTUClientHandler(address)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.Timeout = time.Second * 100
	handler.SlaveId = byte(1)
	err := handler.Connect()
	if err != nil {
		log.Println("TCPClient...")
		client = modbus.TCPClient("127.0.0.1:3333")
	} else {
		log.Println("RTUClient...")
		client = modbus.NewClient(handler)
	}
	defer handler.Close()
	//res, err := client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	//if err != nil {
	//	log.Fatal("WriteMultipleRegisters", err)
	//} else {
	//	log.Println("res", res)
	//}
	for {
		address := rand.Intn(125)
		quantity := rand.Intn(10) + 1
		log.Println("send: ", address, quantity)
		registers, err := client.ReadHoldingRegisters(uint16(address), uint16(quantity))
		if err != nil {
			log.Println("client: ", err)
			continue
		}
		log.Println("recv: ", len(registers)/2, registers)
		time.Sleep(time.Second)
	}
}

func TestClient(t *testing.T) {
	//runClient("/dev/pts/4")
	runClient(`\\.\COM14`)
}

func TestModbus(t *testing.T) {
	// Server
	s := mbserver.NewServer()
	err := s.ListenRTU(&serial.Config{
		Address:  `\\.\COM14`,
		BaudRate: 9600,
		DataBits: 8,
		StopBits: 1,
		Parity:   "N",
		Timeout:  time.Second * 10,
	})
	//err = s.ListenTCP("127.0.0.1:3333")
	if err != nil {
		t.Fatalf("failed to listen, got %v\n", err)
	}
	defer s.Close()

	// Allow the server to start and to avoid a connection refused on the client
	time.Sleep(1 * time.Millisecond)

	// Client
	handler := modbus.NewRTUClientHandler(`\\.\COM16`)
	handler.BaudRate = 9600
	handler.StopBits = 1
	handler.DataBits = 8
	handler.Parity = "N"
	handler.Timeout = time.Second * 10
	handler.SlaveId = byte(1)
	//handler = modbus.NewTCPClientHandler("127.0.0.1:3333")
	// Connect manually so that multiple requests are handled in one connection session
	err = handler.Connect()
	if err != nil {
		t.Errorf("failed to connect, got %v\n", err)
		t.FailNow()
	}
	defer handler.Close()
	client := modbus.NewClient(handler)

	// Coils
	results, err := client.WriteMultipleCoils(100, 9, []byte{255, 1})
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	} else {
		t.Log("WriteMultipleCoils:", results)
	}

	results, err = client.ReadCoils(100, 16)
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	} else {
		t.Log("ReadCoils:", results)
	}
	expect := []byte{255, 1}
	got := results
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	// Discrete inputs
	results, err = client.ReadDiscreteInputs(0, 64)
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	}
	// test: 2017/05/14 21:09:53 modbus: sending 00 01 00 00 00 06 ff 02 00 00 00 40
	// test: 2017/05/14 21:09:53 modbus: received 00 01 00 00 00 0b ff 02 08 00 00 00 00 00 00 00 00
	expect = []byte{0, 0, 0, 0, 0, 0, 0, 0}
	got = results
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	// Holding registers
	results, err = client.WriteMultipleRegisters(1, 2, []byte{0, 3, 0, 4})
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	}
	// received: 00 01 00 00 00 06 ff 10 00 01 00 02
	expect = []byte{0, 2}
	got = results
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	results, err = client.ReadHoldingRegisters(1, 2)
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	}
	expect = []byte{0, 3, 0, 4}
	got = results
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}

	// Input registers
	s.InputRegisters[65530] = 1
	s.InputRegisters[65535] = 65535
	results, err = client.ReadInputRegisters(65530, 6)
	if err != nil {
		t.Errorf("expected nil, got %v\n", err)
		t.FailNow()
	}
	expect = []byte{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255}
	got = results
	if !isEqual(expect, got) {
		t.Errorf("expected %v, got %v", expect, got)
	}
}

func isEqual(a interface{}, b interface{}) bool {
	expect, _ := json.Marshal(a)
	got, _ := json.Marshal(b)
	if string(expect) != string(got) {
		return false
	}
	return true
}
