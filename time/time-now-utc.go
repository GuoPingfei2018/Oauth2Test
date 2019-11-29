/**
 * @Author: Guo PingFei
 * @Description:    time.Now().Unix() unix时间戳是从1970年1月1日（UTC/GMT的午夜）开始所经过的秒数，不考虑闰秒。
 * @File:  time-now-utc
 * @Version: 1.0.0
 * @Date: 11/12/2019 4:24 PM
 */

package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(time.Now())


	// func Now() Time
	fmt.Println(time.Now().Unix())


	// func Now() Time
	fmt.Println(time.Now().UTC().Unix())

	 fmt.Println("********************************")


	time.Sleep(time.Duration(2)*time.Second)

	fmt.Println(time.Now())

	// func Now() Time
	fmt.Println(time.Now().Unix())


	// func Now() Time
	fmt.Println(time.Now().UTC().Unix())
	
}


