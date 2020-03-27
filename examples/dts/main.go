package main

import (
	"github.com/Atian-OE/DTSSDK_Golang/dtssdk"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("===========dts连接测试v0.0.1===============")
	c := make(chan *dtssdk.DTSSDKClient)
	ip := "127.0.0.1"
	if len(os.Args) == 2 {
		ip = os.Args[1]
	}
	for {
		go func() {
			client := dtssdk.NewDTSClient(ip)
			client.CallConnected(func(s string) {
				c <- client
				client.Close()
			})
		}()
		select {
		case <-c:
			log.Println("dts连接 [", ip, "] 成功....")
			time.Sleep(time.Second * 3)
		case <-time.After(time.Second * 3):
			log.Println("dts连接 [", ip, "] 失败....")
		}
	}
}
