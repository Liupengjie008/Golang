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
