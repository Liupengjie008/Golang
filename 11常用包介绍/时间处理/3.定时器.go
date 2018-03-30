定时器的使用
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.NewTicker(time.Second*3)  //3三秒执行一次
    // t := time.NewTicker(执行周期)    
    for v := range t.C {
        fmt.Println("hello, ", v)
    }
}

输出结果：
$ go run main.go 
hello,  2018-03-30 16:08:35.930617 +0800 CST m=+3.005158519
hello,  2018-03-30 16:08:38.930654 +0800 CST m=+6.005111785
hello,  2018-03-30 16:08:41.930753 +0800 CST m=+9.005128262
hello,  2018-03-30 16:08:44.930819 +0800 CST m=+12.005110030
hello,  2018-03-30 16:08:47.930948 +0800 CST m=+15.005155812

如果在定时器到期之前，使用Stop()，那么就不会再有元素写入通道内，那么等待接受该通道元素所在的goroutine将被阻塞，恢复被停止的定时器的唯一途径是使用Reset()方法重置；定时器可以复用，尤其是在for循环中复用可以减少程序的资源占用，这时需要Reset()方法来重置定时器。

一次定时器：

package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    select {
        case <- time.After(time.Second*5):  // 5秒后执行
        // case <- time.After(周期):
            fmt.Println("after")
    }
    duration := time.Since(start)
    fmt.Println("运行时间：",duration)
}


输出结果：
$ go run main.go 
after
运行时间： 5.003661304s