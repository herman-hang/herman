package jobs

import (
	"fmt"
	"github.com/herman-hang/herman/kernel/kafka"
)

// Consumer 调用消费者执行消费
// @param topic string 消费主题
// @return kafkaConsumer 返回kafka消费者结构体
func Consumer(topic string) (kafkaConsumer kafka.Consumer) {
	kafkaConsumer = kafka.Consumer{
		Topic:        topic,
		MessageQueue: make(chan []byte, 1000),
	}

	kafkaConsumer.Exec()

	return kafkaConsumer
}

// Dispatch 消息队列
// @param data 待消费数据
// @param jobFunc 消费者函数
// @return 返回一个闭包
func Dispatch(data map[string]interface{}, jobFunc func(topic string)) {
	topic := fmt.Sprintf("%s", data["topic"])
	// 调用生产者
	go kafka.Send(topic, data["data"].(map[string]interface{}))
	// 调用消费者
	go jobFunc(topic)
}
