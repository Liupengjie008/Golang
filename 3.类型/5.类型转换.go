Golang 不支持隐式类型转换，即便是从窄向宽转换也不行。

var b byte = 100
// var n int = b // Error: cannot use b (type byte) as type int in assignment 
var n int = int(b) // 显式转换

使用括号避免优先级错误。
 
*Point(p) // 相当于 *(Point(p)) 
(*Point)(p)
<-chan int(c) // 相当于 <-(chan int(c)) 
(<-chan int)(c)

同样不能将其他类型当 bool 值使用。

a := 100
if a {   	// Error: non-bool a (type int) used as if condition
    println("true")
}


类型转换用于将一种数据类型的变量转换为另外一种类型的变量。Go 语言类型转换基本格式如下：
type_name(expression)
type_name 为类型，expression 为表达式。

实例
以下实例中将整型转化为浮点型，并计算结果，将结果赋值给浮点型变量：
package main

import "fmt"

func main() {
   var sum int = 17
   var count int = 5
   var mean float32
   
   mean = float32(sum)/float32(count)
   fmt.Printf("mean 的值为: %f\n",mean)
}
以上实例执行输出结果为：
mean 的值为: 3.400000

Go各种类型转换及函数的高级用法

整形转字符串
fmt.Println(strconv.Itoa(100))
该方法的源码是：
// Itoa is shorthand for FormatInt(i, 10).
func Itoa(i int) string {
    return FormatInt(int64(i), 10)
}
可以看出是FormatInt方法的简单实现。

字符串转整形
i, _ := strconv.Atoi("100")
fmt.Println(i)

64位整形转字符串
var i int64
i = 0x100
fmt.Println(strconv.FormatInt(i, 10))
FormatInt第二个参数表示进制，10表示十进制。

字节转32位整形
b := []byte{0x00, 0x00, 0x03, 0xe8}
bytesBuffer := bytes.NewBuffer(b)

var x int32
binary.Read(bytesBuffer, binary.BigEndian, &x)
fmt.Println(x)
其中binary.BigEndian表示字节序，相应的还有little endian。通俗的说法叫大端、小端。

32位整形转字节
var x int32
x = 106
bytesBuffer := bytes.NewBuffer([]byte{})
binary.Write(bytesBuffer, binary.BigEndian, x)
fmt.Println(bytesBuffer.Bytes())

字节转字符串
fmt.Println(string([]byte{97, 98, 99, 100}))

字符串转字节
fmt.Println([]byte("abcd"))

零值
变量在定义时没有明确的初始化时会赋值为 零值 。

零值是：

数值类型为 0 ，
布尔类型为 false ，
字符串为 "" （空字符串）。

类型转换
表达式 T(v) 将值 v 转换为类型 T 。

一些关于数值的转换：

var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
或者，更加简单的形式：

i := 42
f := float64(i)
u := uint(f)
与 C 不同的是 Go 的在不同类型之间的项目赋值时需要显式转换。 试着移除例子中 float64 或 int 的转换看看会发生什么。

类型推导     
在定义一个变量却并不显式指定其类型时（使用 := 语法或者 var = 表达式语法）【全局变量不适用】， 变量的类型由（等号）右侧的值推导得出。

当右值定义了类型时，新变量的类型与其相同：

var i int
j := i // j 也是一个 int
但是当右边包含了未指名类型的数字常量时，新的变量就可能是 int 、 float64 或 complex128 。 这取决于常量的精度：

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128
尝试修改演示代码中 v 的初始值，并观察这是如何影响其类型的。

