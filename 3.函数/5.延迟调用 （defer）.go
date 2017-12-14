Golang延迟调：
关键字 defer 用于注册延迟调用。这些调用直到 return 前才被执 ，通常 于释放资源或错误处理。

func test() error {
    f, err := os.Create("test.txt")
    if err != nil { return err }

    defer f.Close()     // 注册调用，而不是注册函数。必须提供参数，哪怕为空。
    
    f.WriteString("Hello, World!")
    return nil 
}

多个 defer 注册，按 FILO 次序执行 ( 先进后出 )。哪怕函数或某个延迟调用发生错误，这些调用依旧会被执行。

func test(x int) {
    defer println("a")
    defer println("b")

    defer func() {
        println(100 / x)    // div0 异常未被捕获，逐步往外传递，最终终止进程。
    }()

    defer println("c")
}

func main() {
    test(0)
}

输出:
c
b
a
panic: runtime error: integer divide by zero


*延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
 
func test() {
    x, y := 10, 20

    defer func(i int) {
        println("defer:", i, y) // y 闭包引用
    }(x) // x 被复制
     
    x += 10
    y += 100
    println("x =", x, "y =", y)
}

输出:
x = 20 y = 120
defer: 10 120


*滥用 defer 可能会导致性能问题，尤其是在一个 "大循环" 里。
 
var lock sync.Mutex

func test() {
    lock.Lock()
    lock.Unlock()
}

func testdefer() {
    lock.Lock()
    defer lock.Unlock()
}

func BenchmarkTest(b *testing.B) {
    for i := 0; i < b.N; i++ {
        test() 
    }
}

func BenchmarkTestDefer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        testdefer()
    }
}

输出:
BenchmarkTest       50000000    43 ns/op 
BenchmarkTestDefer  20000000    128 ns/op