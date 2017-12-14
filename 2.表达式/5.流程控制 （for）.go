Golang for支持三种循环方式，包括类似 while 的语法。

s := "abc"

for i, n := 0, len(s); i < n; i++ { // 常见的 for 循环，支持初始化语句。
    println(s[i])
}

n := len(s)
for n > 0 {				// 替代 while (n > 0) {}
    println(s[n])		// 替代 for (; n > 0;) {}
n-- }

for {					// 替代 while (true) {}
    println(s)			// 替代 for (;;) {}
}


不要期望编译器能理解你的想法，在初始化语句中计算出全部结果是个好主意。

func length(s string) int {
    println("call length.")
    return len(s)
}

func main() {
    s := "abcd"
    
    for i, n := 0, length(s); i < n; i++ { 	// 避免多次调用 length 函数。
        println(i, s[i])
	} 
}

输出:
call length.
0 97
1 98
2 99
3 100