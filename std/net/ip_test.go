package net_test

import (
	"net"
	"testing"
)

func TestIPNet(t *testing.T) {
	interfaces, err := net.Interfaces()
	if err != nil {
		t.Fatal(err)
	}
	for _, i := range interfaces {
		t.Log("----------------------")
		addrs, err := i.Addrs()
		if err != nil {
			t.Fatal(err)
		}
		for _, v := range addrs {
			i, ok := v.(*net.IPNet)
			if !ok {
				t.SkipNow()
			}
			t.Log("IP: ", i.IP.String())
			t.Log("Mask: ", i.Mask.String())
			t.Log("Network: ", i.Network())
			t.Log("IsGlobalUnicast: ", i.IP.IsGlobalUnicast())
			t.Log("Contains: ", i.Contains(net.ParseIP("192.168.0.61")))
		}
	}
}

func TestIPv4(t *testing.T) {
	pv4 := net.IPv4(192, 168, 0, 61)
	t.Log(pv4.String())
	t.Log("IsGlobalUnicast: ", pv4.IsGlobalUnicast())
	t.Log("IsInterfaceLocalMulticast: ", pv4.IsInterfaceLocalMulticast())
	t.Log("IsUnspecified: ", pv4.IsUnspecified())
	t.Log("To16: ", pv4.To16())
	t.Log("DefaultMask: ", pv4.DefaultMask().String())
}
