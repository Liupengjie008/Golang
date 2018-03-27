创建消费者：consumer
package main

import (
    "fmt"
    "strings"
    "sync"
    "github.com/Shopify/sarama"
)

var (
    wg sync.WaitGroup
)

func main() {
    //创建消费者
    consumer, err := sarama.NewConsumer(strings.Split("192.168.1.125:9092", ","), nil)
    if err != nil {
        fmt.Println("Failed to start consumer: %s", err)
        return
    }
    //设置分区
    partitionList, err := consumer.Partitions("nginx_log")
    if err != nil {
        fmt.Println("Failed to get the list of partitions: ", err)
        return
    }
    fmt.Println(partitionList)
    //循环分区
    for partition := range partitionList {
        pc, err := consumer.ConsumePartition("nginx_log", int32(partition), sarama.OffsetNewest)
        if err != nil {
            fmt.Printf("Failed to start consumer for partition %d: %s\n", partition, err)
            return
        }
        defer pc.AsyncClose()
        go func(pc sarama.PartitionConsumer) {
            wg.Add(1)
            for msg := range pc.Messages() {
                fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s", msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
                fmt.Println()
            }
            wg.Done()
        }(pc)
    }
    //time.Sleep(time.Hour)
    wg.Wait()
    consumer.Close()
}




创建生产者：producer
package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"192.168.1.125:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}

	defer client.Close()
	for {
		msg := &sarama.ProducerMessage{}
		msg.Topic = "nginx_log"
		msg.Value = sarama.StringEncoder("this is a good test, my message is good")

		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send message failed,", err)
			return
		}

		fmt.Printf("pid:%v offset:%v\n", pid, offset)
		time.Sleep(10 * time.Millisecond)
	}
}



输出结果：
$ ls
consumer.go 	producer.go

启动zookeeper
启动kafka
启动kafka链接zookeeper

开启生产者：
$ go run producer.go 
pid:0 offset:1678
pid:0 offset:1679
pid:0 offset:1680
pid:0 offset:1681
pid:0 offset:1682
pid:0 offset:1683
pid:0 offset:1684
pid:0 offset:1685
pid:0 offset:1686


开启消费者：    //消费代码有问题待解决
$ go run consumer.go 
[0]
