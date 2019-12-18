/**
 * @Author: Guo PingFei
 * @Description:log测试不同日志级别分为不同的包
 * @File:  log1
 * @Version: 1.0.0
 * @Date: 11/8/2019 2:05 PM
 */

package main
import (
	"log"
	"os"
)
func init() {
	file := "./logs/" +"message"+ ".txt"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	log.SetPrefix("[qSkipTool]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return
}

func main() {
	log.Println("Hello Davis!") // log 还是可以作为输出的前缀
	return
}
