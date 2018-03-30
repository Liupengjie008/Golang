int转json
package main

import (

    "fmt"
    "encoding/json"

)

func testInt() {
    var age = 100
    data, err := json.Marshal(age)
    if err != nil {
        fmt.Printf("json.marshal failed, err:", err)
        return
    }

    fmt.Printf("%s\n", string(data))
}

func main() {
    testInt()
    fmt.Println("----")
}




输出结果：
100
----