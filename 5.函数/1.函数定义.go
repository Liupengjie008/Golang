Golang函数定义不支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。

• 无需声明原型。
• 支持不定 变参。
• 支持多返回值。
• 支持命名返回参数。 
• 支持匿名函数和闭包。

使用关键字 func 定义函数，左大括号依旧不能另起一行。

func test(x, y int, s string) (int, string) { // 类型相同的相邻参数，参数类型可合并。 多返回值必须用括号。
	n := x + y 		 
	return n, fmt.Sprintf(s, n)
}


函数是第一类对象，可作为参数传递。建议将复杂签名定义为函数类型，以便于阅读。
 
func test(fn func() int) int {
    return fn()
}

type FormatFunc func(s string, x, y int) string // 定义函数类型。

func format(fn FormatFunc, s string, x, y int) string {
    return fn(s, x, y)
}

func main() {
    s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。
    
    s2 := format(func(s string, x, y int) string {
        return fmt.Sprintf(s, x, y)
    }, "%d, %d", 10, 20)
    
    println(s1, s2)
}


有返回值的函数，必须有明确的终止语句，否则会引发编译错误。