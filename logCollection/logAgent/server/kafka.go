package server

import (
	"fmt"
	"sync"
	"time"

	"github.com/Shopify/sarama"
	"github.com/astaxie/beego/logs"
)

var kafkaWg sync.WaitGroup

func (t TailInfo) SendToKafka() {
	if t.ExitSign {
		wg.Done()
		return
	}
	for {
		message, ok := <-MessageChan
		if !ok {
			time.Sleep(time.Millisecond * 100) // 停止100毫秒
			continue
		}
		msg := &sarama.ProducerMessage{}
		msg.Topic = message.Topic
		msg.Value = sarama.StringEncoder(message.LineLog)
		_, _, err = KafkaClient.SendMessage(msg)
		if err != nil {
			err = fmt.Errorf("send message failed,", err)
			logs.Warn(err)
		}
		fmt.Println("send message to kafka", message.Topic, message.LineLog)
	}
	// kafkaWg.Wait()

}

/*
for {
		fmt.Println("send to kafka ")
		message, ok := <-t.MessageChan
		fmt.Println("send to kafka : ", message)
		if !ok {
			fmt.Println("not ok")
			time.Sleep(time.Millisecond * 100) // 停止100毫秒
			continue
		}
		kafkaWg.Add(1)
		msg := &sarama.ProducerMessage{}
		msg.Topic = message.Topic
		msg.Value = sarama.StringEncoder(message.LineLog)
		_, _, err = KafkaClient.SendMessage(msg)
		if err != nil {
			err = fmt.Errorf("send message failed,", err)
			logs.Warn(err)
		}
		fmt.Println("send message to kafka", message.Topic, message.LineLog)

	}
	kafkaWg.Wait()
*/
