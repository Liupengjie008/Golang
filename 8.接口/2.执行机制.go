Golang 接口执行机制 ：接口对象由接口表 (interface table) 指针和数据指针组成。

runtime.h
struct Iface
{
    Itab* tab;
    void*    data;
};

struct Itab {
    InterfaceType*    inter;
    Type*             type;
    void (*fun[])(void);
};


接口表存储元数据信息，包括接口类型、动态类型，以及实现接口的方法指针。无论是反射还是通过接口调用方法，都会用到这些信息。

数据指针持有的是目标对象的只读复制品，复制完整对象或指针。


type User struct {
    id   int
    name string 
}

func main() {
    u := User{1, "Tom"}
    var i interface{} = u

    u.id = 2
    u.name = "Jack"

    fmt.Printf("%v\n", u)
    fmt.Printf("%v\n", i.(User))
}

输出:
{2 Jack}
{1 Tom}


接口转型返回临时对象，只有使用指针才能修改其状态。

type User struct {
    id   int
    name string 
}

func main() {
    u := User{1, "Tom"}
    var vi, pi interface{} = u, &u

    // vi.(User).name = "Jack"       // Error: cannot assign to vi.(User).name
    pi.(*User).name = "Jack"

    fmt.Printf("%v\n", vi.(User))
    fmt.Printf("%v\n", pi.(*User))
}

输出:
{1 Tom}
&{1 Jack}

只有 tab 和 data 都为 nil 时，接口才等于 nil。


var a interface{} = nil     // tab = nil, data = nil
var b interface{} = (*int)(nil)    // tab 包含 *int 类型信息, data = nil

type iface struct {
    itab, data uintptr
}

ia := *(*iface)(unsafe.Pointer(&a))
ib := *(*iface)(unsafe.Pointer(&b))

fmt.Println(a == nil, ia)
fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())

输出:
true  {0 0}
false {505728 0} true




