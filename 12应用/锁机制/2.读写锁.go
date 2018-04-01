2、读写锁     （用于读取多，写入少的操作应用）

type RWMutex

    func (rw *RWMutex) Lock()

    func (rw *RWMutex) RLock()

    func (rw *RWMutex) RLocker() Locker

    func (rw *RWMutex) RUnlock()

    func (rw *RWMutex) Unlock()

RWMutex是一个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景．

func (rw *RWMutex) Lock()　　写锁，如果在添加写锁之前已经有其他的读锁和写锁，则lock就会阻塞直到该锁可用，为确保该锁最终可用，已阻塞的 Lock 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
func (rw *RWMutex) Unlock()　写锁解锁，如果没有进行写锁定，则就会引起一个运行时错误

func (rw *RWMutex) RLock() 读锁，当有写锁时，无法加载读锁，当只有读锁或者没有锁时，可以加载读锁，读锁可以加载多个，所以适用于＂读多写少＂的场景

func (rw *RWMutex)RUnlock()　读锁解锁，RUnlock 撤销单次RLock 调用，它对于其它同时存在的读取器则没有效果。若 rw 并没有为读取而锁定，调用 RUnlock 就会引发一个运行时错误(注：这种说法在go1.3版本中是不对的，例如下面这个例子)。

读写锁的写锁只能锁定一次，解锁前不能多次锁定，读锁可以多次，但读解锁次数最多只能比读锁次数多一次，一般情况下我们不建议读解锁次数多余读锁次数。

读多写少的情况，用读写锁， 携程同时在操作读。

package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

//读写锁
var rwLock sync.RWMutex

func testRWLock() {
    var a map[int]int
    a = make(map[int]int, 5)
    a[8] = 10
    a[3] = 10
    a[2] = 10
    a[1] = 10
    a[18] = 10
    for i := 0; i < 2; i++ {
        go func(b map[int]int) {
            rwLock.Lock()
            b[8] = rand.Intn(100)
            rwLock.Unlock()
        }(a)
    }
    for i := 0; i < 10; i++ {
        go func(b map[int]int) {
            rwLock.RLock() //读锁
            fmt.Println(a)
            rwLock.RUnlock()
        }(a)
    }
    time.Sleep(time.Second * 2)

}
func main() {
    
    testRWLock()
    //读多写少的时候，用读写锁
}

输出结果：
map[3:10 2:10 1:10 18:10 8:81]
map[8:87 3:10 2:10 1:10 18:10]
map[2:10 1:10 18:10 8:87 3:10]
map[2:10 1:10 18:10 8:87 3:10]
map[8:87 3:10 2:10 1:10 18:10]
map[8:87 3:10 2:10 1:10 18:10]
map[2:10 1:10 18:10 8:87 3:10]
map[8:87 3:10 2:10 1:10 18:10]
map[1:10 18:10 8:87 3:10 2:10]
map[8:87 3:10 2:10 1:10 18:10]