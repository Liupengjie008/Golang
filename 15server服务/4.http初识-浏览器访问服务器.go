http初识-浏览器访问服务器
package main

import (
    "fmt"
    "net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handle hello")
    fmt.Fprintf(w, "----- hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("handle login")
    fmt.Fprintf(w, "----- login ")
}

func main() {
    http.HandleFunc("/", Hello)
    http.HandleFunc("/user/login", login)
    err := http.ListenAndServe("0.0.0.0:8880", nil)
    if err != nil {
        fmt.Println("http listen failed")
    }
}


然后在浏览器中输入 :
	http://localhost:8880

页面显示：
----- hello  

终端输出：
handle hello
handle hello


然后在浏览器中输入 :
	http://localhost:8880/user/login

页面显示：
----- login 

终端输出：
handle login
handle hello
