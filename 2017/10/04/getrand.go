package main

import (
	"time"
	"fmt"
	"math/rand"
)


func Generate_Randnum() int{
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)

	fmt.Printf("rand is %v\n", rnd)

	return rnd
}

func main(){
	for true{
		Generate_Randnum()
		time.Sleep(time.Second)
	}

}