package server

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
)

func SendToEs(c ConsumerInfo) {
	if c.ExitSign {
		return
	}
	for {
		message, ok := <-c.MessageChan
		if !ok {
			time.Sleep(time.Millisecond * 100) // 停止100毫秒
			continue
		}
		if message == "" {
			continue
		}

		wg.Add(1)
		type Tweet struct {
			User    string
			Message string
		}
		tweet := Tweet{User: c.Topic, Message: message}
		_, err = EsClient.Index().
			Index(c.Topic).
			Type(c.Topic).
			// Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do()

		if err != nil {
			err = fmt.Errorf("send to es failed,err:%v", err)
			logs.Error(err)
			continue
		}
	}
	wg.Wait()
}
