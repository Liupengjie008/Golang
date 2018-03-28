Golang 在语言层面对并发编程提供支持，一种类似协程，称作 goroutine 的机制。

只需在函数调用语句前添加 go 关键字，就可创建并发执行单元。开发人员无需了解任何执行细节，调度器会自动将其安排到合适的系统线程上执行。goroutine 是一种非常轻量级的实现，可在单个进程里执行成千上万的并发任务。

事实上，入口函数 main 就以 goroutine 运行。另有与之配套的 channel 类型，用以实现 "以通讯来共享内存" 的 CSP 模式。

go func() {
    println("Hello, World!")
}()


调度器不能保证多个 goroutine 执行次序，且进程退出时不会等待它们结束。

默认情况下，进程启动后仅允许一个系统线程服务于 goroutine。可使用环境变量或标准库函数 runtime.GOMAXPROCS 修改，让调度器用多个线程实现多核并行，而不仅仅是并发。
 
package main

import (
    "math"
    "sync"
)

func sum(id int) {
    var x int64
    for i := 0; i < math.MaxUint32; i++ {
        x += int64(i)
    }

    println(id, x)
}

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    for i := 0; i < 2; i++ {
        go func(id int) {
            defer wg.Done()
            sum(id) 
        }(i)
    }
    
    wg.Wait() 
}

输出:

$ go build -o test

$ time -p ./test

0 9223372030412324865
1 9223372030412324865


real   7.70     // 程序开始到结束时间差 (  CPU 时间)
user   7.66     // 用户态所使用 CPU 时间片 (多核累加) 
sys    0.01     // 内核态所使用 CPU 时间片
 
$ GOMAXPROCS=2 time -p ./test


0 9223372030412324865
1 9223372030412324865

real 4.18
user 7.61       // 虽然总时间差不多，但由 2 个核并行，real 时间自然少了许多。 
sys 0.02


调用 runtime.Goexit 将立即终止当前 goroutine 执行，调度器确保所有已注册 defer 延迟调用被执行。

package main

import (
    "runtime"
    "sync"
)

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(1)

    go func() {
        defer wg.Done()
        defer println("A.defer")

        func() {
            defer println("B.defer")
            runtime.Goexit()     // 终止当前 goroutine
            println("B")          // 不会执行 
        }()

        println("A")            // 不会执行
    }()

    wg.Wait() 
}
 
输出:
B.defer
A.defer


和协程 yield 作用类似，Gosched 让出底层线程，将当前 goroutine 暂停，放回队列等待下次被调度执行。

package main

import (
    "runtime"
    "sync"
)

func main() {
    wg := new(sync.WaitGroup)
    wg.Add(2)

    go func() {
        defer wg.Done()

        for i := 0; i < 6; i++ {
            println(i)
            if i == 3 { runtime.Gosched() }
        }
    }()

    go func() {
        defer wg.Done()
        println("Hello, World!")
    }()
    
    wg.Wait() 
}

输出:
$ go run main.go
0
1
2
3
Hello, World!
4
5