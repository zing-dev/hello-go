package modbus_test

import (
	"bytes"
	"fmt"
	"github.com/goburrow/modbus"
	"testing"
	"time"
)

func TestModBus(t *testing.T) {
	handler := modbus.NewRTUClientHandler(fmt.Sprintf(`\\.\%s`, "COM3"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.Timeout = time.Second
	err := handler.Connect()
	if err != nil {
		t.Fatal("Connect", err)
		return
	}
	handler.SlaveId = byte(1)
	client := modbus.NewClient(handler)
	coil, err := client.WriteSingleCoil(0, 0xFF00)
	coil, err = client.WriteSingleCoil(31, 0xFF00)
	time.Sleep(time.Second)
	coil, err = client.WriteSingleCoil(0, 0)
	coil, err = client.WriteSingleCoil(31, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(coil)

	registers, err := client.ReadHoldingRegisters(0, 32)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(registers)

	readCoils, err := client.ReadCoils(0, 0x20)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%8b", readCoils))
	fmt.Println("readCoils", readCoils)

	response, err := handler.Send([]byte{0xFE, 0x01, 0x00, 0x00, 0x00, 0x10, 0x29, 0xC9})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("%x", response))
	fmt.Println(response[0] == 0xfe)
	fmt.Println(response[1] == 0x01)
	fmt.Println(response[2] == 0x01)
	fmt.Println(response[3] == 0x00)
	fmt.Println(response[4] == 0x61)
	fmt.Println(response[5] == 0x9c)

	decode, err := handler.Decode(response)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(decode.FunctionCode, decode.Data)

	response, _ = handler.Send([]byte{0xFE, 0x05, 0x00, 0x1E, 0xFF, 0x00, 0xF8, 0x33})
	fmt.Println(response[3] == 0x1e)
	singleCoil, err := client.WriteSingleCoil(29, 0xff00)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(singleCoil, bytes.Equal(singleCoil, []byte{0xff, 0}))

	//inputRegisters, err := client.ReadInputRegisters(0, 32)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(inputRegisters)

	//coil, err = client.WriteMultipleCoils(1, 2, []byte{0x001,0x002})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//fmt.Println(coil)

	coils, err := client.ReadCoils(0, 32)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(coils))
	for k, b := range coils {
		fmt.Println(k, b)
	}
	time.Sleep(time.Second * 3)
	_ = handler.Close()
}
