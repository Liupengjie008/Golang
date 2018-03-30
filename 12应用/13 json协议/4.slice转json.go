slice转json
package main

import (

    "fmt"
    "encoding/json"

)


func testSlice() {
    var m map[string]interface{}
    var s []map[string]interface{}
    m = make(map[string]interface{})
    m["username"] = "user1"
    m["age"] = 18
    m["sex"] = "man"

    s = append(s, m)

    m = make(map[string]interface{})
    m["username"] = "user2"
    m["age"] = 29
    m["sex"] = "female"
    s = append(s, m)

    data, err := json.Marshal(s)
    if err != nil {
        fmt.Printf("json.marshal failed, err:", err)
        return
    }

    fmt.Printf("%s\n", string(data))
}

func main() {
	testSlice()
	fmt.Println("--------")
}




输出结果：
[{"age":18,"sex":"man","username":"user1"},{"age":29,"sex":"female","username":"user2"}]
--------