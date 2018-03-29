循环控制语句

循环控制语句可以控制循环体内语句的执行过程。
GO 语言支持以下几种循环控制语句：

Goto、Break、Continue
    1.三个语句都可以配合标签(label)使用
    2.标签名区分大小写，定以后若不使用会造成编译错误
    3.continue、break配合标签(label)可用于多层循环跳出
    4.goto是调整执行位置，与continue、break配合标签(label)的结果并不相同


goto 语句    将控制转移到被标记的语句。

Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
goto语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
但是，在结构化程序设计中一般不主张使用goto语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
语法
goto 语法格式如下：
goto label;
..
.
label: statement;

Golang支持在函数内 goto 跳转。标签名区分大小写，未使用标签引发错误。

func main() {
    var i int
    for {
        println(i)
        i++
        if i > 2 { goto BREAK }
    }
BREAK:
    println("break")
EXIT:                 // Error: label EXIT defined and not used
}

goto 实例
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* 循环 */
   LOOP: for a < 20 {
      if a == 15 {
         /* 跳过迭代 */
         a = a + 1
         goto LOOP
      }
      fmt.Printf("a的值为 : %d\n", a)
      a++     
   }  
}
以上实例执行结果为：
a的值为 : 10
a的值为 : 11
a的值为 : 12
a的值为 : 13
a的值为 : 14
a的值为 : 16
a的值为 : 17
a的值为 : 18
a的值为 : 19

控制语句    

break 语句    经常用于中断当前 for 循环或跳出 switch 语句

Go 语言中 break 语句用于以下两方面：
    1.用于循环语句中跳出循环，并开始执行循环之后的语句。
    2.break在switch（开关语句）中在执行一条case后跳出语句的作用。


实例：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      fmt.Printf("a 的值为 : %d\n", a)
      a++
      if a > 15 {
         /* 使用 break 语句跳出循环 */
         break
      }
   }
}
以上实例执行结果为：
a 的值为 : 10
a 的值为 : 11
a 的值为 : 12
a 的值为 : 13
a 的值为 : 14
a 的值为 : 15


Break label 语句：我们在for多层嵌套时，有时候需要直接跳出所有嵌套循环， 这时候就可以用到go的label breaks特征了。

先看一个范例代码：

package main

import (
    "fmt"
)

func main() {
    fmt.Println("1")

Exit:
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            if i+j > 15 {
                fmt.Print("exit")
                break Exit
            }
        }
    }

    fmt.Println("3")
}

执行效果：
1
exit3

注意：label要写在for循环的开始而不是结束的地方。和goto是不一样的。虽然它是直接break退出到指定的位置。

break的标签和goto的标签的区别可以参考下面代码：

JLoop:
    for i := 0; i < 10; i++ {
        fmt.Println("label i is ", i)
        for j := 0; j < 10; j++ {
            if j > 5 {
                //跳到外面去啦，但是不会再进来这个for循环了
                break JLoop
            }
        }
    }

    //跳转语句 goto语句可以跳转到本函数内的某个标签
    gotoCount := 0
GotoLabel:
    gotoCount++
    if gotoCount < 10 {
        goto GotoLabel //如果小于10的话就跳转到GotoLabel
    }


break 标签除了可以跳出 for 循环，还可以跳出 select switch 循环， 参考下面代码：

L:
    for ; count < 8192; count++ {
        select {
        case e := <-self.pIdCh:
            args[count] = e

        default:
            break L // 跳出 select 和 for 循环
        }

    }



continue 语句    跳过当前循环的剩余语句，然后继续进行下一轮循环。

Go 语言的 continue 语句 有点像 break 语句。但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
for 循环中，执行 continue 语句会触发for增量语句的执行。

实例
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 10

   /* for 循环 */
   for a < 20 {
      if a == 15 {
         /* 跳过此次循环 */
         a = a + 1
         continue
      }
      fmt.Printf("a 的值为 : %d\n", a)
      a++    
   }  
}
以上实例执行结果为：
a 的值为 : 10
a 的值为 : 11
a 的值为 : 12
a 的值为 : 13
a 的值为 : 14
a 的值为 : 16
a 的值为 : 17
a 的值为 : 18
a 的值为 : 19


配合标签，break 和 continue 可在多级嵌套循环中跳出。

func main() {
L1:
    for x := 0; x < 3; x++ {
L2:
        for y := 0; y < 5; y++ {
            if y > 2 { continue L2 }
            if x > 1 { break L1 }
            
            print(x, ":", y, " ")
        }
        println() 
    }
}

输出:
0:0  0:1  0:2
1:0  1:1  1:2


附:break 可用于 for、switch、select，而 continue 仅能用于 for 循环。

x := 100

switch {
case x >= 0:
    if x == 0 { break }
    println(x) 
}


goto、continue、break语句：
   
package main

import "fmt"

func main() {

    //goto直接调到LAbEL2
    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                goto LAbEL2
            }
        }
    }
    fmt.Println("PreLAbEL2")
LAbEL2:
    fmt.Println("LastLAbEL2")

    //break跳出和LAbEL1同一级别的循环,继续执行其他的
LAbEL1:
    for {
        for i := 0; i < 10; i++ {
            if i > 3 {
                break LAbEL1
            }
        }
    }
    fmt.Println("OK")

    //continue
LABEL3:

    for i := 0; i < 3; i++ {
        for {
            continue LABEL3
        }
    }
    fmt.Println("ok")
}
输出如下：

LastLAbEL2
OK
ok













