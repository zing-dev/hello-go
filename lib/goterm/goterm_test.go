// +build linux

package main_test

import (
	"github.com/google/goterm/term"
	"testing"
	"time"
)

func TestRead(t *testing.T) {
	pty, err := term.OpenPTY()
	if err != nil {
		t.Fatal(err)
	}
	name, err := pty.PTSName()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("PTSName: ", name)
	data := make([]byte, 10)
	go func() {
		var i byte = 0
		for {
			n, err := pty.Slave.Write([]byte{i})
			if err != nil {
				continue
			}
			t.Log("slave: write ", n, i)
			if i == 100 {
				break
			}
			i++
			time.Sleep(time.Second)
		}
	}()
	for {
		n, err := pty.Master.Read(data)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("read: ", n, data)
		if data[0] == 100 {
			break
		}
	}
}
