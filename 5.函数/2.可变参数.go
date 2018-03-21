Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。

func test(s string, n ...int) string {
    var x int
    for _, i := range n {
        x += i
	}

    return fmt.Sprintf(s, x)
}

func main() {
    println(test("sum: %d", 1, 2, 3))
}

使用 slice 对象做变参时，必须展开。

func main() {
    s := []int{1, 2, 3}
    println(test("sum: %d", s...))
}