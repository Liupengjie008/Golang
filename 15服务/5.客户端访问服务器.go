客户端访问服务器
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    res, err := http.Get("https://www.baidu.com/")
    if err != nil {
        fmt.Println("get err:", err)
        return
    }

    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println("get data err:", err)
        return
    }

    fmt.Println(string(data))
}




$ go run main.go
输出结果： 
<html>
<head>
	<script>
		location.replace(location.href.replace("https://","http://"));
	</script>
</head>
<body>
	<noscript><meta http-equiv="refresh" content="0;url=http://www.baidu.com/"></noscript>
</body>
</html>