Golang 接口转换 ：利用类型推断，可判断接口对象是否某个具体的接口或类型。
 
type User struct {
    id   int
    name string
}

func (self *User) String() string {
    return fmt.Sprintf("%d, %s", self.id, self.name)
}

func main() {
    var o interface{} = &User{1, "Tom"}

    if i, ok := o.(fmt.Stringer); ok {  // ok-idiom
        fmt.Println(i)    
    }

    u := o.(*User)
    // u := o.(User)    // panic: interface is *main.User, not main.User
    fmt.Println(u)
}


还可用 switch 做批量类型判断，不支持 fallthrough。

func main() {
    var o interface{} = &User{1, "Tom"}

    switch v := o.(type) {
    case nil:               // o == nil
        fmt.Println("nil")
    case fmt.Stringer:         // interface
        fmt.Println(v)
    case func() string:         // func
        fmt.Println(v())
    case *User:             // *struct
        fmt.Printf("%d, %s\n", v.id, v.name)
    default:
        fmt.Println("unknown")
    } 
}


超集接口对象可转换为子集接口，反之出错。

type Stringer interface {
    String() string
}

type Printer interface {
    String() string
    Print() 
}

type User struct {
    id   int
    name string 
}

func (self *User) String() string {
    return fmt.Sprintf("%d, %v", self.id, self.name)
}

func (self *User) Print() {
    fmt.Println(self.String())
}

func main() {
    var o Printer = &User{1, "Tom"}
    var s Stringer = o
    fmt.Println(s.String())
}






