反射调用结构体方法
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

func (s Student) Print(){
    fmt.Println(s)
}


func (s Student) Set(name string,age int,score float32){
    s.Age = age
    s.Name = name
    s.Score = score
}


func TestStruct(a interface{}) {
    val := reflect.ValueOf(a)
    kd := val.Kind()

    fmt.Println(val,kd)
    if kd != reflect.Struct {
        fmt.Println("expect struct")
        return
    }
    //获取字段数量
    fields:= val.NumField()
    fmt.Printf("struct has %d field\n",fields)
    //获取字段的类型
    for i:=0;i<fields;i++{
        fmt.Printf("%d %v\n",i,val.Field(i).Kind())
    }
    //获取方法数量
    methods:=val.NumMethod()
    fmt.Printf("struct has %d methods\n",methods)

    //反射调用的Print方法
    var params []reflect.Value
    val.Method(0).Call(params)


}

func main() {
    var a Student = Student{
        Name:  "stu01",
        Age:   18,
        Score: 92.8,
    }
    TestStruct(a)
    // fmt.Println(a)
}



输出结果：
{stu01 18 92.8} struct
struct has 3 field
0 string
1 int
2 float32
struct has 2 methods
{stu01 18 92.8}