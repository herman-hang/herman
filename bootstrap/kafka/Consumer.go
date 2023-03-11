package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/fatih/color"
	"github.com/herman-hang/herman/bootstrap/core"
	"github.com/herman-hang/herman/servers/settings"
	"go.uber.org/zap"
	"sync"
)

// Consumer 消费者结构体
type Consumer struct {
	Topic        string
	MessageQueue chan []byte
}

//Consume 消费者
// @return void
func (k *Consumer) Consume() {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer([]string{fmt.Sprintf("%s:%d",
		settings.Config.Kafka.Host,
		settings.Config.Kafka.Port,
	)}, config)
	if err != nil {
		zap.S().Error(color.RedString(fmt.Sprintf("New consumer err: %v", err)))
		return
	}

	defer func(consumer sarama.Consumer) {
		if err := consumer.Close(); err != nil {
			zap.S().Error(color.RedString(fmt.Sprintf("Close consumer err: %v", err)))
			return
		}
	}(consumer)

	// 先查询该 topic 有多少分区
	partitions, err := consumer.Partitions(k.Topic)
	if err != nil {
		zap.S().Error(color.RedString(fmt.Sprintf("Partitions err: %v", err)))
		return
	}
	var wg sync.WaitGroup
	wg.Add(len(partitions))
	// 然后每个分区开一个 goroutine 来消费
	for _, partitionId := range partitions {
		// 不开异步会导致一个消费完才会消费另外一个
		go k.consumeByPartition(consumer, k.Topic, partitionId, &wg)
	}
	wg.Wait()
}

// consumeByPartition 执行消费
// @param sarama.Consumer consumer 消费者
// @param string topic 消息主题
// @param int32 partitionId 分区ID
// @param *sync.WaitGroup wg 同异步机制，用于管理goroutine便于执行完成返回一个信号
// @return void
func (k *Consumer) consumeByPartition(consumer sarama.Consumer, topic string, partitionId int32, wg *sync.WaitGroup) {
	defer wg.Done()
	partitionConsumer, err := consumer.ConsumePartition(topic, partitionId, sarama.OffsetNewest)
	if err != nil {
		zap.S().Error(color.RedString(fmt.Sprintf("Consume partition err: %v", err)))
		return
	}
	defer func(partitionConsumer sarama.PartitionConsumer) {
		if err := partitionConsumer.Close(); err != nil {
			zap.S().Error(color.RedString(fmt.Sprintf("Partition consumer err: %v", err)))
		}
	}(partitionConsumer)
	for message := range partitionConsumer.Messages() {
		core.Log.Infof("[Consumer] partitionid: %d; offset:%d, value: %s\n", message.Partition, message.Offset, string(message.Value))
		k.MessageQueue <- message.Value
	}
}
