package modbus

import (
	"bytes"
	"fmt"
	"github.com/goburrow/modbus"
	"math/bits"
	"testing"
	"time"
)

var (
	address                          = "COM3"
	handler *modbus.RTUClientHandler = nil
	client  modbus.Client
)

func init() {
	build()
}

func build() {
	handler = modbus.NewRTUClientHandler(fmt.Sprintf(`\\.\%s`, address))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "N"
	handler.Timeout = time.Second
	err := handler.Connect()
	if err != nil {
		panic(err)
	}
	handler.SlaveId = byte(1)
	client = modbus.NewClient(handler)
}

func openOne(address uint16) ([]byte, error) {
	return client.WriteSingleCoil(address, 0xff00)
}

func closeOne(address uint16) ([]byte, error) {
	return client.WriteSingleCoil(address, 0)
}

func statusOne(address, quantity uint16) ([]byte, error) {
	return client.ReadCoils(address, quantity)
}

func TestMarquee(t *testing.T) {
	i := 0
	for i < 32 {
		_, _ = openOne(uint16(i))
		time.Sleep(time.Second / 30)
		if i >= 4 {
			_, _ = closeOne(uint16(i - 4))
		} else {
			_, _ = closeOne(uint16(32 + i - 4))
		}
		i++
		if i == 32 {
			i = 0
		}
	}
}

func TestStatusOne(t *testing.T) {
	fmt.Println(4 & 1)
	fmt.Println(4 & 1 << 1)
	fmt.Println(4 & 1 << 2)
	fmt.Println(4 & 1 << 3)
	fmt.Println(fmt.Sprintf("%.8b", 4))
	fmt.Println(bits.Reverse8(4))
	fmt.Println(fmt.Sprintf("%.8b", bits.Reverse8(4)))

	_, _ = openOne(0)
	_, _ = openOne(10)
	_, _ = openOne(16)
	_, _ = openOne(20)
	_, _ = openOne(26)
	_, _ = openOne(31)
	one, _ := statusOne(0, 32)
	str := ""
	for len(one) < 4 {
		one = append(one, 0)
	}
	for k, v := range one {
		fmt.Println(k, v)
		str += fmt.Sprintf("%.8b", bits.Reverse8(v))
	}
	for k, c := range str {
		fmt.Printf("第 %d 路: ", k+1)
		if c == '1' {
			fmt.Println("开")
		} else {
			fmt.Println("关")
		}
	}
}

func TestOne(t *testing.T) {
	_, _ = openOne(1)
	fmt.Println(statusOne(1, 1))
	_, _ = closeOne(1)
	fmt.Println(statusOne(1, 1))

	_, _ = openOne(2)
	fmt.Println(statusOne(2, 1))

	fmt.Println(statusOne(0, 32))
	fmt.Println(statusOne(1, 31))
	fmt.Println(statusOne(0, 8))
	fmt.Println(statusOne(8, 16))
}

func TestIO3200Open(t *testing.T) {
	coil, err := openOne(0)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(coil, []byte{0xff, 00}) {
		t.Log("success")
	}
}

func TestIO3200Close(t *testing.T) {
	coil, err := closeOne(0)
	if err != nil {
		t.Fatal(err)
	}
	if bytes.Equal(coil, []byte{0, 00}) {
		t.Log("success")
	}
}

func TestIO3200Status(t *testing.T) {
	coil, err := statusOne(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(coil)
}
