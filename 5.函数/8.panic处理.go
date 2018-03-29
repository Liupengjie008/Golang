panic处理
package main

import (
    "io"
    "log"
    "net/http"
)


func SimpleServer(w http.ResponseWriter, request *http.Request) {
    io.WriteString(w, "hello, world")
    panic("test test")
}

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
*/
func main() {
    http.HandleFunc("/test1", logPanics(SimpleServer))

    if err := http.ListenAndServe(":8088", nil); err != nil {
    }
}
/*
type HandlerFunc func(ResponseWriter, *Request)
*/
func logPanics(handle http.HandlerFunc) http.HandlerFunc {
    return func(writer http.ResponseWriter, request *http.Request) {
        defer func() {
            if x := recover(); x != nil {
                log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
            }
        }()
        handle(writer, request)
    }
}


然后在浏览器中输入 :
	http://localhost:8088

页面显示：
hello, world

控制台输出的panic信息：
2018/03/23 17:51:19 [[::1]:60659] caught panic: test test
