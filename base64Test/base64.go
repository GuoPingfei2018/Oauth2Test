/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  base64
 * @Version: 1.0.0
 * @Date: 11/27/2019 2:04 PM
 */

package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func main() {
	//读原图片
	ff, _ := os.Open("./base64Test/auth.png")
	defer ff.Close()
	sourcebuffer := make([]byte, 500000)
	n, _ := ff.Read(sourcebuffer)
	//base64压缩
	sourcestring := base64.StdEncoding.EncodeToString(sourcebuffer[:n])

	//写入临时文件
	ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)
	//读取临时文件
	cc, _ := ioutil.ReadFile("a.png.txt")

	//解压
	dist, _ := base64.StdEncoding.DecodeString(string(cc))
	//写入新文件
	f, _ := os.OpenFile("./base64Test/test.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)
	
}
