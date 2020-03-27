package os

import (
	"log"
	"testing"
)

func TestExec1(t *testing.T) {
	exec1()
}

func TestExec2(t *testing.T) {
	exec2()
}

func TestExec3(t *testing.T) {
	exec3()
}

func TestExec4(t *testing.T) {
	exec4()
}

func TestExec5(t *testing.T) {
	if exec5("192.168.0.215") == nil {
		log.Println("不存在")
	} else {
		log.Println("存在")
	}
	log.Println(exec5("192.168.0.215"))
	if err := exec5("192.168.0.215"); err != nil {
		log.Println("存在", err)
	} else {
		log.Println("不存在")
	}
}

func TestExec6(t *testing.T) {
	log.Println(exec6("192.168.0.215"))
	log.Println(exec6("192.168.0.216"))
}
