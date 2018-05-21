import "regexp"

regexp包实现了正则表达式搜索。

func Match(pattern string, b []byte) (matched bool, err error)
// Match检查b中是否存在匹配pattern的子序列。更复杂的用法请使用Compile函数和Regexp对象。
// pattern：要查找的正则表达式
// b：要在其中进行查找的 []byte
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现
package main

import (
	"fmt"
	"regexp"
)

func main() {
	ret, err := regexp.Match("H.*", []byte("Hello world"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}


func MatchString(pattern string, s string) (matched bool, err error)
// MatchString类似Match，但匹配对象是字符串。
// 判断在 s 中能否找到正则表达式 pattern 所匹配的子串
// pattern：要查找的正则表达式
// r：要在其中进行查找的字符串
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现

package main

import (
	"fmt"
	"regexp"
)

func main() {
	ret, err := regexp.MatchString("H.* ", "Hello World!")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}



func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
// MatchReader类似Match，但匹配对象是io.RuneReader。
// 判断在 r 中能否找到正则表达式 pattern 所匹配的子串
// pattern：要查找的正则表达式
// r：要在其中进行查找的 RuneReader 接口
// matched：返回是否找到匹配项
// err：返回查找过程中遇到的任何错误
// 此函数通过调用 Regexp 的方法实现

package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {
	r := bytes.NewReader([]byte("Hello World!"))
	ret, err := regexp.MatchReader("H.* ", r)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret)
}





