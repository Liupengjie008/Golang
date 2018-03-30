模板渲染存储文件
(index.html同模板渲染终端index.html保持不变)

package main

import (
    "os"
    "fmt"
    "html/template"
    // "io"
    "net/http"
)

var myTemplate *template.Template

type Person struct {
    Name string
    age  string
}


func userInfo(w http.ResponseWriter,r *http.Request) {

    p := Person{Name:"Murphy",age:"28"}

    myTemplate.Execute(os.Stdout,p)
    file,err := os.OpenFile("demo.dat", os.O_CREATE|os.O_WRONLY, 0755)
    if err != nil {
        fmt.Println("open failed err:", err)
        return
    }
    myTemplate.Execute(file,p)
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


命令行：
$ go run main.go 

然后在浏览器中输入 :
	http://localhost:8880/user/info

命令行显示：
	<html>
	
	<head>
		<title>
		</title>
	</head>
	
	<body>
		<p>
			
			hello,Murphy
			{Murphy 28}
		</p>
	</body>
	</html>


命令行：
control+c

^Csignal: interrupt

$ ls
demo.dat	index.html	main.go

$ cat demo.dat 
<html>

<head>
    <title>
    </title>
</head>

<body>
    <p>
        
        hello,Murphy
        {Murphy 28}
    </p>
</body>

