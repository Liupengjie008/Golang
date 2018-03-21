Elem反射操作基本类型
用来获取指针指向的变量，相当于： var a *int;

package main

import (
    "fmt"
    "reflect"
)

func main() {

    var b int = 1
    b = 200
    testInt(&b)
    fmt.Println(b)
}

//fv.Elem()用来获取指针指向的变量
func testInt(b interface{}) {
    val := reflect.ValueOf(b)
    val.Elem().SetInt(100)
    c := val.Elem().Int()

    fmt.Printf("get value  interface{} %d\n", c)
    fmt.Printf("string val:%d\n", val.Elem().Int())
}


输出结果：
get value  interface{} 100
string val:100
100