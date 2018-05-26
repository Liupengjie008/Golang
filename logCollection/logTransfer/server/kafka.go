package server

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

func ReadKafka(c ConsumerInfo) {
	//设置分区
	partitionList, err := c.Consumer.Partitions(c.Topic)
	if err != nil {
		err = fmt.Errorf("Failed to get the list of partitions: ", err)
		logs.Warn(err)
		return
	}
	//循环分区
	for partition := range partitionList {
		pc, err := c.Consumer.ConsumePartition(c.Topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			err = fmt.Errorf("Failed to start consumer for partition %d: %s\n", partition, err)
			logs.Warn(err)
			return
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(pc sarama.PartitionConsumer) {
			defer wg.Done()
			for msg := range pc.Messages() {
				// fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
				// fmt.Println()
				c.MessageChan <- string(msg.Value)
				fmt.Println(string(msg.Value))
				// 如果 停止收集 消费掉kafka所有信息 返回
				if c.ExitSign && string(msg.Value) == "" {
					return
				}
			}

		}(pc)
	}
	wg.Wait()
}
