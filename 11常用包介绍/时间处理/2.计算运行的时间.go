函数time.Since()
计算golang运行的时间是非常有用的性能衡量指标，特别是在并发基准测试中。下面将介绍如何简单地使用Go语言来计算程序运行的时间。

简单地使用Golang的time.Since()函数即可。下面有一个完整例子展示这个用法。

time.Hour 表示1小时
time.Minute 表示1分钟
time.Second 表示1秒
time.Millisecond	表示1毫秒
time.Microsecond	表示1微妙
time.Nanosecond	表示1纳秒

// 休眠5秒
time.Sleep(5 * time.Second)

package main

import (
    "fmt"
    "time"
)

func StartCac() {
    t1 := time.Now()

	for i := 0; i < 10000; i++ {
		time.Sleep(2*time.Nanosecond) 	//休眠2纳秒
        // continue
    }
    elapsed := time.Since(t1)
    fmt.Println("App elapsed: ", elapsed)
}

func main(){
    StartCac()
}

输出结果：
App elapsed:  4.031139ms
