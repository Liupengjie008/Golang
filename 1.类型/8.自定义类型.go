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

var a struct { x int `a` }
var b struct { x int `ab` }

// cannot use a (type struct { x int "a" }) as type struct { x int "ab" } in assignment b=a


可用 type 在全局或函数内定义新类型。
  
func main() {
    type bigint int64

    var x bigint = 100
    println(x) 
}

新类型不是原类型的别名，除拥有相同数据存储结构外，它们之间没有任何关系，不会持有原类型任何信息。除非目标类型是未命名类型，否则必须显式转换。

x := 1234
var b bigint = bigint(x) // 必须显式转换，除 是常量。 
var b2 int64 = int64(b)

var s myslice = []int{1, 2, 3} // 未命名类型，隐式转换。 
var s2 []int = s

