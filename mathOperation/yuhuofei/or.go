/**
 * @Author: Guo PingFei
 * @Description:      测试go语言中或运算符
 * @File:  or
 * @Version: 1.0.0
 * @Date: 11/12/2019 1:50 PM
 */

package main

import "fmt"

func main() {

	var flag1 = true
	var flag2 = false
	var test1 = true
	var test2 = false

	if flag1 && flag2 {
		fmt.Println(" flag1 flag2 与运算")
	}

	if flag1 || flag2 {
		fmt.Println(" flag1 flag2 或运算")
	}

	if flag1 && test1 {
		fmt.Println(" flag1 test1 与运算")
	}

	if flag1 || test1 {
		fmt.Println(" flag1 test1 或运算")
	}

	if flag1 && test2 {
		fmt.Println(" flag1 test2 与运算")
	}

	if flag1 || test2 {
		fmt.Println(" flag1 test2 或运算")
	}
	
}
