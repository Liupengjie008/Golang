编写 Hello World

     创建文件 program1.go，不写入任何内容。按照如下的命令尝试进行编译

     cmd-prompt> go run program1.go

     将会打印出如下错误：

     package :
     program1.go:1:1: expected 'package', found 'EOF'

     在Go语言中，所有文件必须隶属于某一个包。当前，只需要理解在文件的头部声明一个package name就可以了，其中package为关键字，name为你自己起的一个包名字。
     在大型的程序中，包可以很好的有条理的组织各种功能。
     例如，如果你想写一个关于交通工具的虚拟模型，你应该把所有属于car的模型放入一个叫做cars的包中，把所有属于bus的模型放入buses的包中。
     组织相关的功能只是包的一种用途，在后续文章中会讲述更多内容。

     现在让我们在刚刚创建的文件中添加一条语句，之后重新执行命令

     program1.go内容

     package main

     执行之前的命令之后，会打印如下错误：

     runtime.main: undefined: main.main

     Go程序启动时，需要在文件中有一个可标识入口。就像汽车必须有一把启动点火的钥匙、电脑需要有一个开机键，Go程序中需要有一个main函数。

     在文件中添加另外一行，并且重试

     program1.go内容

     package main

     func main() {}

     执行命令go run program1.go

     程序正确执行，但是由于我们没有做任何其它操作，程序很快就退出了。

     恭喜恭喜，到目前为止，我们已经创建了自己的第一个程序。虽然没啥卵用，但是已经可以正常运行了。

     让我们继续添加一行

     program1.go内容

     package main

     func main() {
           Println("Hello world")
     }

     尝试运行，将会打印如下错误

     program1.go:4: undefined: Println

     Println是向屏幕输入内容。执行命令之后，编译器报未定义。为什么呢？可记得前面提到的包？对喽，这里我们就需要用到包了。像Println这样的函数存放在某些包中。然而，当前这些包由于我们没有主动引入，但不能使用滴。如果我们需要使用这些包中的功能，首先需要import它们。这就像我们从海外购入汽车一样一样的。Ok，让我们import试试。

     函数Println和其它读写文本和字符的函数，都存放在一个叫做fmt的包中——formatting的缩写

     Go程序语言遵从短小精悍。如果你写Java代码，你会以一个很长的命名方式进行定义。比如上面的formatting package会很正常的被叫做formatting。但是在Go语言中，就要打破这种常规，追求精简。开始的时候，并不能适应这种模式，但是使用过一段时间之后，真尼玛太棒了。代码变得简洁了，阅读起来更快了，奇怪的是也没有降低可读性。个人观点哦~~~
至此，让我们在添加几行代码

     package main

     import "fmt"

     func main() {
           fmt.Println("Hello world")
     }

     运行程序go run program1.go，输出如下：

     Hello world

     哇哦，是不是很神奇，我们只是在package下面添加了一个import语句，第一个Go程序已经正常运行了。import之后，Println可以通过 包名.的方式进行调用。知道了吧，就这么简单。