package os

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

func exec1() {
	cmd := exec.Command("cmd", "/C", "dir", ".")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))

	output, err = exec.Command("cmd", "/C", "123").Output()
	if err != nil {
		log.Println(">", err, "<")
	}
	log.Println(string(output))

	err = exec.Command("cmd", "/C", "123").Run()
	if err != nil {
		log.Println(">>>", string(err.(*exec.ExitError).Stderr))
		log.Println(">", err, "<")
	}
	log.Println("==============================================================")
	output, err = exec.Command("cmd", "/C", `netstat -an |find /c "8888"`).Output()
	if err != nil {
		log.Println(">", err, "<")
	}
	log.Println(string(output))

}

func exec2() {
	output, err := exec.Command("cmd", "/C", `file`).Output()
	if err != nil {
		log.Fatal(err)
	}
	rd := bufio.NewReader(bytes.NewBuffer(output))
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			log.Println("读完了.....")
			break
		}
		if err != nil {
			log.Fatal("读错了....", err)
		}
		str := string(line)
		if strings.HasPrefix(str, "Telnet") {
			all := regexp.MustCompile(`\d+`).FindAllString(str, -1)
			log.Println(all)
		}
	}
}

func exec3() {
	output, err := exec.Command("cmd", "/C", `CPUInfoService.exe`).Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
	/*rd := bufio.NewReader(bytes.NewBuffer(output))
	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			log.Println("读完了.....")
			break
		}
		if err != nil {
			log.Fatal("读错了....", err)
		}
		str := string(line)
		if strings.HasPrefix(str, "Telnet") {
			all := regexp.MustCompile(`\d+`).FindAllString(str, -1)
			log.Println(all)
		}
	}*/
}

func exec4() {
	output, err := exec.Command("cmd", "/C", "ping 192.168.0.81 -n 1").Output()
	if err != nil {
		log.Fatal(err)
	}
	/**
	Pinging 192.168.0.81 with 32 bytes of data:
	Reply from 192.168.0.111: Destination host unreachable.

	Ping statistics for 192.168.0.81:
	    Packets: Sent = 1, Received = 1, Lost = 0 (0% loss),

	*/
	/**
	Pinging 192.168.0.83 with 32 bytes of data:
	Reply from 192.168.0.83: bytes=32 time=8ms TTL=128

	Ping statistics for 192.168.0.83:
	    Packets: Sent = 1, Received = 1, Lost = 0 (0% loss),
	Approximate round trip times in milli-seconds:
	    Minimum = 8ms, Maximum = 8ms, Average = 8ms
	*/
	log.Println(string(output))
}

func exec5(ip string) error {
	const PingUnreachable = "Destination host unreachable"
	const PingUnreachable2 = "无法访问目标主机"
	output, err := exec.Command("cmd", "/C", fmt.Sprintf(`ping %s -n 1`, ip)).Output()
	if err != nil {
		return errors.New("执行ping命令失败")
	}
	if strings.Contains(string(output), PingUnreachable) || strings.Contains(string(output), PingUnreachable2) {
		return nil
	}
	return errors.New("")
}

func exec6(ip string) error {
	output, err := exec.Command("cmd", "/C", `arp -a`).Output()
	if err != nil {
		return errors.New("执行ping命令失败")
	}
	if !strings.Contains(string(output), ip) {
		return nil
	}
	return errors.New("")
}
