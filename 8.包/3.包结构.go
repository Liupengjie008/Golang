Golang 包结构 ：所有代码都必须组织在 package 中。

• 源 件头部以 "package <name>" 声明包名称。
• 包由同  录下的多个源码 件组成。
• 包名类似 namespace，与包所在 录名、编译 件名 关。 •  录名最好不  main、all、std 这三个保留名称。
• 可执  件必须包含 package main，  函数 main。
包中成员以名称 字   写决定访问权限。 • public:  字  写，可被包外访问。
• internal:  字  写，仅包内成员可以访问。 该规则适 于全局变量、全局常量、类型、结构字段、函数、 法等。
8.3.1 导 包
使 包成员前，必须先  import 关键字导 ，但不能形成导 循环。 import "相对 录/包主 件名"
相对 录是指从 <workspace>/pkg/<os_arch> 开始的  录，以标准库为例:
在导 时，可指定包成员访问 式。 如对包重命名，以避免同名冲突。
Go 学习笔记, 第 4 版
 说明:os.Args 返回命令 参数，os.Exit 终 进程。 要获取正确的可执  件路径，可  filepath.Abs(exec.LookPath(os.Args[0]))。
  import "fmt"      ->  /usr/local/go/pkg/darwin_amd64/fmt.a
import "os/exec"  ->  /usr/local/go/pkg/darwin_amd64/os/exec.a
 import     "yuhen/test"
import  M  "yuhen/test"
import  .  "yuhen/test"
import  _  "yuhen/test"
// 默认模式: test.A
// 包重命名: M.A
// 简便模式: A
//  导 模式: 仅让该包执 初始化函数。
77

 未使 的导 包，会被编译器视为错误 (不包括 "import _")。 ./main.go:4: imported and not used: "fmt"
对于当前 录下的 包，除使 默认完整导 路径外，还可使  local  式。
main.go
8.3.2  定义路径
可通过 meta 设置为代码库设置 定义路径。
server.go
Go 学习笔记, 第 4 版
  workspace
|+--- src
|+--- learn
|+--- main.go |+--- test
|+--- test.go
 import "learn/test" // 正常模式
import "./test" // 本地模式，仅对 go run main.go 有效。
 package main
import (
    "fmt"
"net/http" )
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, `<meta name="go-import"
                   content="test.com/qyuhen/test git https://github.com/qyuhen/test">`)
}
func main() {
    http.HandleFunc("/qyuhen/test", handler)
    http.ListenAndServe(":80", nil)
}
该 例使  定义域名 test.com 重定向到 github。
78

 Go 学习笔记, 第 4 版
 $ go get -v test.com/qyuhen/test
Fetching https://test.com/qyuhen/test?go-get=1
https fetch failed.
Fetching http://test.com/qyuhen/test?go-get=1
Parsing meta tags from http://test.com/qyuhen/test?go-get=1 (status code 200)
get "test.com/qyuhen/test": found meta tag http://test.com/qyuhen/test?go-get=1
test.com/qyuhen/test (download)
test.com/qyuhen/test
如此，该库就有两个有效导 路径，可能会导致存储两个本地副本。为此，可以给库添加 专 的 "import comment"。当 go get 下载完成后，会检查本地存储路径和该注释是否  致。
github.com/qyuhen/test/abc.go
如继续  github 路径，会导致 go build 失败。
 package test // import "test.com/qyuhen/test"
func Hello() {
    println("Hello, Custom import path!")
}
 $ go get -v github.com/qyuhen/test
github.com/qyuhen/test (download)
package github.com/qyuhen/test
  imports github.com/qyuhen/test
  imports github.com/qyuhen/test: expects import "test.com/qyuhen/test"
这就强制包 户使 唯 路径，也便于 后将包迁移到其他位置。 资源:Go 1.4 Custom Import Path Checking
8.3.3 初始化 初始化函数:
• 每个源 件都可以定义 个或多个初始化函数。 • 编译器不保证多个初始化函数执 次序。
 79

 • 初始化函数在单 线程被调 ，仅执  次。 • 初始化函数在包所有全局变量初始化后执 。 • 在所有初始化函数结束后才执  main.main。 •  法调 初始化函数。
因为 法保证初始化函数执 顺序，因此全局变量应该直接  var 初始化。
Go 学习笔记, 第 4 版
 var now = time.Now()
func init() {
    fmt.Printf("now: %v\n", now)
}
func init() {
    fmt.Printf("since: %v\n", time.Now().Sub(now))
}
可在初始化函数中使  goroutine，可等待其结束。
 var now = time.Now()
func main() {
    fmt.Println("main:", int(time.Now().Sub(now).Seconds()))
}
func init() {
    fmt.Println("init:", int(time.Now().Sub(now).Seconds()))
    w := make(chan bool)
    go func() {
        time.Sleep(time.Second * 3)
        w <- true
}()
<-w }
输出:
不应该滥 初始化函数，仅适合完成当前 件中的相关环境设置。
 init: 0
main: 3