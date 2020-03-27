package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func ip1() {

	pv4 := net.IPv4(127, 0, 0, 1)
	fmt.Println(pv4.Equal(net.IP{}))
	fmt.Println(pv4.Equal(pv4)) //true

	fmt.Println(pv4.String())             //127.0.0.1
	fmt.Printf("%d\n", pv4.DefaultMask()) //[255 0 0 0]

	fmt.Println(pv4.IsLoopback()) //true

	//全球单播地址
	fmt.Println(pv4.IsGlobalUnicast()) //false

	//本地组播地址
	fmt.Println(pv4.IsInterfaceLocalMulticast()) //false

	//链路本地组播地址
	fmt.Println(pv4.IsLinkLocalMulticast()) //false

	//路本地单播地址
	fmt.Println(pv4.IsLinkLocalUnicast()) //false

	fmt.Println(pv4.To16()) //127.0.0.1

	text := make([]byte, 0)
	p := &pv4
	fmt.Println(p.UnmarshalText(text))
	fmt.Printf("%s\n", text)
	fmt.Printf("%p\n", p)
}

func ip2() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok /*&& !ipnet.IP.IsLoopback() */ {
			//if ipnet.IP.To4() != nil && ipnet.IP.To4().IsGlobalUnicast() {
			if ipnet.IP.To4() != nil && ipnet.IP.To4().IsGlobalUnicast() {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

func ip3() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(addrs[0].(*net.IPNet))
	face, err := net.InterfaceByName("WLAN")
	if err != nil {
		log.Fatal(err)
	}
	addrs, err = face.Addrs()
	if err != nil {
		log.Fatal(err)
	}

	for k, v := range addrs {
		ipNet := v.(*net.IPNet)
		if ipNet.IP.IsGlobalUnicast() {
			log.Println(k, v.(*net.IPNet).IP.To4())
		}
	}
}
