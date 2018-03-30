读写参数
文件打开模式：

const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和O_CREATE配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
权限控制：

r ——> 004
w ——> 002
x ——> 001

读取文件
os.Open || os.OpenFile
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // file, err := os.Open("/tmp/test")
    file, err := os.OpenFile("/tmp/test", os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        fmt.Println("Open file error: ", err)
        return
    }
    defer file.Close()    //关闭文件

    reader := bufio.NewReader(file)    //带缓冲区的读写
    for {
        str, err := reader.ReadString('\n')    // 循环读取一行
        if err != nil {
            fmt.Println("read string failed, err: ", err)
            return
        }
        fmt.Println("read string is %s: ", str)
    }
}
 

readline 行读取
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
)

func main() {
    file, err := os.Open("C:/test.log")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()
    reader := bufio.NewReader(file)
    var line []byte
    for {
        data, prefix, err := reader.ReadLine()
        if err == io.EOF {
            break
        }

        line = append(line, data...)
        if !prefix {
            fmt.Printf("data:%s\n", string(line))
            line = line[:]
        }

    }
}
 

读取整个文件：
"io/ioutil" 包实现了读取整个文件功能

package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    fileName := "/tmp/test"

    file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0666)
    if err != nil {
        fmt.Println("Open file error: ", err)
        return
    }
    defer file.Close()

    buf, err := ioutil.ReadAll(file)
    //buf, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
        return
    }
    fmt.Printf("%s\n", string(buf))
}