golang中sync包实现了两种锁Mutex （互斥锁）和RWMutex（读写锁），其中RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能。

1、互斥锁     （用于基本上全是写入操作的应用）

type Mutex

    func (m *Mutex) Lock()

    func (m *Mutex) Unlock()

go mutex是互斥锁，只有Lock和Unlock两个方法，在这两个方法之间的代码不能被多个goroutins同时调用到。

其中Mutex为互斥锁，Lock()加锁，Unlock()解锁，使用Lock()加锁后，便不能再次对其进行加锁，直到利用Unlock()解锁对其解锁后，才能再次加锁．适用于读写不确定场景，即读写次数没有明显的区别，并且只允许只有一个读或者写的场景，所以该锁叶叫做全局锁。

func (m *Mutex) Unlock()用于解锁m，如果在使用Unlock()前未加锁，就会引起一个运行错误．已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁。

互斥锁只能锁定一次，当在解锁之前再次进行加锁，便会死锁状态，如果在加锁前解锁，便会报错“panic: sync: unlock of unlocked mutex”


同一时刻只有一个携程在操作：

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]uint64)
	lock sync.Mutex    //互斥锁
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	fmt.Println(t.n, sum)
	// lock.Lock()
	m[t.n] = sum
	// lock.Unlock()
}

func main() {
	for i := 0; i < 16; i++ {
		t := &task{n: i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	// lock.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	// lock.Unlock()
}

输出结果：（有时也会正常执行）
$ go run main.go 
1 1
2 1
7 720
fatal error: concurrent map writes
fatal error: concurrent map writes

goroutine 7 [running]:
runtime.throw(0x10c80e7, 0x15)

package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m    = make(map[int]uint64)
	lock sync.Mutex
)

type task struct {
	n int
}

func calc(t *task) {
	var sum uint64
	sum = 1
	for i := 1; i < t.n; i++ {
		sum *= uint64(i)
	}

	fmt.Println(t.n, sum)
	lock.Lock()
	m[t.n] = sum
	lock.Unlock()
}

func main() {
	for i := 0; i < 16; i++ {
		t := &task{n: i}
		go calc(t)
	}

	time.Sleep(10 * time.Second)
	lock.Lock()
	for k, v := range m {
		fmt.Printf("%d! = %v\n", k, v)
	}
	lock.Unlock()
}


输出结果：

4 6
3 2
10 362880
5 24
2 1
8 5040
1 1
9 40320
0 1
6 120
13 479001600
14 6227020800
7 720
15 87178291200
12 39916800
11 3628800
2! = 1
6! = 120
15! = 87178291200
3! = 2
10! = 362880
1! = 1
4! = 6
8! = 5040
14! = 6227020800
11! = 3628800
9! = 40320
0! = 1
7! = 720
12! = 39916800
5! = 24
13! = 479001600

建议：同一个互斥锁的成对锁定和解锁操作放在同一层次的代码块中。

package main
 
import (
    "fmt"
    "sync"
    "time"
)
 
func main(){
    //声明
    var mutex sync.Mutex
    fmt.Println("Lock the lock. (G0)")
    //加锁mutex
    mutex.Lock()
 
    fmt.Println("The lock is locked.(G0)")
    for i := 1; i < 4; i++ {
        go func(i int) {
            fmt.Printf("Lock the lock.(G%d)\r\n", i)
            mutex.Lock()
            fmt.Printf("The lock is locked.(G%d)\r\n", i)
        }(i)
    }
    //休息一会,等待打印结果
    time.Sleep(time.Second)
    fmt.Println("Unlock the lock. (G0)")
    //解锁mutex
    mutex.Unlock()
 
    fmt.Println("The lock is unlocked. (G0)")
    //休息一会,等待打印结果
    time.Sleep(time.Second)
}

输出结果：

Lock the lock. (G0)
The lock is locked.(G0)
Lock the lock.(G3)
Lock the lock.(G2)
Lock the lock.(G1)
Unlock the lock. (G0)
The lock is unlocked. (G0)
The lock is locked.(G3)




