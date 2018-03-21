包的概念
每个 Go 程序都是由包组成的。
 
程序运行的入口是包 main 。

下面这个程序使用并导入了包 "fmt" 和 "math/rand" 。

按照惯例，包名与导入路径的最后一个目录一致。例如，"math/rand" 包由 package rand 语句开始。

注意：这个程序的运行环境是确定性的，因此 rand.Intn 每次都会返回相同的数字。 （为了得到不同的随机数，需要提供一个随机数种子，参阅 rand.Seed。）

package main

import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.Println("My favorite number is", rand.Intn(10))
} 

导入
这个代码用圆括号组合了导入，这是“打包”导入语句。

同样可以编写多个导入语句，例如：

import "fmt"
import "math"
不过使用打包的导入语句是更好的形式。

import (
    "fmt"
    "math"
)

导出名
在 Go 中，首字母大写的名称是被导出的。

在导入包之后，你只能访问包所导出的名字，任何未导出的名字是不能被包外的代码访问的。

Foo 和 FOO 都是被导出的名称。名称 foo 是不会被导出的。

包别名的
    应用：
    package main

    import(
        f “fat”
    )

    func main(){
        f.Println(“包的别名”)
    }


package 基本的管理单元
	同一个package下面，可以有非常多的不同文件，只要 每个文件的头部    都有 如 "package xxx" 的相同name
	就可以 在主方法中使用 xxx.Method()调用不同文件中的方法了。
	文件夹名字可以和这个package 名称不一致，
	比如我有个文件夹名字是mypackage,其中包含了a.go,b.go, c.go三个文件 :
	mypackage	
	  | --a.go
	  | --b.go
	  | --c.go
	
	比如a.go中有 Saya(),b.go中有Sayb()  而几个文件共同的package name 确是testpackage
	
	所以在 主函数中调用a.go 和b.go文件中的各自方法只要用，testpackage.Saya() ,testpackage.Sayb()即可
	还有默认的init方法，在import进来的时候就去执行了，而且允许每个文件中都有init()这个方法，当然是每个都会执行。
	
	
包的导入语法：import 命令用来导入包文件。
	go import 用法
	import "fmt"		最常用的一种形式（系统包）
	import "./test"		导入同一目录下test包中的内容（相对路径）
	import “shorturl/model” 	加载gopath/src/shorturl/model模块（绝对路径）
	import f "fmt"		导入fmt，并给他启别名ｆ
	import . "fmt" 		将fmt启用别名"."，这样就可以直接使用其内容，而不用再添加fmt，如fmt.Println可以直接写成Println
	import  _ "fmt" 	表示不使用该包，而是只是使用该包的init函数，并不显示的使用该包的其他内容。注意：这种形式的import，当import时就执行了fmt包中的init函数，而不能够使用该包的其他函数。