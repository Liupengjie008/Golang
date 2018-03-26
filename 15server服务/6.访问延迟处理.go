访问延迟处理
package main

import (
    "fmt"
    "net/http"
    "net"
    "time"
)

var url = []string{
    "http://www.baidu.com",
    "http://google.com",
    "http://taobao.com",
}

func main() {

    for _, v := range url {
        //type Client struct{}
        c := http.Client{
            //type Transport struct {}
            Transport: &http.Transport {
                //Dial func(network, addr string) (net.Conn, error)
                Dial:func(network, addr string) (net.Conn, error){
                    timeout := time.Second*2
                    return net.DialTimeout(network, addr, timeout)
                },
            },
        }
        resp, err := c.Head(v)
        if err != nil {
            fmt.Printf("head %s failed, err:%v\n", v, err)
            continue
        }

        fmt.Printf("head %s succ, status:%v\n", v ,resp.Status)
    }
}



$ go run main.go 
输出结果：

head http://www.baidu.com succ, status:200 OK
head http://google.com failed, err:Head http://google.com: dial tcp [2404:6800:4008:801::200e]:80: i/o timeout
head http://taobao.com succ, status:200 OK