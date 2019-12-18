/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  log4go-3
 * @Version: 1.0.0
 * @Date: 11/11/2019 10:42 AM
 */

package log3_2

import (
	"github.com/jeanphorn/log4go"
)

func Test() {

	// ##########################  测试log

	log4go.LOGGER("Info").Info("22222222222222222222              Info test")

	log4go.LOGGER("Info").Warn("22222222222222222222               warn test")

	log4go.LOGGER("Info").Error("22222222222222222222               error test")

}