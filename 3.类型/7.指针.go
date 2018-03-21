Golang 支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。

• 默认值 nil，没有 NULL 常量。
• 操作符 "&" 取变量地址，"*" 透过指针访问目标对象。
• 不支持指针运算，不支持 "->" 运算符，直接用 "." 访问目标成员。

func main() {
    type data struct{ a int }
    
    var d = data{1234}
    var p *data

    p = &d
    fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。 
}

输出:
0xc042008240, 1234


不能对指针做加减法等运算。

x := 1234
p := &x
p++      // Error: invalid operation: p += 1 (mismatched types *int and int)

可以在 unsafe.Pointer 和任意类型指针间进 转换。

func main() {
    x := 0x12345678

    p := unsafe.Pointer(&x)     // *int -> Pointer
    n := (*[4]byte)(p)          // Pointer -> *[4]byte
    
    for i := 0; i < len(n); i++ {
        fmt.Printf("%X ", n[i])
    } 
}

输出:
78 56 34 12


返回局部变量指针是安全的，编译器会根据需要将其分配在 GC Heap 上。

func test() *int {
    x := 100
    return &x   // 在堆上分配 x 内存。但在内联时，也可能直接分配在目标栈。 
}

将 Pointer 转换成 uintptr，可变相实现指针运算。
  
func main() {
    d := struct {
        s   string
        x   int
    }{"abc", 100}

    p := uintptr(unsafe.Pointer(&d))    // *struct -> Pointer -> uintptr
    p += unsafe.Offsetof(d.x)       // uintptr + offset

    p2 := unsafe.Pointer(p)         // uintptr -> Pointer
    px := (*int)(p2)            // Pointer -> *int
    *px = 200               // d.x = 200
    
    fmt.Printf("%#v\n", d)

}

输出:
struct { s string; x int }{s:"abc", x:200}

注意:GC 把 uintptr 当成普通整数对象，它无法阻止 "关联" 对象被回收。