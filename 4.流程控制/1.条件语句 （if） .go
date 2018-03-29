Go 语言条件语句：
     条件语句需要开发者通过指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。

Go 语言提供了以下几种条件判断语句：

if 语句    if 语句 由一个布尔表达式后紧跟一个或多个语句组成。

Go 编程语言中 if 语句的语法如下：
	• 可省略条件表达式括号。
	• 持初始化语句，可定义代码块局部变量。 
	• 代码块左 括号必须在条件表达式尾部。

	if 布尔表达式 {
	/* 在布尔表达式为 true 时执行 */
	}
If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则不执行。

 x := 0

// if x > 10		// Error: missing condition in if statement
// {
// }

if n := "abc"; x > 0 { 	// 初始化语句未必就是定义变量， 如 println("init") 也是可以的。
    println(n[2])
} else if x < 0 {	// 注意 else if 和 else 左大括号位置。
    println(n[1])
} else {
    println(n[0])
}

 
*不支持三元操作符(三目运算符) "a > b ? a : b"。



实例
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10
   /* 使用 if 语句判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)
}
以上代码执行结果为：
a 小于 20
a 的值为 : 10

if...else 语句    if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。

语法

Go 编程语言中 if...else 语句的语法如下：
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
If 在布尔表达式为 true 时，其后紧跟的语句块执行，如果为 false 则执行 else 语句块。
实例
package main

import "fmt"

func main() {
   /* 局部变量定义 */
   var a int = 100
   /* 判断布尔表达式 */
   if a < 20 {
       /* 如果条件为 true 则执行以下语句 */
       fmt.Printf("a 小于 20\n" )
   } else {
       /* 如果条件为 false 则执行以下语句 */
       fmt.Printf("a 不小于 20\n" )
   }
   fmt.Printf("a 的值为 : %d\n", a)

}
以上代码执行结果为：
a 不小于 20
a 的值为 : 100

if 嵌套语句    你可以在 if 或 else if 语句中嵌入一个或多个 if 或 else if 语句。

语法

Go 编程语言中 if...else 语句的语法如下：
if 布尔表达式 1 {
   /* 在布尔表达式 1 为 true 时执行 */
   if 布尔表达式 2 {
      /* 在布尔表达式 2 为 true 时执行 */
   }
}
你可以以同样的方式在 if 语句中嵌套 else if...else 语句
实例
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   /* 判断条件 */
   if a == 100 {
       /* if 条件语句为 true 执行 */
       if b == 200 {
          /* if 条件语句为 true 执行 */
          fmt.Printf("a 的值为 100 ， b 的值为 200\n" )
       }
   }
   fmt.Printf("a 值为 : %d\n", a )
   fmt.Printf("b 值为 : %d\n", b )
}
以上代码执行结果为：
a 的值为 100 ， b 的值为 200
a 值为 : 100
b 值为 : 200









