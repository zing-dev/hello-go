package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//svc.GetServiceStatus("dfvs")

	//ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	//process, err := os.StartProcess("C:\\Users\\admin\\OneDrive\\workspace\\rust\\dfvs\\target\\release\\dfvs.exe",nil,nil)
	//if err != nil {
	//	return
	//}
	//process.Wait()
	//cmd := exec.Command("C:\\Users\\admin\\OneDrive\\workspace\\rust\\dfvs\\target\\release\\dfvs.exe")
	os.Chdir("C:\\Users\\admin\\OneDrive\\workspace\\rust\\dfvs\\target\\release")
	cmd := exec.Command("dfvs.exe")
	fmt.Println("1111111111111111")
	read, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(false)
	}
	fmt.Println("222222222222")
	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("333333333333")
	reader := bufio.NewReader(read)
	fmt.Println(reader.Buffered())
	for reader.Buffered() > 0 {
		line, prefix, err := reader.ReadLine()
		fmt.Println("44444444444")

		fmt.Println(line, prefix, err)
	}
	fmt.Println("55555555555555")

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}


	//outr, outw, err := os.Pipe()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer outr.Close()
	//defer outw.Close()
	//
	//inr, inw, err := os.Pipe()
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer inr.Close()
	//defer inw.Close()
	//
	//proc, err := os.StartProcess("", flag.Args(), &os.ProcAttr{
	//	Files: []*os.File{inr, outw, outw},
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//
	//inr.Close()
	//outw.Close()
	//
	//stdoutDone := make(chan struct{})
	//go pumpStdout(ws, outr, stdoutDone)
	//go ping(ws, stdoutDone)
	//
	//pumpStdin(ws, inw)
	//
	//// Some commands will exit when stdin is closed.
	//inw.Close()
	//
	//// Other commands need a bonk on the head.
	//if err := proc.Signal(os.Interrupt); err != nil {
	//	log.Println("inter:", err)
	//}
	//
	//select {
	//case <-stdoutDone:
	//case <-time.After(time.Second):
	//	// A bigger bonk on the head.
	//	if err := proc.Signal(os.Kill); err != nil {
	//		log.Println("term:", err)
	//	}
	//	<-stdoutDone
	//}
	//
	//if _, err := proc.Wait(); err != nil {
	//	log.Println("wait:", err)
	//}
}
