Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。

• 只能为当前包内命名类型定义方法。
• 参数 receiver 可任意命名。如方法中未曾使 ，可省略参数名。
• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。 
• 不支持方法重载，receiver 只是参数签名的组成部分。
• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。


没有构造和析构方法，通常用简单工厂模式返回对象实例。

type Queue struct {
    elements []interface{}
}

func NewQueue() *Queue {    // 创建对象实例。
    return &Queue{make([]interface{}, 10)}
}

func (*Queue) Push(e interface{}) error {  // 省略 receiver 参数名。
    panic("not implemented")
}

// func (Queue) Push(e int) error {   // Error: method redeclared: Queue.Push
//     panic("not implemented")
// }

func (self *Queue) length() int {    // receiver 参数名可以是 self、this 或其他。
    return len(self.elements)
}


方法不过是一种特殊的函数，只需将其还原，就知道 receiver T 和 *T 的差别。

type Data struct{
    x int
}

func (self Data) ValueTest() {     // func ValueTest(self Data);
    fmt.Printf("Value: %p\n", &self)
}

func (self *Data) PointerTest() {  // func PointerTest(self *Data);
    fmt.Printf("Pointer: %p\n", self)
}

func main() {
    d := Data{}
    p := &d
    fmt.Printf("Data: %p\n", p)

    d.ValueTest()     // ValueTest(d)
    d.PointerTest()     // PointerTest(&d)

    p.ValueTest()       // ValueTest(*p)
    p.PointerTest()     // PointerTest(p)
}

输出:
Data: 0xc0420361d0
Value: 0xc0420361f0
Pointer: 0xc0420361d0
Value: 0xc0420361f8
Pointer: 0xc0420361d0


从 1.4 开始，不再支持多级指针查找方法成员。
 
type X struct{}

func (*X) test() {
    println("X.test")
}

func main() {
    p := &X{}
    p.test()
    
    // Error: calling method with receiver &p (type **X) requires explicit dereference (&p).test()
}