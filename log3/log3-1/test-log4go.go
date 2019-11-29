/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  log4go-3
 * @Version: 1.0.0
 * @Date: 11/11/2019 10:42 AM
 */

package log3_1

import (
	"github.com/jeanphorn/log4go"
)

func Test() {

	// ##########################  测试log
	log4go.LOGGER("Info").Info("11111111111111111111              Info test")

	log4go.LOGGER("Warn").Warn("11111111111111111111              warn test")

	log4go.LOGGER("Error").Error("11111111111111111111              error test")

}