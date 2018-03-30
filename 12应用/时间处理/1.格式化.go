package main

import(
    "fmt"
    "time"
)

func main() {
    
    // 当前时间戳
    now := time.Now().Unix()
    fmt.Println(now)

    // 当前格式化时间 
    fmt.Println(time.Now().Format("2006-01-02 15:04:05")) 
    // 这是个奇葩,必须是这个时间点, 据说是go诞生之日, 记忆方法:6-1-2-3-4-5

    // 时间戳转str格式化时间
    str_time1 := time.Unix(0, 0).Format("2006-01-02 15:04:05")
    fmt.Println(str_time1)
    str_time2 := time.Unix(1522393808, 0).Format("2006年01月02日 15时04分05秒")
    fmt.Println(str_time2)
    // str_time3 := time.Unix("1522393808",0).Format("2006-01-02 15:04:05")
    // cannot use "1522393808" (type string) as type int64 in argument to time.Unix
    str_time4 := time.Unix(1522393808, 0).Format("06-01-02 15:04:05")
    fmt.Println(str_time4)
    str_time5 := time.Unix(1522393808, 0).Format("01-02 15:04")
    fmt.Println(str_time5)

    // str格式化时间转时间戳 
    // 方法一   2018-03-30 15:24:59
    the_time := time.Date(2018, 3, 30, 15, 24, 59, 0, time.Local)
    unix_time := the_time.Unix()
    fmt.Println(unix_time)
    fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))
    // 方法二 , 使用time.Parse
    /*
        返回的不是本地时间, 而是 UTC , 会自动加8小时.
    */
    the_time, err := time.Parse("2006-01-02 15:04:05", "2018-03-30 15:24:59")
    if err == nil {
        unix_time := the_time.Unix()
        fmt.Println(unix_time)
        fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))		
    }
    // 使用time.ParseInLocation
    the_time, err = time.ParseInLocation("2006-01-02 15:04:05", "2018-03-30 15:24:59",time.Local)
    if err == nil {
        unix_time := the_time.Unix()
        fmt.Println(unix_time)
        fmt.Println(time.Unix(unix_time,0).Format("2006-01-02 15:04:05"))		
    }

    // 格式化当前时间
    lasttime := time.Now().Format("2006-01-02 15:04:05")
    fmt.Println(lasttime)

}


输出结果：
$ go run main.go 
1522395496
2018-03-30 15:38:16
1970-01-01 08:00:00
2018年03月30日 15时10分08秒
18-03-30 15:10:08
03-30 15:10
1522394699
2018-03-30 15:24:59
1522423499
2018-03-30 23:24:59
1522394699
2018-03-30 15:24:59
2018-03-30 15:38:16

