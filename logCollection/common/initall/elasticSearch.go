package initall

import (
	"fmt"

	"github.com/astaxie/beego/logs"

	elastic "gopkg.in/olivere/elastic.v2"
)

func InitES() (client *elastic.Client, err error) {
	client, err = elastic.NewClient(elastic.SetSniff(LogConfAll.EsConf.EsSniff), elastic.SetURL(LogConfAll.EsConf.EsAddr...))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}
	logs.Error("init ES success")
	return
}
