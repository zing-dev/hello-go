package main

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"time"
)

const DefaultIniName = "config.ini"

type Cfg struct {
	Address  string
	Username string
	Password string
}

func config() *Cfg {
	f, err := ini.ShadowLoad(DefaultIniName)
	if err != nil {
		file, err := os.Create(DefaultIniName)
		if err != nil {
			log.Panicln(err)
		}
		err = file.Close()
		f, err = ini.ShadowLoad(DefaultIniName)
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := &Cfg{}
	err = f.Section("MQtt").MapTo(cfg)
	if err != nil || cfg.Address == "" {
		cfg = &Cfg{
			Address: "192.168.0.215",
		}
		err := f.Section("MQtt").ReflectFrom(cfg)
		if err != nil {
			log.Println("MQtt ReflectFrom", err)
		}
		err = f.SaveTo(DefaultIniName)
		if err != nil {
			log.Println(err)
		}
	}
	log.Println("address:", cfg.Address)
	return cfg
}

func main() {
	opts := mqtt.NewClientOptions()
	opts.SetClientID("zing")
	opts.SetProtocolVersion(4)
	opts.AutoReconnect = false
	opts.CleanSession = false
	opts.ConnectTimeout = time.Second * 3
	opts.SetPingTimeout(10 * time.Second)
	for {
		cfg := config()
		opts.SetUsername(cfg.Username)
		opts.SetPassword(cfg.Password)
		opts.Servers = nil
		opts.AddBroker(cfg.Address)
		c := mqtt.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
		}
		if true == c.IsConnected() {
			log.Println("mqtt已连接....")
		} else {
			log.Println("mqtt未连接....")
		}
		time.Sleep(time.Second * 2)
	}
}
