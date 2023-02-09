package time

import (
	"fmt"
	"testing"
	"time"
)

func LocalDiffDateTime(start, end time.Time) string {
	var (
		str  = ""
		diff = start.Sub(end)
		abs  = diff.Abs()
	)

	if y := int(abs.Hours()) / 24 / 365; y != 0 {
		str += fmt.Sprintf("%d年", y)
	}
	if d := int(abs.Hours()) / 24; d != 0 {
		str += fmt.Sprintf("%02d天", d)
	}
	if h := int(abs.Hours()) % 24; h != 0 {
		str += fmt.Sprintf("%02d小时", h)
	}
	if m := int(abs.Minutes()) % 60; m != 0 {
		str += fmt.Sprintf("%02d分钟", m)
	}
	if s := int(abs.Seconds()) % 60; s != 0 {
		str += fmt.Sprintf("%02d秒", s)
	}
	if diff > 0 {
		return fmt.Sprintf("%s后", str)
	} else if diff == 0 {
		return "刚刚"
	} else {
		return fmt.Sprintf("%s前", str)
	}
}

func Test1(t *testing.T) {
	start := time.Now()
	end := time.Now().Add(time.Hour*24 + time.Hour*4 + time.Minute*30 + time.Second)
	diff := end.Sub(start)
	fmt.Println(diff)
	fmt.Println(int(diff.Hours())/24, "天")    // 1 d
	fmt.Println(int(diff.Hours())%24, "小时")   // 3 h
	fmt.Println(int(diff.Minutes())%60, "分钟") // 3 m
	fmt.Println(int(diff.Seconds())%60, "秒")  // 3 m
	fmt.Println(LocalDiffDateTime(start, end))
}
