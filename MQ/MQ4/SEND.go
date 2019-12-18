/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  SEND
 * @Version: 1.0.0
 * @Date: 2019/12/12 13:42
 */

package main

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/streadway/amqp"
)

// amqp://<user>:<password>@<ip>:<port>
var addr = "amqp://root:root@127.0.0.1:5672" //test

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
	// 定义交换“direct”、“fanout”、“topic”和“headers”
	err = ch.ExchangeDeclare("happy", amqp.ExchangeTopic, true, false, false, false, nil)
	if nil != err {
		logs.Error(err)
		return
	}

	data := fmt.Sprintf("hello,world!!!")
	//a.b.c.d.e 为发布key,以.分割；
	err = ch.Publish("happy", "a.b.c.d.e", false, false,
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(data),
			DeliveryMode: amqp.Transient, // 1=non-persistent, 2=persistent
		})
	if nil != err {
		logs.Error(err)
		return
	}

}