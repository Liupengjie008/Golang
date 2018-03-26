自定义io.Writer渲染
(index.html同模板渲染终端index.html保持不变)

我们看下myTemplate.Execute源码

func (t *Template) Execute(wr io.Writer, data interface{}) error {
    if err := t.escape(); err != nil {
        return err
    }
    return t.text.Execute(wr, data)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

所以我们实现Writer接口方法Write就可以自定义了

package main

import (
    // "os"
    "fmt"
    "html/template"
    "io"
    "net/http"
)

var myTemplate *template.Template

type Person struct {
    Name string
    age  string
}

type Result struct {
    output string
}
/*
自定义实现接口
type Writer interface {
    Write(p []byte) (n int, err error)
}
*/
func (p *Result) Write(b []byte) (n int, err error) {
    fmt.Println("called by template")
    p.output += string(b)
    p.output += "*"
    return len(b), nil
}

/*
func WriteString(w Writer, s string) (n int, err error) {
    if sw, ok := w.(stringWriter); ok {
        return sw.WriteString(s)
    }
    return w.Write([]byte(s))
}
*/

func userInfo(w http.ResponseWriter,r *http.Request) {
    p := Person{Name:"Murphy",age:"28"}

    resultWriter := &Result{}
    io.WriteString(resultWriter, "hello world\r\n")

    myTemplate.Execute(resultWriter,p)
    fmt.Println("render data:",resultWriter.output)
}

func initTemplate(fileName string) (err error){
    myTemplate,err  = template.ParseFiles(fileName)
    if err != nil{
        fmt.Println("parse file err:",err)
        return
    }
    return
}
/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
*/

func main() {
    initTemplate("./index.html")
    http.HandleFunc("/user/info", userInfo)
    err := http.ListenAndServe("0.0.0.0:8880", nil)
    if err != nil {
        fmt.Println("http listen failed")
    }
}


在浏览器中输入:
    http://localhost:8880/user/info 

    
然后在终端就会渲染出如下内容：
$ go run main.go 
called by template
called by template
called by template
called by template
called by template
called by template
render data: hello world
*<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        
        hello,*Murphy*
        *{Murphy 28}*
    </p>
</body>
</html>*
