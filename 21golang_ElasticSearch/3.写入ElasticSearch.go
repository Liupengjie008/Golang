package main

import (
    "fmt"
    "gopkg.in/olivere/elastic.v2"
)
//go get gopkg.in/olivere/elastic.v2
type Tweet struct {
    User    string
    Message string
}

func main() {
    client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://192.168.1.103:9200/"))
    if err != nil {
        fmt.Println("connect es error", err)
        return
    }

    fmt.Println("conn es succ")

    for i := 0; i < 20; i++ {
        tweet := Tweet{User: "olivere", Message: "Take Five"}
        _, err = client.Index().
            Index("twitter").
            Type("tweet").
            Id(fmt.Sprintf("%d", i)).
            BodyJson(tweet).
            Do()
        if err != nil {
            // Handle error
            panic(err)
            return
        }
    }

    fmt.Println("insert succ")
}



输出结果：
$ go run main.go 
conn es succ
insert succ


浏览器配置索引模式：
http://localhost:5601/app/kibana#/management/kibana/index?_g=()

配置索引模式
为了使用Kibana，您必须配置至少一个索引模式。索引模式用于识别运行搜索和分析的Elasticsearch索引。它们也用于配置字段。

索引名称或模式：
	twitter

点击 Discover 按钮查看数据：
http://localhost:5601/app/kibana#/discover?_g=()&_a=(columns:!(_source),index:twitter,interval:auto,query:(match_all:()),sort:!(_score,desc))
