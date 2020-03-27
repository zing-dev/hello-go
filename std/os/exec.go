package os

import (
	"bufio"
	"bytes"
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
