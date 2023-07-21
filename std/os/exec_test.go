package os

import (
	"log"
	"os/exec"
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

func TestLookPath(t *testing.T) {
	path, err := exec.LookPath("git")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(path)
	cmd := exec.Command(path, "version")
	res, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	log.Println(string(res))
}

func TestCommand(t *testing.T) {
	cmd := exec.Command("go", "mod", "graph")
	data, err := cmd.Output()
	if err != nil {
		t.Fatal(err)
	}
	log.Print(string(data))
}

func TestCombinedOutput(t *testing.T) {
	p, err := exec.LookPath("go")
	if err != nil {
		t.Fatal(err)
	}
	cmd := exec.Cmd{
		Path: p,
		Args: []string{
			"mod", "graph",
		},
	}
	log.Println(cmd.CombinedOutput())
}

func TestExec(t *testing.T) {
	output, err := exec.Command("cmd", "/C", "1.cmd").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(output))
}
