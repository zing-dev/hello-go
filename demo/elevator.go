package main

import (
	"fmt"
	"math"
)

var A = [8]bool{true, false, false, false, false, false, false, false}
var B = [8]bool{true, false, false, false, false, false, false, false}
var ak, bk = 0, 0

func main() {
	fmt.Println("初始电梯状态")
	fmt.Println(A)
	fmt.Println(B)
	var i int
	for true {
		fmt.Println("选择去的楼层,退出按0:")
		_, _ = fmt.Scanln(&i)
		if i < 0 || i > 8 {
			fmt.Println("楼层数错误")
		} else if i == 0 {
			break
		} else {
			handle(i)
			fmt.Println(A)
			fmt.Println(B)
		}
	}
}

func handle(i int) {

	for key, val := range A {
		if val {
			ak = key
		}
	}

	for key, val := range B {
		if val {
			bk = key
		}
	}

	if (i-1) == ak || (i-1) == bk {
		return
	}
	if math.Abs(float64(i-1-ak)) <= math.Abs(float64(i-1-bk)) {
		A[i-1] = true
		A[ak] = false
	} else {
		B[i-1] = true
		B[bk] = false
	}
}
