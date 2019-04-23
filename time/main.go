package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now().String()) //2019-04-23 13:41:51.825852 +0800 CST m=+0.001974201

	fmt.Println(time.Friday) //Friday

	fmt.Println(time.Now().Year())          // 2019
	fmt.Println(time.Now().Local().Month()) //April
	fmt.Println(time.Now().Day())           // 23
	fmt.Println(time.Now().Hour())          // 20
	fmt.Println(time.Now().Minute())        // 10
	fmt.Println(time.Now().Second())        // 55

	date := time.Date(2019, 1, 1, 1, 1, 1, 1, time.Local)
	fmt.Println(date.Year())

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().Add(1000 * 1000 * 1000).String())           // + 1s
	fmt.Println(time.Now().Add(1000 * 1000 * 1000 * 60 * 60).String()) // + 1h

	fmt.Println(time.Nanosecond)  // 1ns 一纳秒
	fmt.Println(time.Microsecond) // 1µs 一微秒
	fmt.Println(time.Millisecond) // 1ms 一毫秒
	fmt.Println(time.Second)      // 1s 一秒
	fmt.Println(time.Minute)      // 1m0s 一分钟
	fmt.Println(time.Hour)        // 1h0m0s 一小时

	//minDuration Duration = -1 << 63
	fmt.Println(-1 << 63) // -9223372036854775808
	//maxDuration Duration = 1<<63 - 1
	fmt.Println(1<<63 - 1) // 9223372036854775807

	now := time.Now
	fmt.Println(now()) //2019-04-23 20:22:10.6935552 +0800 CST m=+0.012686901

	n := now()
	fmt.Println(n)

	fmt.Println(n.Unix()) //1556022256

	fmt.Println(n.Zone()) //CST 28800

	fmt.Println(n.Local()) //2019-04-23 20:25:10.4033476 +0800 CST

	fmt.Println(n.AddDate(10, 1, 1)) //2029-05-24 20:25:40.6828302 +0800 CST

	fmt.Println(n.After(now().AddDate(0, 0, -1))) //true
	fmt.Println(n.Before(now().Add(1)))           //true

	fmt.Println(n.Clock()) //20 27 51
	hour, min, sec := n.Clock()
	fmt.Printf("%d:%d:%d\n", hour, min, sec) //20:29:7

	fmt.Println(n.Equal(now()))        //true
	fmt.Println(n.Equal(now().Add(1))) //false

	//"2006-01-02 15:04:05.999999999 -0700 MST"
	fmt.Println(n.Format("2006-01-02 15:04:05"))
	fmt.Println(n.Format("MST"))     //CST
	fmt.Println(n.Format("January")) //April
	fmt.Println(n.Format("02"))      //23
	fmt.Println(n.Format("_2"))      //23

	fmt.Println(n.YearDay()) //113

	fmt.Println(n.Truncate(time.Minute)) //2019-04-23 20:49:00 +0800 CST
	fmt.Println(n.Truncate(time.Hour))   // 2019-04-23 20:00:00 +0800 CST

	fmt.Println(n.Round(time.Hour)) //2019-04-23 21:00:00 +0800 CST

	fmt.Println(time.Hour.Hours())

	fmt.Println(time.Duration(time.Hour * 10).Seconds()) // 36000秒
	fmt.Println(time.Duration(time.Hour * 10).Truncate(time.Minute).Minutes()) // 600分钟
}
