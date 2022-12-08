package net

import (
	"fmt"
	"log"
	"net/netip"
	"net/url"
	"strconv"
	"testing"
)

func TestParseURL(t *testing.T) {
	for i, raw := range []string{
		"10.0.0.10",
		"tcp://10.0.0.10:8080",
		"http://10.0.0.10",
		"http://10.0.0.10:8080",
		"wss://10.0.0.10:8080",
	} {
		u, err := url.Parse(raw)
		if err != nil {
			log.Fatal(i, err)
		}
		fmt.Println(u.String(), u.Scheme, u.Hostname(), u.Port())
	}
	for i, raw := range []string{
		"10.0.0.10",
	} {
		addr, err := netip.ParseAddr(raw)
		if err != nil {
			log.Fatal(i, err)
		}
		fmt.Println(addr.String())
	}
	fmt.Println(strconv.Atoi(" "))
}
