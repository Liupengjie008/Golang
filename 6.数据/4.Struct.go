Golang Struct 值类型，赋值和传参会复制全部内容。可用 "_" 定义补位字段，支持指向自身类型的指针成员。


type Node struct {
    _    int
    id   int
    data *byte
    next *Node
}

func main() {
    n1 := Node{
        id: 1,
        data: nil, 
    }

    n2 := Node{
        id:   2,
        data: nil,
        next: &n1, 
    }
}


顺序初始化必须包含全部字段，否则会出错。

type User struct {
    name string
    age int 
}

u1 := User{"Tom", 20}
u2 := User{"Tom"}   // Error: too few values in struct initializer
 

支持匿名结构，可用作结构成员或定义变量。

type File struct {
    name string
    size int
    attr struct {
        perm int
        owner int 
    }
}

f := File{
    name: "test.txt",
    size: 1025,
    // attr: {0755, 1},     // Error: missing type in composite literal
}

f.attr.owner = 1
f.attr.perm = 0755

var attr = struct {
    perm  int
    owner int
}{2, 0755}

f.attr = attr


支持 "=="、"!=" 相等操作符，可用作 map 键类型。

type User struct {
    id   int
    name string 
}

m := map[User]int{
    User{1, "Tom"}: 100,
}


*可定义字段标签，用反射读取。标签是类型的组成部分。

var u1 struct { name string "username" }
var u2 struct { name string }

u2 = u1   // Error: cannot use u1 (type struct { name string "username" }) as type struct { name string } in assignment


空结构 "节省" 内存， 如用来实现 set 数据结构，或者实现没有 "状态" 只有方法的 "静态类"。

var null struct{}

set := make(map[string]struct{})
set["a"] = null


**匿名字段 ：匿名字段不过是一种语法糖，从根本上说，就是一个与成员类型同名 (不含包名) 的字段。被匿名嵌入的可以是任何类型，当然也包括指针。
 
type User struct {
    name string
}

type Manager struct {
    User
    title string
}

m := Manager{
    User: User{"Tom"}, // 匿名字段的显式字段名，和类型名相同。 
    title: "Administrator",
}


可以像普通字段那样访问匿名字段成员，编译器从外向内逐级查找所有层次的匿名字段，直到发现目标或出错。

type Resource struct {
    id int
}

type User struct {
    Resource
    name string 
}

type Manager struct {
    User
    title string
}

var m Manager
m.id = 1
m.name = "Jack"
m.title = "Administrator"


外层同名字段会遮蔽嵌入字段成员，相同层次的同名字段也会让编译器无所适从。解决方法是使用显式字段名。

type Resource struct {
    id   int
    name string 
}

type Classify struct {
    id int
}

type User struct {
    Resource    // Resource.id 与 Classify.id 处于同一层次。
    Classify
    name string     // 遮蔽 Resource.name。
}

u := User{
    Resource{1, "people"},  
    Classify{100},  
    "Jack",
}

println(u.name)         // User.name: Jack
println(u.Resource.name)     // people

// println(u.id)        // Error: ambiguous selector u.id
println(u.Classify.id)  // 100



不能同时嵌入某一类型和其指针类型，因为它们名字相同。

type Resource struct {
    id int
}

type User struct {
    *Resource
    // Resource     // Error: duplicate field Resource
    name string 
}

u := User{
    &Resource{1},
    "Administrator",
}

println(u.id)
println(u.Resource.id)


**面向对象 ：面向对象三大特征 ，Go 仅支持封装，尽管匿名字段的内存布局和行为类似继承。没有 class 关键字，没有继承、多态等等。

type User struct {
    id   int
    name string 
}

type Manager struct {
    User
    title string
}

m := Manager{User{1, "Tom"}, "Administrator"}

// var u User = m // Error: cannot use m (type Manager) as type User in assignment （没有继承， 然也不会有多态）

var u User = m.User // 同类型拷 。


内存布局和 C struct 相同，没有任何附加的 object 信息。

   |<----- User:24 ---->|<-title:16->|
   +--------+-----------+------------+  +---------------+
m  |   1    |  string   |  string    |  | Administrator |  [n]byte
   +--------+-----------+------------+  +---------------+
               |              |                   |
               |              +--->>>------->>>---+ 
               +--->>>------------>>>-----+
                 | +--->>>------>>>-+ |   |
                 |                    |   |
   +--------+-----------+          +---------+
u  | 1      |   string  |          |   Tom   |   [n]byte
   +--------+-----------+          +---------+
   |<-id:8->|<-name:16->|



可用 unsafe 包相关函数输出内存地址信息。
m      : 0x2102271b0, size: 40, align: 8
m.id   : 0x2102271b0, offset: 0
m.name : 0x2102271b8, offset: 8
m.title: 0x2102271c8, offset: 24
