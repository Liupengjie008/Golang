在Go中http包的Server中，每一个请求在都有一个对应的goroutine去处理。请求处理函数通常会启动额外的goroutine用来访问后端服务，比如数据库和RPC服务。用来处理一个请求的goroutine通常需要访问一些与请求特定的数据，比如终端用户的身份认证信息，验证相关的令牌，请求的截止时间。然后系统才能释放这些goroutine占用的资源。

在Google内部，开发了Context包，专门用来简化对于处理单个请求的多个goroutine之间与请求域的数据，取消信号，截止时间等相关操作，这些操作可能涉及多个API调用。 go get golang.org/x/net/context命令获取这个包。本文要讲的就是如果使用这个包，同时也会提供一个完整的例子。

注意: 使用时遵循context规则

1. 不要将 Context 放入结构体，Context应该作为第一个参数传
   入，命名为ctx。
2. 即使函数允许，也不要传入nil的 Context。如果不知道用哪种 
   Context，可以使用context.TODO()。
3. 使用context的Value相关方法,只应该用于在程序和接口中传递
   和请求相关数据，不能用它来传递一些可选的参数
4. 相同的 Context 可以传递给在不同的goroutine；Context 是
   并发安全的。



Context结构
type Context interface {
    Deadline() (deadline time.Time, ok bool)
    Done() <-chan struct{}
    Err() error
    Value(key interface{}) interface{}
}

Deadline()  返回一个time.Time，是当前 Context 的应该结束的时间，ok 表示是否有 deadline
Done()  返回一个struct{}类型的只读 channel
Err()   返回 Context 被取消时的错误
Value(key interface{})  是 Context 自带的 K-V 存储功能

// Deadline会返回一个超时时间，Goroutine获得了超时时间后，例如可以对某些io操作设定超时时间。

// Done方法返回一个信道（channel），当Context被撤销或过期时，该信道是关闭的，即它是一个表示Context是否已关闭的信号。

// 当Done信道关闭后，Err方法表明Context被撤的原因。

// Value可以让Goroutine共享一些数据，当然获得数据是协程安全的。但使用这些数据的时候要注意同步，比如返回了一个map，而这个map的读写则要加锁。

网络请求超时控制
package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "time"
)
type Result struct {
    r   *http.Response
    err error
}
func process() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    //释放资源
    defer cancel()
    tr := &http.Transport{}
    client := &http.Client{Transport: tr}
    resultChan := make(chan Result, 1)
    //发起请求
    req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
    // req, err := http.NewRequest("GET", "http://www.google.com", nil)
    if err != nil {
        fmt.Println("http request failed, err:", err)
        return
    }
    /*
    func (c *Client) Do(req *Request) (*Response, error)
    */
    go func() {
        resp, err := client.Do(req)
        pack := Result{r: resp, err: err}
        //将返回信息写入管道(正确或者错误的)
        resultChan <- pack
    }()
    select {
    case <-ctx.Done():
        tr.CancelRequest(req)
        er:= <-resultChan
        fmt.Println("Timeout!",er.err)
    case res := <-resultChan:
        defer res.r.Body.Close()
        out, _ := ioutil.ReadAll(res.r.Body)
        fmt.Printf("Server Response: %s", out)
    }
    return
}
func main() {
    process()
}

如果修改下代码：
req, err := http.NewRequest("GET", "http://google.com", nil)

请求超时，输出日志信息如下：
Timeout! Get http://www.google.com: net/http: request canceled while waiting for connection


WithValue 传递元数据：

package main

import (
    "context"
    "fmt"
)

func process(ctx context.Context) {
    ret,ok := ctx.Value("trace_id").(int)
    if !ok {
        ret = 21342423
    }

    fmt.Printf("ret:%d\n", ret)

    s , _ := ctx.Value("session").(string)
    fmt.Printf("session:%s\n", s)
}

func main() {
    ctx := context.WithValue(context.Background(), "trace_id", 13483434)
    ctx = context.WithValue(ctx, "session", "sdlkfjkaslfsalfsafjalskfj")
    process(ctx)
}

输出结果：
ret:13483434
session:sdlkfjkaslfsalfsafjalskfj

通过Context我们也可以传递一些必须的元数据，这些数据会附加在Context上以供使用。

package main

import (
	"context"
	"fmt"
	"time"
)

var key string = "name"

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	//附加值
	valueCtx := context.WithValue(ctx, key, "【监控1】")
	go watch(valueCtx)
	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}
func watch(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			//取出值
			fmt.Println(ctx.Value(key), "监控退出，停止了...")
			return
		default:
			//取出值
			fmt.Println(ctx.Value(key), "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}


输出结果：
【监控1】 goroutine监控中...
【监控1】 goroutine监控中...
【监控1】 goroutine监控中...
【监控1】 goroutine监控中...
【监控1】 goroutine监控中...
可以了，通知监控停止
【监控1】 监控退出，停止了...




超时控制 WithDeadline
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(4 * time.Second)
	// d := time.Now().Add(2 * time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}

}
输出结果：
overslept

如果将上面代码修改为：
d := time.Now().Add(2 * time.Second)

输出结果：
context deadline exceeded


WithCancel
我们来了解一个利用context结束goroutine的demo

package main

import (
	"context"
	"fmt"
	"time"
)

/*
 创建一个管道chan，启动goroutine
 for循环存数据
**/
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				//执行defer cancel操作后，就会执行到该select入库
				fmt.Println("i exited")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func test() {

	ctx, cancel := context.WithCancel(context.Background())
	//当取数据n == 5时候，执行defer cancel操作
	defer cancel()
	intChan := gen(ctx)
	for n := range intChan {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
func main() {
	test()
	time.Sleep(time.Second * 5)
}
输出结果：
1
2
3
4
5
i exited









