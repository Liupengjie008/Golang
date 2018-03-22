从一个文件拷贝到另一个文件

package main

import (
    "fmt"
    "io"
    "os"
)

func CopyFile(dstName, srcName string) (writeen int64, err error) {
    src, err := os.Open(dstName)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer src.Close()

    dst, err := os.OpenFile(srcName, os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

func main() {
    CopyFile("/tmp/test", "/tmp/test_copy1")
    fmt.Println("copy done.")
}