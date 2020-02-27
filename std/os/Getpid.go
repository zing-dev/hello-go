package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Getpid ", os.Getpid())

	fmt.Println("Getppid ", os.Getppid())

}
