/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  log4go-3
 * @Version: 1.0.0
 * @Date: 11/11/2019 10:42 AM
 */

package main

import (
	"github.com/jeanphorn/log4go"
	log3_12 "oauth2Test/log/log3/log3-1"
	log3_22 "oauth2Test/log/log3/log3-2"
)

func main() {
	// load config file, it's optional
	// or log.LoadConfiguration("./example.json", "json")
	// config file could be json or xml
	log4go.LoadConfiguration("./trigger_log.json")
	
	//log.LOGGER("Test").Info("category Test info test message: %s", "new test msg")
	//log.LOGGER("Test").Debug("category Test debug test ...")
	//
	//// Other category not exist, test
	//log.LOGGER("DEBUG").Debug("category Other debug test ...")
	//
	//// socket log test
	//log.LOGGER("TestSocket").Debug("category TestSocket debug test ...")

	// original log4go test
	//log.Info("nomal info test ...")
	//log.Debug("nomal debug test ...")


	// ##########################  测试log


	log4go.LOGGER("DEBUG").Debug("category Other debug test ...")

	log4go.LOGGER("Info").Fine("Fine test")

	log4go.LOGGER("Trace").Trace("Trace test")

	log4go.Info("test")
	log4go.Warn("test")
	log4go.Error("test")

	log4go.LOGGER("Info").Info("Info test")

	log4go.LOGGER("Warn").Warn("warn test")

	log4go.LOGGER("Error").Error("error test")

	log3_12.Test()
	log3_22.Test()



	log4go.Close()

}