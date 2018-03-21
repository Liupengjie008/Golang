json反序列化为map
package main

import (
    "encoding/json"
    "fmt"
)


func testMap() (ret string, err error) {
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["username"] = "user1"
    m["age"] = 18
    m["sex"] = "man"

    data, err := json.Marshal(m)
    if err != nil {
        err = fmt.Errorf("json.marshal failed, err:", err)
        return
    }

    ret = string(data)
    return
}

func test2() {
    data, err := testMap()
    if err != nil {
        fmt.Println("test map failed, ", err)
        return
    }

    var m map[string]interface{}
    err = json.Unmarshal([]byte(data), &m)
    if err != nil {
        fmt.Println("Unmarshal failed, ", err)
        return
    }
    fmt.Println(m)
}



func main() {
    test2()
}




输出结果：
map[age:18 sex:man username:user1]