package main

import (
	"fmt"
)

func main() {

	//slice

	s1 := make([]byte, 10)
	fmt.Println(s1) // [0 0 0 0 0 0 0 0 0 0]

	i := byte(0)
	for i < byte(len(s1)) {
		s1[i] = i
		i++
	}
	fmt.Println(s1) //[0 1 2 3 4 5 6 7 8 9]

	fmt.Println(s1[0])   //0
	fmt.Println(s1[0:0]) //[]
	fmt.Println(s1[0:1]) //[0]
	fmt.Println(s1[0:5]) //[0 1 2 3 4]

	fmt.Println(s1[3:5]) //[3 4]
	fmt.Println(s1[3:3]) //[]

	fmt.Println(s1[1:3:5]) //[1 2]
	fmt.Println(s1[1:5:8]) //[1 2 3 4]

	fmt.Println(s1[1:]) //[1 2 3 4 5 6 7 8 9]
	fmt.Println(s1[:7]) //[0 1 2 3 4 5 6]
	fmt.Println(s1[:])  //[0 1 2 3 4 5 6 7 8 9]

	fmt.Println(len(s1)) //10
	fmt.Println(cap(s1)) //10

	fmt.Println(len(s1[1:3:5])) //2 = 3-1
	fmt.Println(cap(s1[1:3:5])) //4 = 5-1

	s2 := append(s1, 10, 11, 12)
	fmt.Println(s1)      //[0 1 2 3 4 5 6 7 8 9 10 11 12]
	fmt.Println(len(s1)) //10
	fmt.Println(cap(s1)) //10

	//先将旧的slice容量乘以2，如果乘以2后的容量仍小于新的slice容量，则取新的slice容量(append多个elems)
	//如果新slice小于等于旧slice容量的2倍，则取旧slice容量乘以2
	//如果旧的slice容量大于1024，则新slice容量取旧slice容量乘以1.25
	fmt.Println(len(s2)) //13
	fmt.Println(cap(s2)) //32 = 16 * 2

	//s2 := make([]string, 10)
	//fmt.Println(s2) //[         ]

}
