package os

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
)

func TestEnv(t *testing.T) {
	environ := os.Environ()

	for key, value := range environ {
		fmt.Printf("%d : %s\n", key, value)
	}

	path := os.Getenv("PATH")

	paths := make([]string, 0)

	if runtime.GOOS == "linux" {
		paths = strings.Split(path, ":")
		fmt.Println(os.Getegid())
		fmt.Println(os.Geteuid())
		fmt.Println(os.Getuid())

	}

	if runtime.GOOS == "windows" {
		paths = strings.Split(path, ";")
	}

	for key, value := range paths {
		if value == "" {
			continue
		}
		fmt.Printf("%d %s\n", key, value)
	}

	//windows zing
	//linux xenial
	fmt.Println(os.Hostname())

	//windows C:\Users\zhang
	//linux error
	//fmt.Println(os.UserHomeDir())

}

func TestGetEnv(t *testing.T) {
	value := os.Getenv("go")
	log.Println(value)
	err := os.Setenv("go", "hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(os.Getenv("go"))
}

func TestLookupEnv(t *testing.T) {
	value, exist := os.LookupEnv("go")
	log.Println(value, exist)

	err := os.Setenv("go", "hello")
	if err != nil {
		log.Fatal(err)
	}

	value, exist = os.LookupEnv("go")
	log.Println(value, exist)

	value, exist = os.LookupEnv("GOPROXY")
	log.Println(value, exist)
}

func TestExpandEnv(t *testing.T) {
	_ = os.Setenv("NAME", "gopher")
	_ = os.Setenv("BURROW", "/usr/gopher")

	log.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))
	log.Println(os.ExpandEnv("$name lives in ${burrow}."))
}
