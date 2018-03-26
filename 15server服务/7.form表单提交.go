form表单提交
package main

import (
    "io"
    // "log"
    "net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></body></html>`


func FormServer(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    switch request.Method {
    case "GET":
        io.WriteString(w, form)
    case "POST":
        request.ParseForm()
        io.WriteString(w, request.Form["in"][1])
        io.WriteString(w, "\n")
        io.WriteString(w, request.FormValue("in"))
    }
}

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
*/
func main() {
    http.HandleFunc("/test2", FormServer)
    if err := http.ListenAndServe(":8088", nil); err != nil {
    }
}


然后在浏览器中输入 :
	http://localhost:8088/test2


页面显示：
	两个输入框
	第一个输入框输入：123
	第二个输入框输入：456
	提交表单

页面显示：
	456 123