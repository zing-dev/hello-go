// +build windows

package modbus_test

import (
	"fmt"
	"github.com/tarm/serial"
	"github.com/xiegeo/modbusone"
	"os"
	"os/signal"
	"testing"
)

func TestModbusOne(t *testing.T) {
	const size = 0x10000

	var discretes [size]bool
	var coils [size]bool
	var inputRegisters [size]uint16
	var holdingRegisters [size]uint16

	config := serial.Config{
		Name:     `\\.\COM14`,
		Baud:     9600,
		StopBits: 1,
		Parity:   'N',
	}
	s, err := serial.OpenPort(&config)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "open serial error: %v\n", err)
		os.Exit(1)
	}
	com := modbusone.NewSerialContext(s, 9600)
	defer func() {
		fmt.Printf("%+v\n", com.Stats())
		_ = com.Close()
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		fmt.Printf("%+v\n", com.Stats())
		fmt.Println("close serial port")
		_ = com.Close()
		os.Exit(0)
	}()

	id, err := modbusone.Uint64ToSlaveID(1)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "set slaveID error: %v\n", err)
		os.Exit(1)
	}
	var device modbusone.Server
	device = modbusone.NewRTUServer(com, id)
	h := modbusone.SimpleHandler{
		ReadDiscreteInputs: func(address, quantity uint16) ([]bool, error) {
			fmt.Printf("ReadDiscreteInputs from %v, quantity %v\n", address, quantity)
			return discretes[address : address+quantity], nil
		},
		WriteDiscreteInputs: func(address uint16, values []bool) error {
			fmt.Printf("WriteDiscreteInputs from %v, quantity %v\n", address, len(values))
			for i, v := range values {
				discretes[address+uint16(i)] = v
			}
			return nil
		},

		ReadCoils: func(address, quantity uint16) ([]bool, error) {
			fmt.Printf("ReadCoils from %v, quantity %v\n", address, quantity)
			return coils[address : address+quantity], nil
		},
		WriteCoils: func(address uint16, values []bool) error {
			fmt.Printf("WriteCoils from %v, quantity %v\n", address, len(values))
			for i, v := range values {
				coils[address+uint16(i)] = v
			}
			return nil
		},

		ReadInputRegisters: func(address, quantity uint16) ([]uint16, error) {
			fmt.Printf("ReadInputRegisters from %v, quantity %v\n", address, quantity)
			return inputRegisters[address : address+quantity], nil
		},
		WriteInputRegisters: func(address uint16, values []uint16) error {
			fmt.Printf("WriteInputRegisters from %v, quantity %v\n", address, len(values))
			for i, v := range values {
				inputRegisters[address+uint16(i)] = v
			}
			return nil
		},

		ReadHoldingRegisters: func(address, quantity uint16) ([]uint16, error) {
			fmt.Printf("ReadHoldingRegisters from %v, quantity %v\n", address, quantity)
			return holdingRegisters[address : address+quantity], nil
		},
		WriteHoldingRegisters: func(address uint16, values []uint16) error {
			fmt.Printf("WriteHoldingRegisters from %v, quantity %v\n", address, len(values))
			for i, v := range values {
				holdingRegisters[address+uint16(i)] = v
			}
			return nil
		},

		OnErrorImp: func(req modbusone.PDU, errRep modbusone.PDU) {
			fmt.Printf("error received: %v from req: %v\n", errRep, req)
		},
	}
	err = device.Serve(&h)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "serve error: %v\n", err)
		os.Exit(1)
	}

}
