/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  Foo
 * @Version: 1.0.0
 * @Date: 11/11/2019 2:29 PM
 */

package Foo

import "fmt"

func TestPanic() {

	defer func (){
		fmt.Println("defer end...")
	}()

	panic("error test.....")

	fmt.Println("END")
	
}
