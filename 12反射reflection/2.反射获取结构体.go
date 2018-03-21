反射获取结构体

package main

import (
    "fmt"
    "reflect"
)

type Student struct {
    Name  string
    Age   int
    Score float32
}


func test(b interface{}) {
    t := reflect.TypeOf(b)
    fmt.Println(t)

    v := reflect.ValueOf(b)
    fmt.Println(v)

    k := v.Kind()
    fmt.Println(k)


    iv := v.Interface()
    fmt.Println(iv)

    stu, ok := iv.(Student)
    if ok {
        fmt.Printf("%v %T\n", stu, stu)
    }
}

func main() {
    var a Student = Student{
        Name:  "stu01",
        Age:   18,
        Score: 92,
    }
    test(a)
}


输出结果：
main.Student
{stu01 18 92}
struct
{stu01 18 92}
{stu01 18 92} main.Student