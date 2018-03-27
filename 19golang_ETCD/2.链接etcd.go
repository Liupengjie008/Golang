package main

import (
    "fmt"
    "github.com/coreos/etcd/clientv3"
    "time"
)

func main() {
    /*
        DialTimeout time.Duration `json:"dial-timeout"`
        Endpoints []string `json:"endpoints"`
    */
    cli, err := clientv3.New(clientv3.Config{
        Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
        DialTimeout: 5 * time.Second,
    })
    if err != nil {
        fmt.Println("connect failed, err:", err)
        return
    }

    fmt.Println("connect succ")
    defer cli.Close()
}


输出结果：
$ go run main.go 
connect succ