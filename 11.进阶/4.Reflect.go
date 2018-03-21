Reflect
没有运行期类型对象，实例也没有附加字段用来表明身份。只有转换成接口时，才会在其 itab 内部存储与该类型有关的信息，Reflect 所有操作都依赖于此。

Type
以 struct 为例，可获取其全部成员字段信息，包括非导出和匿名字段。
 
type User struct {
    Username string
}

type Admin struct {
    User
    title string
}

func main() {
    var u Admin
    t := reflect.TypeOf(u)
    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type)
    }
}

输出:
User main.User // 可进一步递归。 
title string


如果是指针，应该先使用 Elem 方法获取目标类型，指针本身是没有字段成员的。
 
func main() {
    u := new(Admin)

    t := reflect.TypeOf(u)
    if t.Kind() == reflect.Ptr {
        t = t.Elem() 
    }

    for i, n := 0, t.NumField(); i < n; i++ {
        f := t.Field(i)
        fmt.Println(f.Name, f.Type)
    }
}

同样，value-interface 和 pointer-interface 也会导致方法集存在差异。

type User struct {
}

type Admin struct {
    User
}

func (*User) ToString() {}

func (Admin) test() {}

func main() {
    var u Admin

    methods := func(t reflect.Type) {
        for i, n := 0, t.NumMethod(); i < n; i++ {
            m := t.Method(i)
            fmt.Println(m.Name)
        }
    }

    fmt.Println("--- value interface ---")
    methods(reflect.TypeOf(u))

    fmt.Println("--- pointer interface ---")
    methods(reflect.TypeOf(&u))
}

输出:
--- value interface ---
test
--- pointer interface ---
ToString
test


可直接用名称或序号访问字段，包括用多级序号访问嵌入字段成员。

type User struct {
    Username string
    age int 
}

type Admin struct {
    User
    title string
}

func main() {
    var u Admin
    t := reflect.TypeOf(u)

    f, _ := t.FieldByName("title")
    fmt.Println(f.Name)

    f, _ = t.FieldByName("User")   // 访问嵌入字段。
    fmt.Println(f.Name)

    f, _ = t.FieldByName("Username") // 直接访问嵌入字段成员，会自动深度查找。 
    fmt.Println(f.Name)

    f = t.FieldByIndex([]int{0, 1})    // Admin[0] -> User[1] -> age
    fmt.Println(f.Name)
}

输出:
title
User
Username
age


字段标签可实现简单元数据编程，比如标记 ORM Model 属性。
 
type User struct {
    Name string `field:"username" type:"nvarchar(20)"`
    Age  int    `field:"age" type:"tinyint"`
}

func main() {
    var u User

    t := reflect.TypeOf(u)
    f, _ := t.FieldByName("Name")

    fmt.Println(f.Tag)
    fmt.Println(f.Tag.Get("field"))
    fmt.Println(f.Tag.Get("type"))
}

输出:
field:"username" type:"nvarchar(20)"
username
nvarchar(20)


可从基本类型获取所对应复合类型。
 
var (
    Int    = reflect.TypeOf(0)
    String = reflect.TypeOf("")
)

func main() {
    c := reflect.ChanOf(reflect.SendDir, String)
    fmt.Println(c)

    m := reflect.MapOf(String, Int)
    fmt.Println(m)
    
    s := reflect.SliceOf(Int)
    fmt.Println(s)

    t := struct{ Name string }{}
    p := reflect.PtrTo(reflect.TypeOf(t))
    fmt.Println(p)
}

输出:
chan<- string
map[string]int
[]int
*struct { Name string }


与之对应，方法 Elem 可返回复合类型的基类型。

func main() {
    t := reflect.TypeOf(make(chan int)).Elem()
    fmt.Println(t)
}


方法 Implements 判断是否实现了某个具体接口，AssignableTo、ConvertibleTo 用于赋值和转换判断。
 
type Data struct {
}

func (*Data) String() string {
    return ""
}

func main() {
    var d *Data
    t := reflect.TypeOf(d)

// 没法直接获取接口类型，好在接口本身是个 struct，创建一个空指针对象，这样传递给 TypeOf 转换成 interface{} 时就有类型信息了。。

    it := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

    // 为啥不是 t.Implements(fmt.Stringer)，完全可以由编译器生成。
    fmt.Println(t.Implements(it))
}


某些时候，获取对齐信息对于内存自动分析是很有用的。

type Data struct {
    b   byte
    x int32 
}

func main() {
    var d Data

    t := reflect.TypeOf(d)
    fmt.Println(t.Size(), t.Align())   // sizeof，以及最宽字段的对齐模数。

    f, _ := t.FieldByName("b")
    fmt.Println(f.Type.FieldAlign())   // 字段对齐。
}

输出:
8 4 
1



Value 和 Type 使用方法类似，包括使用 Elem 获取指针目标对象。

type User struct {
    Username string
    age int 
}

type Admin struct {
    User
    title string
}

func main() {
    u := &Admin{User{"Jack", 23}, "NT"}
    v := reflect.ValueOf(u).Elem()
    
    fmt.Println(v.FieldByName("title").String())    // 用转换方法获取字段值
    fmt.Println(v.FieldByName("age").Int())         // 直接访问嵌入字段成员
    fmt.Println(v.FieldByIndex([]int{0, 1}).Int())  // 用多级序号访问嵌入字段成员
}
 
输出:
NT
23 
23


除具体的 Int、String 等转换方法，还可返回 interface{}。只是非导出字段方法使用，需用 CanInterface 判断一下。

type User struct {
    Username string
    age int 
}

func main() {
    u := User{"Jack", 23}
    v := reflect.ValueOf(u)

    fmt.Println(v.FieldByName("Username").Interface())
    fmt.Println(v.FieldByName("age").Interface())
}

输出:
Jack
panic: reflect.Value.Interface: cannot return value obtained from unexported field or method



当然，转换成具体类型不会引发 panic。
 
func main() {
    u := User{"Jack", 23}
    v := reflect.ValueOf(u)

    f := v.FieldByName("age")

    if f.CanInterface() {
        fmt.Println(f.Interface())
    } else {
        fmt.Println(f.Int())
    } 
}


除 struct，其他复合类型 array、slice、map 取值示例。

func main() {
    v := reflect.ValueOf([]int{1, 2, 3})
    for i, n := 0, v.Len(); i < n; i++ {
        fmt.Println(v.Index(i).Int())

    }

    fmt.Println("---------------------------")

    v = reflect.ValueOf(map[string]int{"a": 1, "b": 2})
    for _, k := range v.MapKeys() {
        fmt.Println(k.String(), v.MapIndex(k).Int())
    }
}

输出:
1
2
3 
--------------------------- 
a 1
b 2



需要注意，Value 某些方法没有遵循 "comma ok" 模式，而是返回 ZeroValue，因此需要用 IsValid 判断一下是否可用。
 
func (v Value) FieldByName(name string) Value {
    v.mustBe(Struct)
    if f, ok := v.typ.FieldByName(name); ok {
        return v.FieldByIndex(f.Index)
    }
    return Value{}
}

type User struct {
    Username string
    age int 
}

func main() {
    u := User{}
    v := reflect.ValueOf(u)

    f := v.FieldByName("a")
    fmt.Println(f.Kind(), f.IsValid())
}

输出:
invalid false


另外，接口是否为 nil，需要 tab 和 data 都为空。可使用 IsNil 方法判断 data 值。

func main() {
    var p *int
    var x interface{} = p
    fmt.Println(x == nil)
    v := reflect.ValueOf(p)
    fmt.Println(v.Kind(), v.IsNil())
}

输出:
false
ptr true


将对象转换为接口，会发生复制行为。该复制品只读，无法被修改。所以要通过接口改变目标对象状态，必须是 pointer-interface。

就算是指针，我们依然没法将这个存储在 data 的指针指向其他对象，只能透过它修改目标对象。因为目标对象并没有被复制，被复制的只是指针。

  
type User struct {
    Username string
    age int 
}

func main() {
    u := User{"Jack", 23}

    v := reflect.ValueOf(u)
    p := reflect.ValueOf(&u)

    fmt.Println(v.CanSet(), v.FieldByName("Username").CanSet())
    fmt.Println(p.CanSet(), p.Elem().FieldByName("Username").CanSet())
}

输出:
false  false
false  true


非导出字段无法直接修改，可改用指针操作。
 
type User struct {
    Username string
    age int 
}

func main() {
    u := User{"Jack", 23}
    p := reflect.ValueOf(&u).Elem()

    p.FieldByName("Username").SetString("Tom")

    f := p.FieldByName("age")
    fmt.Println(f.CanSet())

    // 判断是否能获取地址。 
    if f.CanAddr() {
        age := (*int)(unsafe.Pointer(f.UnsafeAddr()))
        // age := (*int)(unsafe.Pointer(f.Addr().Pointer())) 
        // 等同 *age = 88
    }

    // 注意 p 是 Value 类型，需要还原成接 才能转型。
    fmt.Println(u, p.Interface().(User))
}

输出:
false
{Tom 88} {Tom 88}


复合类型修改示例。
 
func main() {
    s := make([]int, 0, 10)
    v := reflect.ValueOf(&s).Elem()

    v.SetLen(2)
    v.Index(0).SetInt(100)
    v.Index(1).SetInt(200)

    fmt.Println(v.Interface(), s)

    v2 := reflect.Append(v, reflect.ValueOf(300))
    v2 = reflect.AppendSlice(v2, reflect.ValueOf([]int{400, 500}))

    fmt.Println(v2.Interface())

    fmt.Println("----------------------")

    m := map[string]int{"a": 1}
    v = reflect.ValueOf(&m).Elem()

    v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(100)) // update
    v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(200)) // add

    fmt.Println(v.Interface(), m)
}


输出:
[100 200] [100 200]
[100 200 300 400 500]
----------------------
map[a:100 b:200] map[a:100 b:200]



Method 可获取方法参数、返回值类型等信息。
 
type Data struct {
}

func (*Data) Test(x, y int) (int, int) {
    return x + 100, y + 100
}

func (*Data) Sum(s string, x ...int) string {
    c := 0
    for _, n := range x {
        c += n
    }

    return fmt.Sprintf(s, c)
}

func info(m reflect.Method) {
    t := m.Type

    fmt.Println(m.Name)

    for i, n := 0, t.NumIn(); i < n; i++ {
        fmt.Printf("  in[%d] %v\n", i, t.In(i))
    }

    for i, n := 0, t.NumOut(); i < n; i++ {
        fmt.Printf("  out[%d] %v\n", i, t.Out(i))
    }
}

func main() {
    d := new(Data)
    t := reflect.TypeOf(d)

    test, _ := t.MethodByName("Test")
    info(test)
    
    sum, _ := t.MethodByName("Sum")
    info(sum) 
}


输出:
Test
  in[0] *main.Data  // receiver
  in[1] int
  in[2] int
  out[0] int
  out[1] int
Sum
  in[0] *main.Data
  in[1] string
  in[2] []int
  out[0] string


动态调用方法很简单，按 In 列表准备好所需参数即可 (不包括 receiver)。

func main() {
    d := new(Data)
    v := reflect.ValueOf(d)

    exec := func(name string, in []reflect.Value) {
        m := v.MethodByName(name)
        out := m.Call(in)
        for _, v := range out {
            fmt.Println(v.Interface())
        } 
    }

    exec("Test", []reflect.Value{
        reflect.ValueOf(1),
        reflect.ValueOf(2),
    })

    fmt.Println("-----------------------")

    exec("Sum", []reflect.Value{
        reflect.ValueOf("result = %d"),
        reflect.ValueOf(1),
        reflect.ValueOf(2),
    })
}


输出:
101
102
-----------------------
result = 3


如改用 CallSlice，只需将变参打包成 slice 即可。
  
func main() {
    d := new(Data)
    v := reflect.ValueOf(d)

    m := v.MethodByName("Sum")

    in := []reflect.Value{
        reflect.ValueOf("result = %d"),
        reflect.ValueOf([]int{1, 2}),   // 将变参打包成 slice。
    }

    out := m.CallSlice(in)

    for _, v := range out {
        fmt.Println(v.Interface())
    } 
}

非导出方法无法调用，甚至无法透过指针操作，因为接口类型信息中没有该方法地址。


利用 Make、New 等函数，可实现近似泛型操作。

var (
    Int    = reflect.TypeOf(0)
    String = reflect.TypeOf("")
)

func Make(T reflect.Type, fptr interface{}) {
    
    // 实际创建 slice 的包装函数。
    swap := func(in []reflect.Value) []reflect.Value {

        // --- 省略算法内容 --- //

    // 返回和类型匹配的 slice 对象。 
    return []reflect.Value{
        reflect.MakeSlice(
            reflect.SliceOf(T), // slice type
            int(in[0].Int()), // len
            int(in[1].Int()) // cap
        ), 
    }
}

    // 传入的是函数变量指针，因为我们要将变量指向 swap 函数。 
    fn := reflect.ValueOf(fptr).Elem()

    // 获取函数指针类型，生成所需 swap function value。 
    v := reflect.MakeFunc(fn.Type(), swap)

    // 修改函数指针实际指向，也就是 swap。
    fn.Set(v) 
}

func main() {
    var makeints func(int, int) []int
    var makestrings func(int, int) []string

    // 用相同算法，生成不同类型创建函数。 
    Make(Int, &makeints) 
    Make(String, &makestrings)

    // 按实际类型使用。
    x := makeints(5, 10) 
    fmt.Printf("%#v\n", x)

    s := makestrings(3, 10)
    fmt.Printf("%#v\n", s)
}

输出:
[]int{0, 0, 0, 0, 0}
[]string{"", "", ""}


原理并不复杂。
1. 核心是提供一个 swap 函数，其中利用 reflect.MakeSlice 生成最终 slice 对象， 因此需要传入 element type、len、cap 参数。

2. 接下来，利用 MakeFunc 函数生成 swap value，并修改函数变量指向，以达到调用 swap 的目的。

3. 当调用具体类型的函数变量时，实际内部调用的是 swap，相关代码会自动转换参 数列表，并将返回结果还原成具体类型返回值。

如此，在共享算法的前提下，无须用 interface{}，无须做类型转换，颇有泛型的效果。



