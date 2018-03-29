Goto、Break、Continue

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
