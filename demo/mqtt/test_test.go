package mqtt

import (
	"encoding/binary"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"log"
	"os"
	"testing"
	"time"
)

func TestConn(t *testing.T) {

	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker("ssl://xx9fx0x.mqtt.iot.gz.baidubce.com:1884")
	opts.SetClientID("zing")
	opts.SetProtocolVersion(4)
	opts.SetUsername("XXXXXX")
	opts.SetPassword("XXXXXXX")
	opts.AutoReconnect = true
	opts.CleanSession = true
	opts.ConnectTimeout = time.Second * 2
	opts.SetPingTimeout(10 * time.Second)
	opts.OnConnect = func(client mqtt.Client) {
		log.Println(fmt.Sprintln("mqtt连接成功"))
	}
	opts.OnConnectionLost = func(client mqtt.Client, e error) {
		log.Println(fmt.Sprintln("mqtt连接意外断开"))
	}
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("test", 0, func(client mqtt.Client, message mqtt.Message) {

		log.Println("=======================>>")
		log.Println(message.Topic(), string(message.Payload()))
		log.Println("=======================>>")

	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	time.Sleep(1000 * time.Second)
	c.Disconnect(250)

}

func TestByte(t *testing.T) {
	log.Println(1 << 3)
	log.Printf("%8b\n", 1)
	log.Printf("%8b\n", 1<<3|1<<6)
	log.Printf("%8d\n", 1<<3|1<<6)
	log.Printf("%8b\n", (1<<3)|1)
	log.Println((1 << 3) | 1)
}
func encodeBytes(field []byte) []byte {
	fieldLength := make([]byte, 2)
	binary.BigEndian.PutUint16(fieldLength, uint16(len(field)))
	return append(fieldLength, field...)
}
func TestName(t *testing.T) {
	//[0 5 104 101 108 108 111]
	log.Println(encodeBytes([]byte("hello")))
	//[0 12 104 101 108 108 111 32 119 111 114 108 100 33]
	log.Println(encodeBytes([]byte("hello world!")))
	log.Println(encodeBytes([]byte("")))
	//[0 3 230 136 145]
	log.Println(encodeBytes([]byte("我")))
	//[104 101 108 108 111]
	log.Println([]byte("hello"))
}
