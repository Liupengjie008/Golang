WaitGroup
WaitGroup在go语言中，用于线程同步，单从字面意思理解，wait等待的意思，group组、团队的意思，WaitGroup就是指等待一组，等待一个系列执行完成后才会继续向下执行。

先说说WaitGroup的用途：它能够一直等到所有的goroutine执行完成，并且阻塞主线程的执行，直到所有的goroutine执行完成。

WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。

Add:添加或者减少等待goroutine的数量

Done:相当于Add(-1)

Wait:执行阻塞，直到所有的WaitGroup数量变成0

package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    wg := sync.WaitGroup{}

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go calc(&wg, i)
    }

    wg.Wait()
    fmt.Println("all goroutine finish")
}
func calc(w *sync.WaitGroup, i int) {

    fmt.Println("calc:", i)
    time.Sleep(time.Second)
    w.Done()
}


运行结果：
$ go run main.go 
calc: 1
calc: 5
calc: 7
calc: 8
calc: 2
calc: 4
calc: 6
calc: 0
calc: 3
calc: 9
all goroutine finish