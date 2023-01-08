package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fp/fp-gin-framework/app/common"
	"github.com/fp/fp-gin-framework/servers/settings"
	"time"
)

var syncProducer sarama.SyncProducer

// newSyncProducer 创建一个生产者
// @return producer err 返回一个生产者和错误信息
func newSyncProducer() (producer sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err = sarama.NewSyncProducer([]string{fmt.Sprintf("%s:%d",
		settings.Config.KafkaConfig.Host,
		settings.Config.KafkaConfig.Port,
	)}, config)

	//defer producer.Close()
	if err != nil {
		return nil, err
	}

	return producer, nil
}

// Send 发送消息到队列
// @param string topic 消息主题
// @param map[string]interface{} data 消息数据
// @return bool error 返回一个bool值和一个错误信息
func Send(topic string, data map[string]interface{}) {
	if syncProducer == nil {
		var err error
		syncProducer, err = newSyncProducer()
		if err != nil {
			common.Log.Errorf("New Sync Producer failed, err:%v", err)
			return
		}
	}
	jsonString, err := json.Marshal(data)
	if err != nil {
		common.Log.Errorf("Producer json failed, err:%v", err)
		return
	}
	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(time.Now().String()),
		Value: sarama.StringEncoder(jsonString),
	}

	// SendMessage：该方法是生产者生产给定的消息
	// 生产成功的时候返回该消息的分区和所在的偏移量
	// 生产失败的时候返回error
	partition, offset, err := syncProducer.SendMessage(msg)
	if err != nil {
		common.Log.Errorf("Producer send message failed, err:%v", err)
		return
	}
	common.Log.Infof("Partition = %d, offset=%d\n", partition, offset)
}
