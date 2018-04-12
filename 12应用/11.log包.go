import "log"

log包实现了简单的日志服务。本包定义了Logger类型，该类型提供了一些格式化输出的方法。本包也提供了一个预定义的“标准”Logger，可以通过辅助函数Print[f|ln]、Fatal[f|ln]和Panic[f|ln]访问，比手工创建一个Logger对象更容易使用。
Logger会打印每条日志信息的日期、时间，默认输出到标准错误。
Fatal系列函数会在写入日志信息后调用os.Exit(1)。
Panic系列函数会在写入日志信息后panic。


Print 示例：
package main

import "log"

func main() {
	log.Print("this is log print test", "\n")
	log.Printf("this is %s", "log printf test\n")
	log.Printf("this is log println test")
}

运行结果：
2018/04/12 14:03:37 this is log print test
2018/04/12 14:03:37 this is log printf test
2018/04/12 14:03:37 this is log println test

对于 log.Fatal 接口，会先将日志内容打印到标准输出，接着调用系统的 os.exit(1) 接口，退出程序并返回状态 1 。
但是有一点需要注意，由于是直接调用系统接口退出，defer函数不会被调用。

Fatal 示例：
package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("first defer")
	}()
	log.Fatal("this is log fatal test", "\n")
	log.Fatalf("this is %s", "log fatalf test\n")
	log.Fatalln("this is log fatalln test\n")
	defer func() {
		fmt.Println("second defer")
	}()
}

运行结果：
2018/04/12 14:12:01 this is log fatal test
exit status 1

Panic 示例：
package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("first defer")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	log.Panic("this is log panic")
	defer func() {
		fmt.Println("second defer")
	}()
}

运行结果：
2018/04/12 14:36:02 this is log panic
first defer
this is log panic





func New(out io.Writer, prefix string, flag int) *Logger
/*
New创建一个Logger。参数out设置日志信息写入的目的地。
参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。
*/
func (l *Logger) SetPrefix(prefix string)
// SetPrefix设置logger的输出前缀。
func (l *Logger) SetFlags(flag int)
// SetFlags设置logger的输出选项。
const (
    // 字位共同控制输出日志信息的细节。不能控制输出的顺序和格式。
    // 在所有项目后会有一个冒号：2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
)
// 这些选项定义Logger类型如何生成用于每条日志的前缀文本。

代码实力：

package main

import (
	"fmt"
	"log"
	"os"
)

func Debug(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Debug]", log.Ldate)

	debugLog.SetPrefix("[Debug]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Debug log")
}
func Waring(logName string) {
	logFile, err := os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Printf("create ./test.log err : %v\n", err)
	}
	if logFile != nil {
		defer func(file *os.File) { file.Close() }(logFile)
	}

	debugLog := log.New(logFile, "[Waring]", log.Ldate)

	debugLog.SetPrefix("[Waring]")
	debugLog.SetFlags(log.Lshortfile)
	debugLog.Println("this is Waring log")
}

func main() {
	logName := "./test.log"
	Debug(logName)
	Waring(logName)
}



