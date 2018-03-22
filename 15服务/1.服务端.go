服务端： 
a. 监听端口 
b. 接收客户端的链接 
c. 创建goroutine，处理该链接

package main

import (
    "fmt"
    "net"
)

func main() {
    fmt.Println("start server...")
    listen, err := net.Listen("tcp", "0.0.0.0:50000")
    if err != nil {
        fmt.Println("listen failed, err:", err)
        return
    }
    for {
        conn, err := listen.Accept()
        if err != nil {
            fmt.Println("accept failed, err:", err)
            continue
        }
        go process(conn)
    }
}
func process(conn net.Conn) {
    defer conn.Close()
    for {
        buf := make([]byte, 512)
        n, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read err:", err)
            return
        }

        fmt.Printf(string(buf[0:n]))
    }
}


在终端启动service  --窗口1
$ go run main.go 
start server...

然后新建 --窗口2 
$ nc localhost 50000 	( nc 是 mac 命令 ，Windows和Linux命令是 telnet localhost 50000 )
输入：
aaaa
bbbb
cccc

窗口1显示：
start server...
aaaa
bbbb
cccc


