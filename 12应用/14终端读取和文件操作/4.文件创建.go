文件创建
创建文件的时候，一定要注意权限问题，一般默认的文件权限是 0666 
关于权限的相关内容，具体可以参考鸟叔p141 这里还是再回顾下，文件属性 r w x r w x r w x，第一位是文件属性，一般常用的 "-" 表示的是普通文件，"d"表示的是目录，golang里面使用os.Create创建文件的时候貌似只能使用0xxx的形式。
比如0666就表示创建了一个普通文件，文件所有者的权限，文件所属用户组的权限，以及其他人对此文件的权限都是110表示可读可写，不可执行。

/*
  递归创建目录
  os.MkdirAll(path string, perm FileMode) error

  path  目录名及子目录
  perm  目录权限位
  error 如果成功返回nil，如果目录已经存在默认什么都不做
*/
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.MkdirAll("./golang/log", 0777)
    if err != nil {
        fmt.Printf("%s", err)
    } else {
        fmt.Println("Create Directory OK!")
    }
}