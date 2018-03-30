模板渲染终端

模板替换 {{.字段名}}



main.go

package main

import (
    "fmt"
    "os"
    "text/template"
)

type Person struct {
    Name string
    age  string
}

func main() {
    t, err := template.ParseFiles("./index.html")
    if err != nil {
        fmt.Println("parse file err:", err)
        return
    }
    p := Person{Name: "Mary", age: "31"}
    if err := t.Execute(os.Stdout, p); err != nil {
        fmt.Println("There was an error:", err.Error())
    }
}




index.html

<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        hello,{{.Name}}
        {{.}}
    </p>
</body>
</html>






然后在终端就可以渲染输出如下：
$ go run main.go 
<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        hello,Mary
        {Mary 31}
    </p>
</body>
</html>