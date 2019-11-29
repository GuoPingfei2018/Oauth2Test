/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  Sprintf2
 * @Version: 1.0.0
 * @Date: 11/18/2019 9:32 AM
 */

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)
func Md5(msg string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(msg))
	return hex.EncodeToString(md5Ctx.Sum(nil))
}
func main() {

	d := time.Now()
	// 格式化拼接函数
	s1 := fmt.Sprintf("%d.%d", d.Unix(), d.Nanosecond())
	fmt.Println("S1的值："+s1)
	fmt.Println("###############################")
	 s := Md5(fmt.Sprintf("%d.%d", d.Unix(), d.Nanosecond()))[0:12]
	 fmt.Println("S的值："+s)

}
