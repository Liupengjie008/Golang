反射获取基本类型
反射：可以在运行时动态获取变量的相关信息 
import "reflect"

a. reflect.TypeOf，获取变量的类型，返回reflect.Type类型 
b. reflect.ValueOf，获取变量的值，返回reflect.Value类型 
c. reflect.Value.Kind，获取变量的类别，返回一个常量 
d. reflect.Value.Interface()，转换成interface{}类型

package main 

import (
    "fmt"
    "reflect"
)

func main(){
    var x float64 = 3.4
    fmt.Println("type:", reflect.TypeOf(x))
    v := reflect.ValueOf(x)
    fmt.Println("value:", v)
    fmt.Println("type:", v.Type())
    fmt.Println("kind:", v.Kind())
    fmt.Println("value:", v.Float())

    fmt.Println(v.Interface())
    fmt.Printf("value is %5.2e\n", v.Interface())
    y := v.Interface().(float64)
    fmt.Println(y)
}


输出如下：
type: float64
value: 3.4
type: float64
kind: float64
value: 3.4
3.4
value is 3.40e+00
3.4