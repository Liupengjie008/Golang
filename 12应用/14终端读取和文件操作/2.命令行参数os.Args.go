package main

import "os"
import "fmt"

func main() {
    //os.Args 提供原始命令行参数访问功能。注意，切片中的第一个参数是该程序的路径，并且 os.Args[1:]保存所有程序的的参数。
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    //你可以使用标准的索引位置方式取得单个参数的值。
    arg := os.Args[3]
    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}

/*
要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件

$ $ go build main.go 
$ ./main a b c d

输出结果：
[./main a b c d]
[a b c d]
c

*/