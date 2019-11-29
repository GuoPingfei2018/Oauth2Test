/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  log2.2
 * @Version: 1.0.0
 * @Date: 11/8/2019 2:32 PM
 */

package main

import (
"fmt"
"github.com/op/go-logging"
"os"
"time"
)
//创建一个名字为examplename的日志对象log
var log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
//创建一个日志输出格式对象format，也就是用什么格式输出
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func main() {

	logFile, err := os.OpenFile("./logs/log1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	logFile1, err := os.OpenFile("./logs/log2.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println(err)
	}

	// For demo purposes, create two backend for os.Stderr.
	//创建一个日志输出对象backend，也就是日志要打印到哪儿，在此是标准错误输出
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(logFile1, "", 0)

	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	//将输出格式与输出对象绑定
	//backend2Formatter := logging.NewBackendFormatter(backend2, format)
	backend2Leveled := logging.AddModuleLevel(backend2)
	backend2Leveled.SetLevel(logging.ERROR, "")


	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	backend1Leveled.SetLevel(logging.INFO, "")

	// Set the backends to be used.
	//将绑定了格式的输出对象设置为日志的输出对象
	//这样log打印每一句话都会按格式输出到backendFormatter所代表的对象里，在此即是标准错误输出
	logging.SetBackend(backend1Leveled, backend2Leveled)


	log.Debugf(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Debugf------ " , Password("secret"))
	log.Info(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Info------ " )
	log.Notice(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Notice------ " )
	log.Warning(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Warning------ ")
	log.Error(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Error------ ")
	log.Critical(time.Now().Format("2006-01-02 15:04:05") + ":" + " ------Critical------ ")


}