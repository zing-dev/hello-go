package os

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

func env() {

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
