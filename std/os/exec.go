package os

import (
	"log"
	"os/exec"
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
