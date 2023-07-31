package main

import (
	"bufio"
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"log"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestProc(t *testing.T) {
	result, err := gproc.ShellExec(gctx.New(), "tasklist  /FI 'IMAGENAME eq msedge.exe' /FO csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func TestProc2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	result, err := gproc.ShellExec(ctx, "systeminfo /FO csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	reader := csv.NewReader(bytes.NewBufferString(result))
	record, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}
	for i, s := range record {
		fmt.Println(i, s)
	}

}

func Test3(t *testing.T) {
	cmd := exec.Command("cmd", "/c", "systeminfo | grep 'System Boot Time'")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	line := strings.Split(string(out), ":")
	if len(line) < 2 {
		fmt.Println("Failed to get the boot time")
		return
	}
	bootTime := strings.TrimSpace(line[1])
	fmt.Println("Boot Time:", bootTime)
}

func Test4(t *testing.T) {
	cmd := exec.Command("cmd", "/c", "systeminfo")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	buffer := bytes.NewBufferString(string(out))
	scanner := bufio.NewScanner(buffer)

	// 逐行读取并打印
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Host Name") {
			line := strings.Split(line, ":")
			if len(line) < 2 {
				continue
			}
			fmt.Println(strings.TrimSpace(line[0]), ":", strings.TrimSpace(line[1]))
		}
		if strings.HasPrefix(line, "OS Name") {
			line := strings.Split(line, ":")
			if len(line) < 2 {
				continue
			}
			fmt.Println(strings.TrimSpace(line[0]), ":", strings.TrimSpace(line[1]))
		}
		if strings.HasPrefix(line, "System Boot Time:") {
			space := strings.TrimSpace(strings.TrimPrefix(line, "System Boot Time:"))
			fmt.Println("System Boot Time:", space)
			parse, err := time.Parse("1/02/2006, 3:04:05 PM", space)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(parse.Format(time.DateTime))
		}
	}

	// 检查错误
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
