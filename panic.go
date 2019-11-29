/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  panic
 * @Version: 1.0.0
 * @Date: 11/11/2019 2:26 PM
 */

package main

import (
	"fmt"
	"oauth2Test/Foo"
)

func main() {

	fmt.Println(11111)
	fmt.Println(2222)
	fmt.Println(3333)
	fmt.Println(4444)

	for i := 0; i < 6; i++ {
		fmt.Println(i)
		go Foo.TestPanic()
	}


	fmt.Println(5555)
	fmt.Println(666)
	fmt.Println(777)
	fmt.Println(888)

	
}
