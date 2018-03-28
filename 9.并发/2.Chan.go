channel存取

先进先出 
线程安全，多个goroutine同时访问，不需要加锁 
channel是有类型的，一个整数的channel只能存放整数

package main

import (
    "fmt"
    // "fmt"

)
type Stu struct{
    name string
}

func main() {
    //int类型
    var intChan chan int 
    intChan = make(chan int,10)
    intChan <- 10

    a := <- intChan
    fmt.Println(a)
    fmt.Println("--------/map类型---------")
    //map类型
    var mapChan chan map[string]string
    mapChan = make(chan map[string]string,10)
    m:= make(map[string]string,16)
    m["stu01"] = "001"
    m["stu02"] = "002"
    m["stu03"] = "003"
    mapChan <- m

    b := <- mapChan
    fmt.Println(b)
    fmt.Println("--------结构体---------")
    //结构体
    var stuChan chan Stu
    stuChan = make(chan Stu,10)
    stu:= Stu{
        name:"safly",
    }

    stuChan <- stu
    tempStu := <- stuChan
    fmt.Println(tempStu)


    fmt.Println("--------结构体内存地址---------")
    //结构体内存地址值
    var stuChanId chan *Stu
    stuChanId = make(chan *Stu,10)
    stuId := &Stu{
        name:"safly",
    }

    stuChanId <- stuId
    tempStuId := <- stuChanId
    fmt.Println(tempStuId)
    fmt.Println(*(tempStuId))

    fmt.Println("-----------接口---------")
    //接口
    var StuInterChain chan interface{}
    StuInterChain = make(chan interface{},10)
    stuInit:= Stu{
        name:"safly",
    }
    //存
    StuInterChain <- &stuInit
    //取
    mFetchStu:= <- StuInterChain
    fmt.Println(mFetchStu)

    //转
    var mStuConvert *Stu
    mStuConvert,ok := mFetchStu.(*Stu)
    if !ok{
        fmt.Println("cannot convert")
        return
    }
    fmt.Println(mStuConvert)
    fmt.Println(*(mStuConvert))

}

输出结果：
10
--------/map类型---------
map[stu03:003 stu02:002 stu01:001]
--------结构体---------
{safly}
--------结构体内存地址---------
&{safly}
{safly}
-----------接口---------
&{safly}
&{safly}
{safly}


package main

import (
    "time"
    "fmt"
)

func main() {
    intChan := make(chan int,5)
    go write(intChan)
    go read(intChan)

    time.Sleep(10 *time.Second)
}

/*
由于管道容量是5，开启go写入10个数据，再写入5个数据，
会阻塞，
然而read每秒会读取一个，然后在会写入一个数据

*/
func write(ch chan int){
    for i:= 0;i<10;i++{
        ch <- i
        fmt.Println("put data:",i)
    }
}

func read(ch chan int){
    for{
        var b int
        b = <- ch
        fmt.Println(b)
        time.Sleep(time.Second)
    }
}


输出结果：
0
put data: 0
put data: 1
put data: 2
put data: 3
put data: 4
put data: 5
1
put data: 6
2
put data: 7
3
put data: 8
4
put data: 9
5
6
7
8
9


channel关闭
channel关闭后，就不能取出数据了

package main

import "fmt"

func main() {
    var ch chan int
    ch = make(chan int, 5)

    for i := 0; i < 5; i++ {
        ch <- i
    }

    close(ch)
    for {
        var b int
        b, ok := <-ch
        if ok == false {
            fmt.Println("chan is close")
            break
        }
        fmt.Println(b)
    }
}

如果将close(ch)注释掉，意思是不关闭管道，那么会出现dead lock死锁 
因为存入管道5个数字，然后无限取数据，会出现死锁

输出结果：
0
1
2
3
4
chan is close



range 遍历 chan

package main

import "fmt"

func main() {
    var ch chan int
    ch = make(chan int, 10)

    for i := 0; i < 10; i++ {
        ch <- i
    }

    close(ch)
    for v := range ch {
        fmt.Println(v)
    }
}

同样如果将close(ch)注释掉，意思是不关闭管道，那么会出现dead lock死锁 
因为存入管道10个数字，然后无限取数据，在取出来第10个数据，在次range管道，会dead lock

输出结果：
0
1
2
3
4
5
6
7
8
9