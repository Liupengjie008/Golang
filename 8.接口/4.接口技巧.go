Golang 接口技巧 ：

让编译器检查，以确保某个类型实现接口。 

var _ fmt.Stringer = (*Data)(nil)

某些时候，让函数直接 "实现" 接口能省不少事。
 
type Tester interface {
    Do()
}

type FuncDo func()
func (self FuncDo) Do() { self() }

func main() {
    var t Tester = FuncDo(func() { println("Hello, World!") })
    t.Do()
}