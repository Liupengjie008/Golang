sync/atomic 库提供了原子操作的支持，原子操作直接有底层CPU硬件支持，因而一般要比基于操作系统API的锁方式效率高些。本文对 sync/atomic 中的基本操作进行一个简单的介绍。

原子增、减值
用于对变量值进行原子增操作，并返回增加后的值。

第一个参数值必须是一个指针类型的值，以便施加特殊的CPU指令。
第二个参数值的类型和第一个被操作值的类型总是相同的。

package main 

import(
    "fmt"
    "sync"
    "sync/atomic"
)

func main(){
    var sum uint32 = 100
    // var sum int32 = 100
    var wg sync.WaitGroup
    for i := 0; i < 50; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            //sum += 1                 //1
            atomic.AddUint32(&sum, 1)  //2
            // atomic.AddInt32(&sum, -1)  //2
        }()
    }
    wg.Wait()
    fmt.Println(sum)
}    

通过对比//1与//2的结果，可以很清楚的看到原子操作起到的作用。使用//1时，可以看到sum的值是不定的，取决于sum的同步访问情况；使用//2时，结果是确定而且正确的，同一时间只有一个goroutine修改sum。

函数原型：

    atomic.AddUint32(addr *uint32, delta uint32) uint32
    atomic.AddUint64(addr *uint64, delta uint64) uint64
    atomic.AddInt32(addr *int32, delta int32) int32
    atomic.AddInt64(addr *int64, delta int64) int64
    atomic.AddUintptr(addr *uintptr, delta uintptr) uintptr