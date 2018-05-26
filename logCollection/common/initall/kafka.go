package initall

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

/*
	初始化NewConfig配置 sarama.NewConfig
	创建生产者sarama.NewSyncProducer
	创建消息sarama.ProducerMessage
	发送消息client.SendMessage
*/
func InitKafka() (client sarama.SyncProducer, err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer(LogConfAll.KafkaConf.KafkaAddr, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	logs.Error("init kafka success")
	return
}

// 创建消费者
func InitKafkaConsumer() (consumer sarama.Consumer, err error) {
	consumer, err = sarama.NewConsumer(LogConfAll.KafkaConf.KafkaAddr, nil)
	if err != nil {
		err = fmt.Errorf("Failed to start consumer: %s", err)
		return
	}
	logs.Error("init kafka Consumer success")
	return
}
