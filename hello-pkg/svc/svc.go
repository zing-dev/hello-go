package svc

import (
	"fmt"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"syscall"
)

var (
	beepFunc = syscall.MustLoadDLL("user32.dll").MustFindProc("MessageBeep")
)


func Beep() {
	beepFunc.Call(0xffffffff)
}

func GetServiceStatus(name string) {
	if name == "" {
		name = "sshd"
	}
	status, err := getServiceStatus(name)
	if err != nil {
		log.Fatalln(err)
	}
	switch status.State {
	case svc.Running:
		fmt.Println(fmt.Sprintf("%s running", name))
	case svc.Paused:
		fmt.Println(fmt.Sprintf("%s paused", name))
	case svc.Stopped:
		fmt.Println(fmt.Sprintf("%s stopped", name))
	}
}

func getServiceStatus(name string) (*svc.Status, error) {
	connect, err := mgr.Connect()
	if err != nil {
		return nil, err
	}

	service, err := connect.OpenService(name)
	if err != nil {
		return nil, err
	}

	status, err := service.Query()
	if err != nil {
		return nil, err
	}
	return &status, err
}
