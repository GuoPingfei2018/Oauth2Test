/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  stringPlus
 * @Version: 1.0.0
 * @Date: 2019/12/12 20:38
 */

package main

import (
	"fmt"
	"strings"
)

func main()  {
	hello := "hello"
	world := "world"
	var test string
	//for i := 0; i < 12; i++ {
	//	test = strings.Join([]string{hello, world}, "_")
	//}
	test = StringsJoin(hello,world,"_")
	
	fmt.Println("test: ",test)

	test1 := strings.Split(hello,"l")
	fmt.Println("test1: ",test1)
	TEL := "15512345678"
	test2 :=  Substring(TEL,len(TEL)-6,len(TEL))
	fmt.Println("test2: ",test2)

}


func StringsJoin(str1 string,str2 string, symbol string)(message string){
	lens := len(str1)+len(str2)
	for i := 0; i < lens; i++ {
		message = strings.Join([]string{str1, str2}, symbol)
	}
	return message
}


func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)
	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	var substring = ""
	for i := start; i < length; i++ {
		substring += string(r[i])
	}

	return substring
}