/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  recover
 * @Version: 1.0.0
 * @Date: 2019/12/12 13:46
 */

package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

// amqp://<user>:<password>@<ip>:<port>
//var addr = "amqp://test:test123@127.0.0.1:5672" //test

func main() {

	// 建立连接
	conn, err := amqp.Dial(addr)
	if nil != err {
		logs.Error(err)
		return
	}

	defer conn.Close()
	// 申请通道
	ch, err := conn.Channel()
	if nil != err {
		logs.Error(err)
		return
	}

	defer ch.Close()
	// 定义交换
	err = ch.ExchangeDeclare("happy", amqp.ExchangeTopic, true, false, false, false, nil)
	if nil != err {
		logs.Error(err)
		return
	}

	queName := "test.test1.test2"
	topic := "a.#"


	// 定义通道
	que, err := ch.QueueDeclare(queName, false, false, false, false, nil)
	if nil != err {
		logs.Error(err)
	}

	err = ch.QueueBind(que.Name, topic, "happy", false, nil)
	if nil != err {
		logs.Error(err)
		return
	}

	msges, err := ch.Consume(que.Name, "", true, false, false, false, nil)
	if nil != err {
		logs.Error(err)
		return
	}

	logs.Info("start recv")

	for msg := range msges {
		fmt.Println(">>> %s", string(msg.Body))

	}

}
