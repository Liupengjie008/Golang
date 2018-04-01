package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

//读写锁
var rwLock sync.RWMutex
var lock sync.Mutex

func testRWLock() {
    var a map[int]int
    a = make(map[int]int, 5)
    var count int32
    a[8] = 10
    a[3] = 10
    a[2] = 10
    a[1] = 10
    a[18] = 10
    for i := 0; i < 2; i++ {
        go func(b map[int]int) {
            //rwLock.Lock() //读写锁的代码
            lock.Lock() //互斥锁的代码
            b[8] = rand.Intn(100)
            time.Sleep(10 * time.Microsecond) //微妙
            //rwLock.Unlock()
            lock.Unlock()

        }(a)
    }
    for i := 0; i < 100; i++ {
        go func(b map[int]int) {
            for {
                //rwLock.RLock() //读写锁的代码
                lock.Lock()
                time.Sleep(time.Millisecond)
                //rwLock.RUnlock()
                lock.Unlock()
                atomic.AddInt32(&count, 1)
            }
        }(a)
    }
    time.Sleep(time.Second * 20)
    fmt.Println(atomic.LoadInt32(&count))
}
func main() {
    //互斥锁
    testRWLock()
    //读多写少的时候，用读写锁
}

输出结果：
15589

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)

//读写锁
var rwLock sync.RWMutex
var lock sync.Mutex

func testRWLock() {
    var a map[int]int
    a = make(map[int]int, 5)
    var count int32
    a[8] = 10
    a[3] = 10
    a[2] = 10
    a[1] = 10
    a[18] = 10
    for i := 0; i < 2; i++ {
        go func(b map[int]int) {
            rwLock.Lock() //读写锁的代码
            // lock.Lock() //互斥锁的代码
            b[8] = rand.Intn(100)
            time.Sleep(10 * time.Microsecond) //微妙
            rwLock.Unlock()
            // lock.Unlock()

        }(a)
    }
    for i := 0; i < 100; i++ {
        go func(b map[int]int) {
            for {
                rwLock.RLock() //读写锁的代码
                // lock.Lock()
                time.Sleep(time.Millisecond)
                rwLock.RUnlock()
                // lock.Unlock()
                atomic.AddInt32(&count, 1)
            }
        }(a)
    }
    time.Sleep(time.Second * 20)
    fmt.Println(atomic.LoadInt32(&count))
}
func main() {
    //互斥锁
    testRWLock()
    //读多写少的时候，用读写锁
}

输出结果：
1525002