单元测试
1. 文件名必须以_test.go结尾
2. 测试函数必须以Test开头
3. 使用go test执行单元测试

目录结构：
	test
	  |
	   —— calc.go
	  |
	   —— calc_test.go



calc.go

package main

func add(a, b int) int {
    return a + b
}

func sub(a, b int) int {
    return a - b
}


calc_test.go

package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    r := add(2, 4)
    if r != 6 {
        t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 6, r)
    }
    t.Logf("test add succ")
}



输出结果：
$ cd test/

$ ls
calc.go		calc_test.go

$ go test 
PASS
ok  	go_dev/test/test	0.007s