练习：
1.定义一个结构体
2.给结构体赋值
3.用反射获取结构体的 下标、结构体名称、类型、值
4.改变结构体的值



package main

import (
    "fmt"
    "reflect"
)

type T struct {
    A int
    B string
}

func main() {
    t := T{23, "skidoo"}
    s := reflect.ValueOf(&t).Elem()
    typeOfT := s.Type()
    for i := 0; i < s.NumField(); i++ {
        f := s.Field(i)
        fmt.Printf("%d: %s %s = %v\n", i,
            typeOfT.Field(i).Name, f.Type(), f.Interface())
    }
    s.Field(0).SetInt(77)
    s.Field(1).SetString("Sunset Strip")
    fmt.Println("t is now", t)
}


输出结果：
0: A int = 23
1: B string = skidoo
t is now {77 Sunset Strip}