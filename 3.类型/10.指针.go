Go 语言指针

我们都知道，变量是一种使用方便的占位符，用于引用计算机内存地址。
Golang 支持指针类型 *T，指针的指针 **T，以及包含包名前缀的 *<package>.T。

• 默认值 nil，没有 NULL 常量。
• 操作符 "&" （取地址符） 取变量地址，"*" （取值符）透过指针访问目标对象。
• 不支持指针运算，不支持 "->" 运算符，直接用 "." 访问目标成员。

以下实例演示了变量在内存中地址：
package main

import "fmt"

func main() {
   var a int = 10   

   fmt.Printf("变量的地址: %x\n", &a  )
}
执行以上代码输出结果为：
变量的地址: c420012058
现在我们已经了解了什么是内存地址和如何去反问它。接下来我们将具体介绍指针。

什么是指针

一个指针变量可以指向任何一个值的内存地址它指向那个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：
var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

var ip *int        /* 指向整型*/      //声明一个int值得指针变量
var fp *float32    /* 指向浮点型 */
本例中这是一个指向 int 和 float32 的指针。

如何使用指针
指针使用流程：
定义指针变量。
为指针变量赋值。
访问指针变量中指向地址的值。
在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量的存储地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}

以上实例执行输出结果为：
a 变量的地址是: c420012058
ip 变量的存储地址: c420012058
*ip 变量的值: 20


直接用指针访问目标对象成员：
package main

import(
	"fmt"
)

func main() {
    type data struct{ a int }
    
    var d = data{1234}
    var p *data

    p = &d
    fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。 
}

输出:
0xc420012058, 1234

不能对指针做加减法等运算。

x := 1234
p := &x
p++      // Error: invalid operation: p += 1 (mismatched types *int and int)

可以在 unsafe.Pointer 和任意类型指针间进 转换。

package main

import(
	"fmt"
	"unsafe"
)

func main() {
    x := 0x12345678

    p := unsafe.Pointer(&x)     // *int -> Pointer
    n := (*[4]byte)(p)          // Pointer -> *[4]byte
    
    for i := 0; i < len(n); i++ {
        fmt.Printf("%X ", n[i])
	} 
	fmt.Println()
}

输出:
78 56 34 12


返回局部变量指针是安全的，编译器会根据需要将其分配在 GC Heap 上。

func test() *int {
    x := 100
    return &x   // 在堆上分配 x 内存。但在内联时，也可能直接分配在目标栈。 
}

将 Pointer 转换成 uintptr，可变相实现指针运算。
  
package main

import(
	"fmt"
	"unsafe"
)

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


Go 空指针     （nil）

当一个指针被定义后没有分配到任何变量时，它的值为 nil。

nil 指针也称为空指针。

nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。

一个指针变量通常缩写为 ptr。

查看以下实例：
package main

import "fmt"

func main() {
   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )
}
以上实例输出结果为：
ptr 的值为 : 0

空指针判断：

if(ptr != nil)     /* ptr 不是空指针 */
if(ptr == nil)    /* ptr 是空指针 */


Go 指针数组

在我们了解指针数组前，先看个实例，定义了长度为 3 的整型数组：
package main

import "fmt"

const MAX int = 3

func main() {

   a := []int{10,100,200}
   var i int

   for i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i, a[i] )
   }
}
以上代码执行输出结果为：
a[0] = 10
a[1] = 100
a[2] = 200
有一种情况，我们可能需要保存数组，这样我们就需要使用到指针。
以下声明了整型指针数组：
var ptr [MAX]*int;
ptr 为整型指针数组。因此每个元素都指向了一个值。以下实例的三个整数将存储在指针数组中：
package main

import "fmt"

const MAX int = 3

func main() {
   a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }
}
以上代码执行输出结果为：
a[0] = 10
a[1] = 100
a[2] = 200

如果一个指针变量存放的又是另一个指针变量的地址，则称这个指针变量为指向指针的指针变量。
当定义一个指向指针的指针变量时，第一个指针存放第二个指针的地址，第二个指针存放变量的地址：

指向指针的指针变量声明格式如下：
var ptr **int

以上指向指针的指针变量为整型。
访问指向指针的指针变量值需要使用两个 * 号，如下所示：
package main

import "fmt"

func main() {

   var a int
   var ptr *int
   var pptr **int

   a = 3000

   /* 指针 ptr 地址 */
   ptr = &a

   /* 指向指针 ptr 地址 */
   pptr = &ptr

   /* 获取 pptr 的值 */
   fmt.Printf("变量 a = %d\n", a )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
以上实例执行输出结果为：
变量 a = 3000
指针变量 *ptr = 3000
指向指针的指针变量 **pptr = 3000


Go 像函数传递指针参数

Go 语言允许向函数传递指针，志需要在函数定义的参数上设置为指针类型即可。
以下实例演示了如何向函数传递指针，并在函数调用后修改函数内的值，：
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int= 200

   fmt.Printf("交换前 a 的值 : %d\n", a )
   fmt.Printf("交换前 b 的值 : %d\n", b )

   /* 调用函数用于交换值
   * &a 指向 a 变量的地址
   * &b 指向 b 变量的地址
   */
   swap(&a, &b);

   fmt.Printf("交换后 a 的值 : %d\n", a )
   fmt.Printf("交换后 b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
   var temp int
   temp = *x    /* 保存 x 地址的值 */
   *x = *y      /* 将 y 赋值给 x */
   *y = temp    /* 将 temp 赋值给 y */
}
以上实例允许输出结果为：
交换前 a 的值 : 100
交换前 b 的值 : 200
交换后 a 的值 : 200
交换后 b 的值 : 100

指针类型转换

Go语言是不允许两个指针类型进行转换的。

我们一般使用*T作为一个指针类型，表示一个指向类型T变量的指针。为了安全的考虑，两个不同的指针类型不能相互转换，比如*int不能转为*float64。

func main() {
    i:= 10
    ip:=&i

    var fp *float64 = (*float64)(ip)

    fmt.Println(fp)
}
以上代码我们在编译的时候，会提示cannot convert ip (type *int) to type *float64，也就是不能进行强制转型。那如果我们还是需要进行转换怎么做呢？这就需要我们使用unsafe包里的Pointer了，下面我们先看看unsafe.Pointer是什么，然后再介绍如何转换。

Pointer

unsafe.Pointer是一种特殊意义的指针，它可以包含任意类型的地址，有点类似于C语言里的void*指针，全能型的。

package main

import(
	"fmt"
	"unsafe"
)

func main() {
    i:= 10
    ip:=&i

    var fp *float64 = (*float64)(unsafe.Pointer(ip))

    *fp = *fp * 3

    fmt.Println(i)
}

输出：
30
以上示例，我们可以把*int转为*float64,并且我们尝试了对新的*float64进行操作，打印输出i，就会发现i的址同样被改变。

以上这个例子没有任何实际的意义，但是我们说明了，通过unsafe.Pointer这个万能的指针，我们可以在*T之间做任何转换。

type ArbitraryType int
type Pointer *ArbitraryType
可以看到unsafe.Pointer其实就是一个*int,一个通用型的指针。

我们看下关于unsafe.Pointer的4个规则。

任何指针都可以转换为unsafe.Pointer
unsafe.Pointer可以转换为任何指针
uintptr可以转换为unsafe.Pointer
unsafe.Pointer可以转换为uintptr
前面两个规则我们刚刚已经演示了，主要用于*T1和*T2之间的转换，那么最后两个规则是做什么的呢？我们都知道*T是不能计算偏移量的，也不能进行计算，但是uintptr可以，所以我们可以把指针转为uintptr再进行便宜计算，这样我们就可以访问特定的内存了，达到对不同的内存读写的目的。

下面我们以通过指针偏移修改Struct结构体内的字段为例，来演示uintptr的用法。

package main

import(
	"fmt"
	"unsafe"
)

func main() {
    u:=new(user)
    fmt.Println(*u)

    pName:=(*string)(unsafe.Pointer(u))
    *pName="张三"

    pAge:=(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)))
    *pAge = 20

    fmt.Println(*u)
}
type user struct {
    name string
    age int
}

输出：
{ 0}
{张三 20}
以上我们通过内存偏移的方式，定位到我们需要操作的字段，然后改变他们的值。

第一个修改user的name值的时候，因为name是第一个字段，所以不用偏移，我们获取user的指针，然后通过unsafe.Pointer转为*string进行赋值操作即可。

第二个修改user的age值的时候，因为age不是第一个字段，所以我们需要内存偏移，内存偏移牵涉到的计算只能通过uintptr，所我们要先把user的指针地址转为uintptr，然后我们再通过unsafe.Offsetof(u.age)获取需要偏移的值，进行地址运算(+)偏移即可。

现在偏移后，地址已经是user的age字段了，如果要给它赋值，我们需要把uintptr转为*int才可以。所以我们通过把uintptr转为unsafe.Pointer,再转为*int就可以操作了。

这里我们可以看到，我们第二个偏移的表达式非常长，但是也千万不要把他们分段，不能像下面这样。

    temp:=uintptr(unsafe.Pointer(u))+unsafe.Offsetof(u.age)
    pAge:=(*int)(unsafe.Pointer(temp))
    *pAge = 20
逻辑上看，以上代码不会有什么问题，但是这里会牵涉到GC，如果我们的这些临时变量被GC，那么导致的内存操作就错了，我们最终操作的，就不知道是哪块内存了，会引起莫名其妙的问题。

小结

unsafe是不安全的，所以我们应该尽可能少的使用它，比如内存的操纵，这是绕过Go本身设计的安全机制的，不当的操作，可能会破坏一块内存，而且这种问题非常不好定位。

当然必须的时候我们可以使用它，比如底层类型相同的数组之间的转换；比如使用sync/atomic包中的一些函数时；还有访问Struct的私有字段时；该用还是要用，不过一定要慎之又慎。

还有，整个unsafe包都是用于Go编译器的，不用运行时，在我们编译的时候，Go编译器已经把他们都处理了。