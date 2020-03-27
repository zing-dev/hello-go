package main

import (
	"log"
	"net"
	"os"
)

func main() {
	name := "WLAN"
	if len(os.Args) == 2 {
		name = os.Args[1]
	}
	log.Println("===============", name, "==v0.0.1============")
	face, err := net.InterfaceByName(name)
	if err != nil {
		log.Fatal("没找到网络接口", err)
	}
	addrs, err := face.Addrs()
	if err != nil {
		log.Fatal("获取地址表失败", err)
	}
	for _, v := range addrs {
		if ip, ok := v.(*net.IPNet); ok && ip.IP.IsGlobalUnicast() {
			log.Println(ip.IP.To4())
		}
	}
}
