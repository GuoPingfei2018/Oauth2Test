/**
 * @Author: Guo PingFei
 * @Description:
 * @File:  map[string][]sting
 * @Version: 1.0.0
 * @Date: 2019/12/6 11:29
 */

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type FailCommands struct {
	Ids       []string `json:"ids"`
	Status    string   `json:"status"`
	States    States   `json:"states"`
	ErrorCode string   `json:"errorCode,omitempty"`
}
type States struct {
	On     bool `json:"on"` //关于电源
	Online bool `json:"online"`
}

//格式version1
type PowerVersion1 struct {
	RequestID string `json:"requestId"`
	Payload   struct {
		Commands []interface{} `json:"commands"`
	} `json:"payload"`
}

func main() {
	var ErrorDeviceInfo map[string][]string
	ErrorDeviceInfo = make(map[string][]string)
	var respPowerVersion PowerVersion1

	var respPowerVersiontest PowerVersion1

	var comment2 FailCommands
	var comment3 FailCommands

	var res = "设备不在线"
	var res2 = "设备key为空或设备不存在"
	respPowerVersion.RequestID = "test"
	if strings.Contains(res, "设备不在线") || strings.Contains(res, "未收到") {
		ErrorDeviceInfo["deviceTurnedOff"] = append(ErrorDeviceInfo["deviceTurnedOff"],"111")
	}

	if strings.Contains(res2, "空或设备不存在") {
		ErrorDeviceInfo["deviceNotFound"] = append(ErrorDeviceInfo["deviceNotFound"],"222")
	}


	if len(ErrorDeviceInfo["deviceTurnedOff"]) > 0 {

		comment2.Ids = append(comment2.Ids, ErrorDeviceInfo["deviceTurnedOff"]...)
		comment2.Status = "ERROR"
		comment2.ErrorCode = "deviceTurnedOff"
		respPowerVersion.Payload.Commands = append(respPowerVersion.Payload.Commands, comment2)
	}

	if len(ErrorDeviceInfo["deviceNotFound"]) > 0 {

		comment3.Ids = append(comment3.Ids, ErrorDeviceInfo["deviceNotFound"]...)
		comment3.Status = "ERROR"
		comment3.ErrorCode = "deviceNotFound"
		respPowerVersion.Payload.Commands = append(respPowerVersion.Payload.Commands, comment3)
	}
	respPowerVersionJson,_ := json.Marshal(respPowerVersion)
	fmt.Println(respPowerVersionJson)
	err := json.Unmarshal([]byte(respPowerVersionJson),&respPowerVersiontest)
 	fmt.Println(err)
	fmt.Println(respPowerVersiontest)





}
