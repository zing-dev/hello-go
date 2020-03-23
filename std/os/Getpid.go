package os

import (
	"fmt"
	"os"
)

func Getpid() {

	fmt.Println("Getpid ", os.Getpid())

	fmt.Println("Getppid ", os.Getppid())

}
