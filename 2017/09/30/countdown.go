package main

import (
	"fmt"
	"time"
)

func main() {
	for true {
		now := time.Now()
		gohome := time.Date(now.Year(), now.Month(), now.Day(), 18, 0, 0, 0, time.Local).Unix()
		var diff int = int(gohome - now.Unix())
		if diff < 0 {
			diff += 3600 * 24
		}
		var h int = int(diff / 3600)
		var m int = int((diff - h*3600) / 60)
		var s int = int(diff % 60)
		fmt.Println("还剩", h, "小时", m, "分钟", s, "秒下班了!")
		time.Sleep(time.Second * 2)
	}

}
