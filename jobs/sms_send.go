package jobs

import (
	"encoding/json"
	"fmt"
	SmsConstant "github.com/herman-hang/herman/app/constants/sms"
	"github.com/herman-hang/herman/kernel/core"
	"github.com/herman-hang/herman/servers/settings"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// SendSms 发送短信队列
// @param string topic 消息主题
// @return void
func SendSms(topic string) {
	var data map[string]interface{}
	// 调用消费者对数据进行消费，并返回结构体
	kafkaConsumer := ExecConsumer(topic)
	for {
		// 从通道取出消费的数据
		message := <-kafkaConsumer.MessageQueue
		// 将取出的JSON数据转为map
		if err := json.Unmarshal(message, &data); err != nil {
			core.Log.Errorf("Consumer sms json data failed, err:%v", err)
		}
		execSend(data)
	}
}

// execSend 执行发送
// @param map[string]interface{} data 带发送数据
// @return void
func execSend(data map[string]interface{}) {
	// 发起http请求
	response, err := http.Get(fmt.Sprintf("%ssms?u=%s&p=%s&m=%s&c=%s",
		settings.Config.Sms.Api,
		settings.Config.Sms.User,
		settings.Config.Sms.Password,
		data["mobile"],
		url.QueryEscape(fmt.Sprintf("%s", data["content"])),
	))
	if err != nil {
		core.Log.Errorf("Sms send failed, mobile:%s content:%s err:%v", data["mobile"], data["content"], err)
		return
	}

	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			core.Log.Errorf("Sms send close failed, err:%v", err)
		}
	}(response.Body)

	bodyBytes, _ := ioutil.ReadAll(response.Body)
	// 转为字符串
	code := string(bodyBytes)
	if SmsConstant.Status[code] != SmsConstant.SendSuccess {
		core.Log.Errorf("Sms send failed, mobile:%s content:%s err:%v", data["mobile"], data["content"], SmsConstant.Status[code])
	}
}
