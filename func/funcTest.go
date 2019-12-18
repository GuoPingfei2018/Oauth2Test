/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  funcTest
 * @Version: 1.0.0
 * @Date: 2019/12/6 14:46
 */

package main

import (
	"fmt"
	"strings"
)

var  Token []string

func main() {

	//Token := strings.Split("stest tstdsd"," ")
	test1()

}

func test1(){
	Token = strings.Split("stest tstdsd"," ")               // 全局中用"=", 局部用":="
 	fmt.Println("Token: ",Token[1])
 	test2()

}


func test2(){
	//Token := strings.Split("stest tstdsd"," ")
	//ss := Token[0]
	fmt.Println("ss: "+Token[0])
}