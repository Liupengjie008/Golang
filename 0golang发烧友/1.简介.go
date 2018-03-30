定时器的使用
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.NewTicker(time.Second)
    for v := range t.C {
        fmt.Println("hello, ", v)
    }
}
如果在定时器到期之前，使用Stop()，那么就不会再有元素写入通道内，那么等待接受该通道元素所在的goroutine将被阻塞，恢复被停止的定时器的唯一途径是使用Reset()方法重置；定时器可以复用，尤其是在for循环中复用可以减少程序的资源占用，这时需要Reset()方法来重置定时器。

一次定时器    
package main

import (
    "fmt"
    "time"
)

func main() {
    select {
        Case <- time.After(time.Second):
         fmt.Println(“after”)
    }
}