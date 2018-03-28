Go中的 struct：

可将类型分为命名和未命名两大类。命名类型包括 bool、int、string 等，而 array、slice、map 等和具体元素类型、长度等有关，属于未命名类型。 

具有相同声明的未命名类型被视为同一类型。

• 具有相同基类型的指针。
• 具有相同元素类型和长度的 array。
• 具有相同元素类型的 slice。
• 具有相同键值类型的 map。
• 具有相同元素类型和传送方向的 channel。
• 具有相同字段序列 (字段名、类型、标签、顺序) 的匿名 struct。 
• 签名相同 (参数和返回值，不包括参数名称) 的 function。
• 方法集相同 ( 方法名、方法签名相同，和次序无关) 的 interface。



struct 特点：
1. 用来自定义复杂数据结构
2. struct里面可以包含多个字段（属性）
3. struct类型可以定义方法，注意和函数的区分
4. struct类型是值类型
5. struct类型可以嵌套
6. Go语言没有class类型，只有struct类型
7. 结构体是用户单独定义的类型，不能和其他类型进行强制转换
8.golang中的struct没有构造函数，一般可以使用工厂模式来解决这个问题
9. 我们可以为struct中的每个字段，写上一个tag。这个tag可以通过反射的机制获取到，最常用的场景就是json序列化和反序列化。

概述
    与C语言struct一样，与java/php等class类似，在Go中，用于扩展类型，面向对象编程(本部分暂未做详细介绍)等

可用 type 在全局或函数内定义新类型。
声明格式：(是类型的组合)

[plain] view plain copy print?
type typeName struct {  
  //...  
}  


var a struct { x int `a` }
var b struct { x int `ab` }

// cannot use a (type struct { x int "a" }) as type struct { x int "ab" } in assignment b=a


package main

import(
    "reflect"
    "fmt"
)

func main() {
    type bigint int64

    var x bigint = 100
    println(x) 
    fmt.Println(reflect.TypeOf(x))

}

输出：
100
main.bigint


新类型不是原类型的别名，除拥有相同数据存储结构外，它们之间没有任何关系，不会持有原类型任何信息。除非目标类型是未命名类型，否则必须显式转换。

x := 1234
var b bigint = bigint(x) // 必须显式转换，除非是常量。 
var b2 int64 = int64(b)

var s myslice = []int{1, 2, 3} // 未命名类型，隐式转换。 
var s2 []int = s


声明及初始化

方法有几种：
[plain] view plain copy print?
var varName typeName             //①  
varName := new(typeName)         //②    同 var  *varName typeName = new(typeName) 
varName := typeName{[初始化值]}  //③  
varName := &typeName{[初始化值]} //④  同 var  *varName typeName = &typeName{[初始化值]}
注: ①③返回 typeName 类型变量；②④返回 *typeName 类型变量；③④[]可省略；若无初始化值，则默认为零值

初始化值可以分为两种:
 a. 有序: typeName{value1, value2, ...} 必须一一对应
 b. 无序: typeName{field1:value1, field2:value2, ...} 可初始化部分值

例:
[plain] view plain copy print?
type Person struct {  
  name string  
  age int  
}  
p := Person{"James", 23}  //有序  
p := Person{age:23}       //无序  


操作

  声明的struct与普通类型一样
  访问结构体中的一个变量名, 用 "." 来连接:
  varName.field 或 (*varName).field
  如操作上面 Person 结构体中的 age : p.age = 35
  也可以作为函数中的参数，返回值类型

如:

[plain] view plain copy print?
func funcName(varName1 typeName2[,varName2 typeName2, ...]) typeName {...}  
example code list
[plain] view plain copy print?

package main  
  
import "fmt"  
  
//1. 声明一个自定义类型名为 Person 的结构体  
type Person struct {  
    name string  
    age int  
}  
  
func main() {  
    //2. 初始化  
    var p1 Person  
    p2 := Person{}  
    p3 := Person{"James", 23}  
    p4 := Person{age:23}  
    fmt.Println(p1, p2, p3, p4)  
    p5 := new(Person)  
    p6 := &Person{}  
    p7 := &Person{"James", 23}  
    p8 := &Person{age:23}  
    fmt.Println(p5, p6, p7, p8)  
    /*********************************/  
    /*print result                   */  
    /*{ 0} { 0} {James 23} { 23}     */  
    /*&{ 0} &{ 0} &{James 23} &{ 23} */  
    /*********************************/  
      
    //3. 操作  
    p1.age = 50  
    p2.age = 25  
    if compareAge(p1, p2) {  
        fmt.Println("p1 is older than p2")  
    } else {  
        fmt.Println("p2 is older than p1")  
    }  
    /*********************************/  
    /*print result                   */  
    /*p1 is older than p2            */  
    /*********************************/  
}  
  
func compareAge(p1, p2 Person) bool {  
    if p1.age > p2.age {  
        return true  
    }  
    return false  
}  

输出：
{ 0} { 0} {James 23} { 23}
&{ 0} &{ 0} &{James 23} &{ 23}
p1 is older than p2

struct的内存布局

struct中的所有字段在内存是连续的，布局如下：




匿名字段：
    声明一个 struct1 可以包含已经存在的 struct2 或者go语言中内置类型作为内置字段，称为匿名字段，即只写了 typeName，无 varName，但是 typeName 不能重复。


匿名字段与面向对象程序语言中的继承
 
声明及初始化:
  如:
[plain] view plain copy print?
type Person struct {  
  name string  
  age int  
  addr string  
}  
type Employee struct {  
  Person          //匿名字段  
  salary int  
  int             //用内置类型作为匿名字段  
  addr string     //类似于重载  
}  
em1 := Employee{Person{"rain", 23, "qingyangqu"}, 5000, 100, "gaoxingqu"}  

操作
  访问方式也是通过 "." 来连接
  相同字段采用最外层优先访问，类似于重载
  em1.addr 访问的是 Employee 中最外层的 addr
  em1.Person.addr 访问的是 Employee 中 Person 中的 addr

example code list

[plain] view plain copy print?

package main  
  
import "fmt"  
  
type Person struct {  
    name string  
    age int  
    addr string  
}  
  
type Employee struct {  
    Person          //匿名字段  
    salary int  
    int             //用内置类型作为匿名字段  
    addr string     //类似于重载  
}  
  
func main() {  
    /*  
    var em1 Employee = Employee{}  
    em1.Person = Person{"rain", 23, "qingyangqu"}  
    em1.salary = 5000  
    em1.int = 100 //使用时注意其意义，此处无  
    em1.addr = "gaoxingqu"  
    */  
    //em1 := Employee{Person{"rain", 23, "qingyangqu"}, 5000, 100, "gaoxingqu"}  
    //初始化方式不一样，但是结果一样  
    em1 := Employee{Person:Person{"rain", 23, "qingyangqu"}, salary:5000, int:100, addr:"gaoxingqu"}  
    fmt.Println(em1)  
    /******************************************************/  
    /*print result                                        */  
    /*{{rain 23 qingyangqu} 5000 100 gaoxingqu}           */  
    /******************************************************/  
      
    fmt.Println("live addr(em1.addr) = ", em1.addr)  
    fmt.Println("work addr(em1.Person.addr) = ", em1.Person.addr)  
    em1.int = 200  //修改匿名字段的值  
    /******************************************************/  
    /*print result                                        */  
    /*live addr(em1.addr) =  gaoxingqu                    */  
    /*work addr(em1.Person.addr) =  qingyangqu            */  
    /******************************************************/  
}  

输出：
{{rain 23 qingyangqu} 5000 100 gaoxingqu}
live addr(em1.addr) =  gaoxingqu
work addr(em1.Person.addr) =  qingyangqu

strut与tag应用
    在处理json格式字符串的时候，经常会看到声明struct结构的时候，属性的右侧还有小米点括起来的内容。形如：

type User struct {
    UserId   int    `json:"user_id" bson:"user_id"`
    UserName string `json:"user_name" bson:"user_name"`
}

这个小米点里的内容是用来干什么的呢？

struct成员变量标签（Tag）说明

要比较详细的了解这个，要先了解一下golang的基础，在golang中，命名都是推荐都是用驼峰方式，并且在首字母大小写有特殊的语法含义：包外无法引用。但是由经常需要和其它的系统进行数据交互，例如转成json格式，存储到mongodb啊等等。这个时候如果用属性名来作为键值可能不一定会符合项目要求。

所以呢就多了小米点的内容，在golang中叫标签（Tag），在转换成其它数据格式的时候，会使用其中特定的字段作为键值。例如上例在转成json格式：

u := &User{UserId: 1, UserName: "tony"}
j, _ := json.Marshal(u)
fmt.Println(string(j))
// 输出内容：{"user_id":1,"user_name":"tony"}
如果在属性中不增加标签说明，则输出：

{"UserId":1,"UserName":"tony"}
可以看到直接用struct的属性名做键值。

其中还有一个bson的声明，这个是用在将数据存储到mongodb使用的。

struct成员变量标签（Tag）获取

那么当我们需要自己封装一些操作，需要用到Tag中的内容时，咋样去获取呢？这边可以使用反射包（reflect）中的方法来获取：

t := reflect.TypeOf(u)
field := t.Elem().Field(0)
fmt.Println(field.Tag.Get("json"))
fmt.Println(field.Tag.Get("bson"))
完整代码如下：

package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

func main() {
    type User struct {
        UserId   int    `json:"user_id" bson:"user_id"`     //多个Key使用空格进行分开，然后使用Get方法获取不同Key的值。
        UserName string `json:"user_name" bson:"user_name"`     
    }
    // 输出json格式
    u := &User{UserId: 1, UserName: "tony"}
    j, _ := json.Marshal(u)
    fmt.Println(string(j))
    // 输出内容：{"user_id":1,"user_name":"tony"}

    // 获取tag中的内容
    t := reflect.TypeOf(u)
    field := t.Elem().Field(0)
    fmt.Println(field.Tag.Get("json"))
    // 输出：user_id
    fmt.Println(field.Tag.Get("bson"))
    // 输出：user_id
}

输出：
{"user_id":1,"user_name":"tony"}
user_id
user_id

有意思的struct大小

我们定义一个struct，这个struct有3个字段，它们的类型有byte,int32以及int64,但是这三个字段的顺序我们可以任意排列，那么根据顺序的不同，一共有6种组合。

type user1 struct {
    b byte
    i int32
    j int64
}
type user2 struct {
    b byte
    j int64
    i int32
}
type user3 struct {
    i int32
    b byte
    j int64
}
type user4 struct {
    i int32
    j int64
    b byte
}
type user5 struct {
    j int64
    b byte
    i int32
}
type user6 struct {
    j int64
    i int32
    b byte
}
根据这6种组合，定义了6个struct，分别位user1，user2，…，user6，那么现在大家猜测一下，这6种类型的struct占用的内存是多少，就是unsafe.Sizeof()的值。

大家可能猜测1+4+8=13，因为byte的大小为1，int32大小为4，int64大小为8，而struct其实就是一个字段的组合，所以猜测struct大小为字段大小之和也很正常。

但是，但是，我可以明确的说，这是错误的。

为什么是错误的，因为有内存对齐存在，编译器使用了内存对齐，那么最后的大小结果就不一样了。现在我们正式验证下，这几种struct的值。

func main() {
    var u1 user1
    var u2 user2
    var u3 user3
    var u4 user4
    var u5 user5
    var u6 user6
    fmt.Println("u1 size is ",unsafe.Sizeof(u1))
    fmt.Println("u2 size is ",unsafe.Sizeof(u2))
    fmt.Println("u3 size is ",unsafe.Sizeof(u3))
    fmt.Println("u4 size is ",unsafe.Sizeof(u4))
    fmt.Println("u5 size is ",unsafe.Sizeof(u5))
    fmt.Println("u6 size is ",unsafe.Sizeof(u6))
}
从以上输出可以看到，结果是：

u1 size is  16
u2 size is  24
u3 size is  16
u4 size is  24
u5 size is  16
u6 size is  16
结果出来了（我的电脑的结果，Mac64位，你的可能不一样），4个16字节，2个24字节，既不一样，又不相同，这说明：

内存对齐影响struct的大小
struct的字段顺序影响struct的大小
综合以上两点，我们可以得知，不同的字段顺序，最终决定struct的内存大小，所以有时候合理的字段顺序可以减少内存的开销。

内存对齐会影响struct的内存占用大小，现在我们就详细分析下，为什么字段定义的顺序不同会导致struct的内存占用不一样。

在分析之前，我们先看下内存对齐的规则：

对于具体类型来说，对齐值=min(编译器默认对齐值，类型大小Sizeof长度)。也就是在默认设置的对齐值和类型的内存占用大小之间，取最小值为该类型的对齐值。我的电脑默认是8，所以最大值不会超过8.
struct在每个字段都内存对齐之后，其本身也要进行对齐，对齐值=min(默认对齐值，字段最大类型长度)。这条也很好理解，struct的所有字段中，最大的那个类型的长度以及默认对齐值之间，取最小的那个。
以上这两条规则要好好理解，理解明白了才可以分析下面的struct结构体。在这里再次提醒，对齐值也叫对齐系数、对齐倍数，对齐模数。这就是说，每个字段在内存中的偏移量是对齐值的倍数即可。

我们知道byte，int32，int64的对齐值分别为1，4，8，占用内存大小也是1，4，8。那么对于第一个structuser1，它的字段顺序是byte、int32、int64，我们先使用第1条内存对齐规则进行内存对齐，其内存结构如下。

bxxx|iiii|jjjj|jjjj
user1类型，第1个字段byte，对齐值1，大小1，所以放在内存布局中的第1位。

第2个字段int32，对齐值4，大小4，所以它的内存偏移值必须是4的倍数，在当前的user1中，就不能从第2位开始了，必须从第5位开始，也就是偏移量为4。第2，3，4位由编译器进行填充，一般为值0，也称之为内存空洞。所以第5位到第8位为第2个字段i。

第3字段，对齐值为8，大小也是8。因为user1前两个字段已经排到了第8位，所以下一位的偏移量正好是8，是第3个字段对齐值的倍数，不用填充，可以直接排列第3个字段，也就是从第9位到第16位为第3个字段j。

现在第一条内存对齐规则后，内存长度已经为16个字节，我们开始使用内存的第2条规则进行对齐。根据第二条规则，默认对齐值8，字段中最大类型长度也是8，所以求出结构体的对齐值位8，我们目前的内存长度为16，是8的倍数，已经实现了对齐。

所以到此为止，结构体user1的内存占用大小为16字节。

现在我们再分析一个user2类型，它的大小是24，只是调换了一下字段i和j的顺序，就多占用了8个字节，我们看看为什么？还是先使用我们的内存第1条规则分析。

bxxx|xxxx|jjjj|jjjj|iiii
按对齐值和其占用的大小，第1个字段b偏移量为0，占用1个字节，放在第1位。

第2个字段j，是int64，对齐值和大小都是8，所以要从偏移量8开始，也就是第9到16位为j，这也就意味着第2到8位被编译器填充。

目前整个内存布局已经偏移了16位，正好是第3个字段i的对齐值4的倍数，所以不用填充，可以直接排列，第17到20位为i。

现在所有字段对齐好了，整个内存大小为1+7+8+4=20个字节，我们开始使用内存对齐的第2条规则，也就是结构体的对齐，通过默认对齐值和最大的字段大小，求出结构体的对齐值为8。

现在我们的整个内存布局大小为20，不是8的倍数，所以我们需要进行内存填充，补足到8的倍数，最小的就是24，所以对齐后整个内存布局为

bxxx|xxxx|jjjj|jjjj|iiii|xxxx
所以这也是为什么我们最终获得的user2的大小为24的原因。
基于以上办法，我们可以得出其他几个struct的内存布局。

user3

iiii|bxxx|jjjj|jjjj
user4

iiii|xxxx|jjjj|jjjj|bxxx|xxxx
user5

jjjj|jjjj|bxxx|iiii
user6

jjjj|jjjj|iiii|bxxx
以上给出了答案，推到过程大家可以参考user1和user2试试。
