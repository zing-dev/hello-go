package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	file, _ = os.OpenFile(`C:\Windows\System32\drivers\etc\services2`, os.O_RDWR, 0777)
	snmp    = "snmp1              161/udp                           #SNMP\r\n"
)

func read1() {
	rd := bufio.NewReader(file)
	for {
		//跟ReadString功能相似，不过是返回的字节切片
		lineBytes, err := rd.ReadBytes('\n')
		if err == io.EOF {
			log.Println("读完了.....")
			break
		}
		if err != nil {
			log.Fatal("读错了....", err)
		}
		line := string(lineBytes)
		if strings.HasPrefix(line, "snmp") && strings.HasSuffix(line, "#SNMP\r\n") {
			log.Println("找到啦 >", line, "<")
		}
		//line, err := rd.ReadString('\n')
		//lineBytes, err := rd.ReadBytes('\n')
		//line := string(lineBytes)
		//line, err := rd.ReadString('\n')
		//line = strings.Trim(line, "\n")
	}
}
func read2() {
	content := ""
	flag := false
	rd := bufio.NewReader(file)
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
		if strings.HasPrefix(str, "1snmp") && strings.HasSuffix(str, "#SNMP") {
			log.Println("找到啦 >", str, "<")
			str = regexp.MustCompile(`\d+`).ReplaceAllString(str, "161")
			log.Println(str)
			flag = true
		}
		content += str + "\r\n"
	}
	seek, err := file.Seek(io.SeekStart, 0)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("seek:", seek)
	if !flag {
		log.Println("没找到...")
		content += snmp
		log.Println(content)
	}
	n, err := bufio.NewWriter(file).WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("n = ", n)
}
func main() {
	read2()
}
