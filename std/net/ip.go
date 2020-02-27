package main

import (
	"fmt"
	"net"
)

func main() {

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
