Golang是静态类型语言，不能在运行期间改变变量类型。
变量
var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。

初始化变量
使用关键字 var 定义变量，自动初始化为零值。如果提供初始化值，可省略变量类型，由编译器自动推断。
var x int
var f float32 = 1.6
var s = "abc"

一次性定义多个变量。

var x, y, z int
var s, n = "abc", 123
var (
	a int
	b float32 
)
func main() {
    n, s := 0x1234, "Hello, World!"
    println(x, s, n)
}

多变量赋值时，先计算所有相关值，然后再从左到右依次赋值。

data, i := [3]int{0, 1, 2}, 0
i, data[i] = 2, 100 	// (i = 0) -> (i = 2), (data[0] = 100)

变量定义可以包含初始值，每个变量对应一个。
	var i, j int = 1, 2
如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型。
	var c, python, java = true, false, "no!"

短声明变量
在函数内部，可以使用更简略的 ":=" 方式定义变量。":=" 简洁赋值语句在明确类型的地方，可以用于替代 var 定义。
func main(){
	k := 3
	c, python, java := true, false, "no!"
	//注意检查，是定义新的局部变量，还是修改全局变量。该方式容易造成错误！	
}	
函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， ":=" 结构不能使用在函数外。

特殊只写变量 "_"，用于忽略值占位。

func test() (int, string) {
    return 1, "abc"
}
func main() {
    _, s := test()
	println(s)
}

编译器会将未使 的局部变量当做错误。

var s string // 全局变量没问题。

func main() {
	i := 0 // Error: i declared and not used。(可使  "_ = i" 规避)
}

注意重新赋值与定义新同名变量的区别。

package main

func main(){
	s := "abc"
	println(&s)
	s, y := "hello", 20		// 重新赋值: 与前 s 在同 层次的代码块中，且有新的变量被定义。 
	println(&s, y)	// 通常函数多返回值 err 会被重复使用。

	{
		s, z := 1000, 30	// 定义新同名变量: 不在同 层次代码块。
		println(&s, z)
	}
}


输出:
$ go run main.go 
0xc42003df60
0xc42003df60 20
0xc42003df58 30


Go 语言变量作用域
	作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：
	函数内定义的变量称为局部变量
	函数外定义的变量称为全局变量
	函数定义中的变量称为形式参数
接下来让我们具体了解局部变量、全局变量和形式参数。

局部变量
	在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。

以下实例中 main() 函数使用了局部变量 a, b, c：

package main

import "fmt"

func main() {
   /* 声明局部变量 */
   var a, b, c int

   /* 初始化参数 */
   a = 10
   b = 20
   c = a + b

   fmt.Printf ("结果： a = %d, b = %d and c = %d\n", a, b, c)
}
以上实例执行输出结果为：
结果： a = 10, b = 20 and c = 30

     全局变量
          在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。

全局变量可以在任何函数中使用，以下实例演示了如何使用全局变量：

package main

import "fmt"

/* 声明全局变量 */
var g int

func main() {

   /* 声明局部变量 */
   var a, b int

   /* 初始化参数 */
   a = 10
   b = 20
   g = a + b

   fmt.Printf("结果： a = %d, b = %d and g = %d\n", a, b, g)
}
以上实例执行输出结果为：
结果： a = 10, b = 20 and g = 30

Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。实例如下：
package main

import "fmt"

/* 声明全局变量 */
var g int = 20

func main() {
   /* 声明局部变量 */
   var g int = 10

   fmt.Printf ("结果： g = %d\n",  g)
}
以上实例执行输出结果为：
结果： g = 10


形式参数

形式参数会作为函数的局部变量来使用。实例如下：

package main

import "fmt"

/* 声明全局变量 */
var a int = 20;

func main() {
   /* main 函数中声明局部变量 */
   var a int = 10
   var b int = 20
   var c int = 0

   fmt.Printf("main()函数中 a = %d\n",  a);
   c = sum( a, b);
   fmt.Printf("main()函数中 c = %d\n",  c);
}

/* 函数定义-两数相加 */
func sum(a, b int) int {
   fmt.Printf("sum() 函数中 a = %d\n",  a);
   fmt.Printf("sum() 函数中 b = %d\n",  b);

   return a + b;
}
以上实例执行输出结果为：
main()函数中 a = 10
sum() 函数中 a = 10
sum() 函数中 b = 20
main()函数中 c = 30
初始化局部和全局变量
不同类型的局部和全局变量默认值为：
数据类型    初始化默认值
int    0
float32    0
pointer    nil
