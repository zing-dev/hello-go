package net_test

import (
	"net"
	"testing"
)

func TestInterfaces(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range interfaces {
		t.Log("-----------------------------------")
		t.Log("Name: ", v.Name)
		t.Log("Flags: ", v.Flags)
		t.Log("Index: ", v.Index)
		t.Log("MTU: ", v.MTU)
		t.Log("HardwareAddr: ", v.HardwareAddr.String())
		addrs, err := v.Addrs()
		if err != nil {
			continue
		}
		for k, v := range addrs {
			t.Log(k, v.Network())
		}
	}
}

func TestInterfaceAddrs(t *testing.T) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		t.Fatal(err)
	}
	for k, v := range addrs {
		t.Log("---------------------------------")
		t.Log(k, v.Network())
	}
}
