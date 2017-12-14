Golang返回值不能用容器对象接收多返回值。只能用多个变量，或 "_" 忽略。
 
func test() (int, int) {
    return 1, 2
}

func main() {
    // s := make([]int, 2)
    // s = test()   // Error: multiple-value test() in single-value context
    
    x, _ := test()
    println(x) 
}


多返回值可直接作为其他函数调用实参。

func test() (int, int) {
    return 1, 2
}

func add(x, y int) int {
    return x + y
}

func sum(n ...int) int {
    var x int
    for _, i := range n {
        x += i
    }

    return x 
}

func main() {
    println(add(test()))
    println(sum(test()))
}


命名返回参数可看做与形参类似的局部变量，最后由 return 隐式返回。
 
func add(x, y int) (z int) { 
    z = x + y
    return 
}

func main() {
    println(add(1, 2))
}


命名返回参数可被同名局部变量遮蔽，此时需要显式返回。

func add(x, y int) (z int) {
    {   // 不能在一个级别，引发 "z redeclared in this block" 错误。
        var z = x + y
        // return   // Error: z is shadowed during return
        return z    // 必须显式返回。
    } 
}


命名返回参数允许 defer 延迟调用通过闭包读取和修改。

func add(x, y int) (z int) {
    defer func() {
        z += 100 
    }()
        
    z=x+y
    return 
}

func main() {
    println(add(1, 2))  // 输出: 103
}


显式 return 返回前，会先修改命名返回参数。
 
func add(x, y int) (z int) {
    defer func() {
        println(z)  // 输出: 203
    }()

    z = x + y 
    return z + 200  // 执行顺序: (z = z + 200) -> (call defer) -> (ret)
}

func main() {
    println(add(1, 2))  // 输出: 203
}

