package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	file, _ = os.OpenFile(`C:\Windows\System32\drivers\etc\services2`, os.O_RDWR, 0777)
	snmp    = "snmp              161/udp                           #SNMP\r\n"
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

func read3() {
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
		if strings.HasPrefix(str, "snmp") && strings.HasSuffix(str, "#SNMP") {
			all := regexp.MustCompile(`\d+`).FindAllString(str, -1)
			log.Println(all)
		}
	}
}

func read4() {
	str := `
以下是 localhost 上的设置

Alt 键已被映射到 'CTRL+A' :     YES
Idle 会话超时             :     1 小时
最多连接次数              :     2
Telnet 端口               :     23
失败的登录企图的最多次数  :     3
断开时结束任务            :     YES
操作模式                  :     Console
身份验证机制              :     NTLM, Password
默认域                    :     ATDTSMKC4L50022
状态                      :     已停止`
	/*rd := bufio.NewReader(bytes.NewBuffer([]byte(str)))
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
	fmt.Println(str)
}

func main() {
	read4()
}
