package main

//Go 提供了一个 flag 包，支持基本的命令行标志解析。我们将用这个包来实现我们的命令行程序示例。
import "flag"
import "fmt"

func main() {
    //基本的标记声明仅支持字符串、整数和布尔值选项。这里我们声明一个默认值为 "foo" 的字符串标志 word并带有一个简短的描述。这里的 flag.String 函数返回一个字符串指针（不是一个字符串值），在下面我们会看到是如何使用这个指针的。
    wordPtr := flag.String("word", "foo", "a string")
    //使用和声明 word 标志相同的方法来声明 numb 和 fork 标志。
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")
    //用程序中已有的参数来声明一个标志也是可以的。注意在标志声明函数中需要使用该参数的指针。
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    //所有标志都声明完成以后，调用 flag.Parse() 来执行命令行解析。
    flag.Parse()
    //这里我们将仅输出解析的选项以及后面的位置参数。注意，我们需要使用类似 *wordPtr 这样的语法来对指针解引用，从而得到选项的实际值。
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

/*
测试这个程序前，最好将这个程序编译成二进制文件，然后再运行这个程序。
$ go build command-line-flags.go
word: opt
numb: 7
fork: true
svar: flag
tail: []
注意到，如果你省略一个标志，那么这个标志的值自动的设定为他的默认值。
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []
位置参数可以出现在任何标志后面。
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]
注意，flag 包需要所有的标志出现位置参数之前（否则，这个标志将会被解析为位置参数）。
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]
使用 -h 或者 --help 标志来得到自动生成的这个命令行程序的帮助文本。
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string
如果你提供一个没有使用 flag 包指定的标志，程序会输出一个错误信息，并再次显示帮助文本。
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...
*/