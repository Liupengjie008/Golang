客户端VS服务端处理流程
a. 建立与服务端的链接 
b. 进行数据收发 
c. 关闭链接

这个例子需要另外的客户端client代码 
clien.go

package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {

    conn, err := net.Dial("tcp", "localhost:50000")
    if err != nil {
        fmt.Println("Error dialing", err.Error())
        return
    }

    defer conn.Close()
    inputReader := bufio.NewReader(os.Stdin)
    for {
        input, _ := inputReader.ReadString('\n')
        trimmedInput := strings.Trim(input, "\r\n")
        if trimmedInput == "Q" {
            return
        }
        _, err = conn.Write([]byte(trimmedInput))
        if err != nil {
            return
        }
    }
}


在终端启动service  --窗口1
$ go run main.go 
start server...

然后新建终端启动client --窗口2 
$ go run main.go 
输入：
aaaa
bbbb
cccc

窗口1显示：
start server...
aaaa
bbbb
cccc