字符串是不可变值类型，内部用指针指向 UTF-8 字节数组。

• 默认值是空字符串 ""。
• 用索引号访问某字节，如 s[i]。
• 不能用序号获取字节元素指针，&s[i] 非法。 
• 不可变类型，无法修改字节数组。
• 字节数组尾部不包含 NULL。

使用索引号访问字符 (byte)。

package main

func main(){
	s := "abc"
	println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)
}

输出:
true true true


使用 "`" 定义不做转义处理的原始字符串，支持跨行。

package main

func main(){
	s := `a
b\r\n\x00
c`
	println(s)
}

输出:
a
b\r\n\x00
c

连接跨行字符串时，"+" 必须在上一行末尾，否则导致编译错误。

s := "Hello, " +
     "World!"
s2 := "Hello, "
    + "World!"    // Error: invalid operation: + untyped string

支持用两个索引号 ([]) 返回子串。 串依然指向原字节数组，仅修改了指针和 度属性。

package main

import(
	"fmt"
)

func main(){
	s := "Hello, World!"

	s1 := s[:5]     // Hello
	s2 := s[7:]     // World!
	s3 := s[1:5]    // ello
	fmt.Println(s,s1,s2,s3)
}

单引号字符常量表示 Unicode Code Point， 持 \uFFFF、\U7FFFFFFF、\xFF 格式。 
对应 rune 类型，UCS-4。

package main

import(
	"fmt"
)

func main() {
    fmt.Printf("%T\n", 'a')
    var c1, c2 rune = '\u6211', '们'
    println(c1 == '我', string(c2) == "\xe4\xbb\xac") 
}

输出:
int32 // rune 是 int32 的别名 
true true

要修改字符串，可先将其转换成 []rune 或 []byte，完成后再转换为 string。无论哪种转换，都会重新分配内存，并复制字节数组。
 
package main

func main() {
    s := "abcd"
    bs := []byte(s)
    
    bs[1] = 'B'
    println(string(bs))

    u := "电脑"
    us := []rune(u)
    
    us[1] = '话'
    println(string(us))
}

输出:
aBcd
电话


for 循环遍历字符串时，也有 byte 和 rune 两种方式。

package main

import(
	"fmt"
)

func main() {
    s := "abc汉字"

    for i := 0; i < len(s); i++ {   // byte
        fmt.Printf("%c,", s[i]) 
    }

    fmt.Println()

    for _, r := range s {   // rune
        fmt.Printf("%c,", r)
	} 
	
	fmt.Println()
}


输出:
a,b,c,æ,±,,å,­
a,b,c,汉,字,

