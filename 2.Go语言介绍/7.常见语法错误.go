常见语法错误

1.多余的 import

新建文件，将如下内容拷贝并执行

ErrProg1.go

package main

import "fmt"
import "os" //excessive - we are not using any function in this package

func main() {
 	fmt.Println("Hello world")
}

输出为：

prog.go:4: imported and not used: os

Go编译器对Go程序非常严格，如果你不使用，就不要有多余的请求。在上面的代码中，试图引入os包，但是在代码当中，并没有使用，Go编译器对这样的行为严厉禁止。去掉第四行代码之后，程序就能正确编译运行。

2.命名区分大小写

ErrProg2.go

package main

import "fmt"

func main() {
 	fmt.println("Hello world")
}

输出为：

prog.go:6: cannot refer to unexported name fmt.println
prog.go:6: undefined: fmt.println

上述代码中，打印函数写的是fmt.println不是之前所写的fmt.Println。Go语言区分大小写，所以在编程时，要严格按照定义的方式进行引用和调用。以下代码都是不正确的:

Package main
iMport "fmt"
import "Fmt"
Func main() {}
Fmt.Println
fmt.println

3.分号分行

如果你学过C、C++、Java、Perl等等，应该已经注意到Go（至少在前面的代码中）没有要求在语句的末尾添加分号。其实在Go语言中，会自动在一行的末尾添加分号。然而，如果在一样有两条表达式，需要用分号显示的进行分割。让我们举个栗子：

ErrProg3.go

package main

import "fmt"

func main() {
 	fmt.Println("Hello world") fmt.Println("Hi again")
}

输出：

prog.go:6: syntax error: unexpected name, expecting semicolon or newline or }

解决以上问题，可以将上述的两条语句放在两行

部分代码

func main() {
 	fmt.Println("Hello world")
    fmt.Println("Hi again")
}

为了说明分号的作用，我们使用以下方式进行修改

package main

import "fmt"

func main() {
 fmt.Println("Hello world"); fmt.Println("Hi again")
}
输出：

Hello world
Hi again

因此在Go语言中，分号能省则省，如果必须使用时，添加上也不会出错。所以，如下代码也是正确滴。

ErrProg4.go

package main;

import "fmt";

func main() {
 fmt.Println("Hello world"); fmt.Println("Hi again");
};

输出为：

Hello world
Hi again

但是也请大家注意这些自动产生的分号。

无效的分号

继续修改上述代码

package main

import "fmt";;

func main() {
 fmt.Println("Hello world")
}

输出为：

prog.go:3: empty top-level declaration

上述问题得原因出现在import后面的第二个分号，第一个分号是正确的，上面的代码已经验证过。第二个分号之前没有任何有效的表达式，所以编译器报了上述的错误。去掉多余的分号，程序可正确运行。

语法和其它问题

每种语言都有各自的语法要求，Go编译器也不例外。很多时候我们会犯一些语法错误，以下会列出一些，以供大家参考。

译者注：以下错误输出就不在翻译了，保持错误输出原汁原味更好，这对大家应该不是问题)
package 'main' //ERROR - no quotes for the package name: package main
package "main" //ERROR - no quotes for the package: package main

package main.x  //ERROR - packages names in go are just one expression.  So either package main or package x.
package main/x  //ERROR - packages names in go are just one expression.  So either package main or package x.

import 'fmt' //ERROR - needs double quotes "fmt"
import fmt //ERROR - needs double quotes "fmt"

func main { } //ERROR - functions have to be followed by parantheses: func main() {}

func main() [] //ERROR - where curly braces are required, only those are allowed.  They are used to contain blocks of code.  func main() {}

func main() { fmt.Println('hello world') } //ERROR - use double quotes for strings: func main() { fmt.Println("hello world") }