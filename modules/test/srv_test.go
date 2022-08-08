package main

import (
	"fmt"
	"golang.org/x/sys/windows/svc/mgr"
	"log"
	"testing"
)

func TestSrv(t *testing.T) {

	const name = "sshd"

	m, err := mgr.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer m.Disconnect()
	s, err := m.OpenService(name)
	if err != nil {
		log.Fatal(fmt.Errorf("could not access service: %v", err))
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(status.State)

}
