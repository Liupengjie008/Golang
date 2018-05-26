package server

import (
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/hpcloud/tail"
)

type TailInfo struct {
	//这里是每个读取日志文件的对象
	Tail     *tail.Tail
	SecLimit *SecondLimit
	Offset   int64 //记录当前位置
	//filename string
	LogConf  CollectionConf
	ExitSign bool
}
type Message struct {
	LineLog string
	Topic   string
}

var MessageChan chan *Message = make(chan *Message, 1000000)

func TailStart(t TailInfo) {
	if t.ExitSign {
		wg.Done()
		return
	}
	//读取文件路径
	filename := t.LogConf.LogPath
	tails, err := tail.TailFile(filename, tail.Config{
		ReOpen:    true,
		Follow:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //偏移量（日志收集第二次读取位置）
		MustExist: false,
		Poll:      true,
	})
	if err != nil {
		fmt.Println("tail file err:", err)
		return
	}

	t = TailInfo{
		Tail:     tails,
		SecLimit: NewSecondLimit(int32(t.LogConf.SendRate)),
		Offset:   0,
		//filename string
		LogConf:  t.LogConf,
		ExitSign: false,
		// MessageChan: make(chan Message, 1000000),
	}

	var msg *tail.Line
	var ok bool
	for {
		msg, ok = <-t.Tail.Lines
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		if msg.Err != nil {
			err = fmt.Errorf("get tails lines err : %v , time is : %v", err, msg.Time)
			logs.Warn(err)
			continue
		}
		str := strings.TrimSpace(msg.Text)
		if len(str) == 0 || str[0] == '\n' {
			continue
		}
		fmt.Println("tail:", msg.Text)
		MessageChan <- &Message{
			LineLog: msg.Text,
			Topic:   t.LogConf.Topic,
		}
	}

}

/*
msg, ok = <-t.Tail.Lines
if !ok {
	fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
	time.Sleep(100 * time.Millisecond)
	continue
}
if msg.Err != nil {
	err = fmt.Errorf("get tails lines err : %v , time is : %v", err, msg.Time)
	logs.Warn(err)
	continue
}
str := strings.TrimSpace(msg.Text)
if len(str) == 0 || str[0] == '\n' {
	continue
}
fmt.Println("tail:", msg.Text)
MessageChan <- &Message{
	LineLog: msg.Text,
	Topic:   t.LogConf.Topic,
}
*/
