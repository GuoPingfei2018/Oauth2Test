/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  jsonTest
 * @Version: 1.0.0
 * @Date: 2019/12/5 14:04
 */

// 参考  https://blog.csdn.net/suiban7403/article/details/79175583
package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type People struct {
	Name string
}

func main() {
	fmt.Println("###############   0   #############")
	test()
	fmt.Println("###############   1   #############")
   	test1()
	fmt.Println("###############   2   #############")
	test2()
}

func test()  {
	data := People{Name: "happy"}
	dataJson,_ := json.Marshal(data)
	dataJsonString := string(dataJson)
	response := &Response{
		Code: 1,
		Msg:  "success",
		Data: dataJsonString,
	}

	b, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(b))
}

func test1()  {
	data := People{Name: "happy"}
	response := &Response{
		Code: 1,
		Msg:  "success",
		Data: data,
	}

	b, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(b))
}

func test2()  {
	data := `{"Name":"Happy"}`
	response := &Response{
		Code: 1,
		Msg:  "success",
		Data: json.RawMessage(data),
	}

	b, err := json.Marshal(&response)
	if err != nil {
		fmt.Println("err", err)
	}

	fmt.Println(string(b))
}