Golang 通过 cgo，可在 Go 和 C/C++ 代码间相互调用。受 CGO_ENABLED 参数限制。
     
package main

/*
   #include <stdio.h>
   #include <stdlib.h>
   void hello() {
        printf("Hello, World!\n");
} */
import "C"

func main() {
    C.hello()
}


调试 cgo 代码是件很麻烦的事，建议单独保存到 .c 文件中。这样可以将其当做独立的 C 程序进行调试。

test.h
#ifndef __TEST_H__
#define __TEST_H__

void hello();

#endif


test.c
#include <stdio.h>
#include "test.h"

void hello() {
    printf("Hello, World!\n");
}

#ifdef __TEST__        // 避免和 Go bootstrap main 冲突。

int main(int argc, char *argv[]) {
    hello();
    return 0; 
}

#endif


main.go
package main

/*
   #include "test.h"
*/
import "C"

func main() {
    C.hello()
}


编译和调试 C，只需在命令行提供宏定义即可。 

$ gcc -g -D__TEST__ -o test test.c


由于 cgo 仅扫描当前目录，如果需要包含其他 C 项 ，可在当前目录新建一个 C文件， 然后用 #include 指令将所需的 .h、.c 都包含进来，记得在 CFLAGS 中使用 "-I" 参数指定原路径。某些时候，可能还需指定 "-std" 参数。


Flags
可使用 #cgo 命令定义 CFLAGS、LDFLAGS 等参数，自动合并多个设置。

/*
   #cgo CFLAGS: -g
   #cgo CFLAGS: -I./lib -D__VER__=1
   #cgo LDFLAGS: -lpthread
   #include "test.h"
*/
import "C"


可设置 GOOS、GOARCH 编译条件，空格表示 OR，逗号 AND，感叹号 NOT。 

#cgo windows,386 CFLAGS: -I./lib -D__VER__=1

 

DataType 数据类型对应关系。
 
   C                    cgo                  sizeof
--------------------+--------------------+--------------------------------
char                    C.char               1
signed char             C.schar              1
unsigned char           C.uchar              1
short                   C.short              2  
unsigned short          C.ushort             2 
int                     C.int                4
unsigned int            C.uint               4
long                    C.long               4 或 8
unsigned long           C.ulong              4 或 8
long long               C.longlong           8 
unsinged long long      C.ulonglong          8
float                   C.float              4
double                  C.double             8
void*                   unsafe.Pointer
char*                   *C.char
size_t                  C.size_t
NULL                    nil
          

可将 cgo 类型转换为标准 Go 类型。

/*
    int add(int x, int y) {
        return x + y; 
    }
*/
import "C"

func main() {
    var x C.int = C.add(1, 2)
    var y int = int(x)
    fmt.Println(x, y)
}


String 字符串转换函数。
 
/*
    #include <stdio.h>
    #include <stdlib.h>
    void test(char *s) {
        printf("%s\n", s);
    }

    char* cstr() {
        return "abcde";
    } 
*/
import "C"

func main() {
    s := "Hello, World!"

    cs := C.CString(s)    // 该函数在 C heap 分配内存，需要调用 free 释放。
    defer C.free(unsafe.Pointer(cs))     // #include <stdlib.h>

    C.test(cs)
    cs = C.cstr()

    fmt.Println(C.GoString(cs))
    fmt.Println(C.GoStringN(cs, 2))
    fmt.Println(C.GoBytes(unsafe.Pointer(cs), 2))
}


输出:
Hello, World!
abcde
ab
[97 98]


C.malloc/free 分配 C heap 内存。

/*
   #include <stdlib.h>
*/
import "C"

func main() {
    m := unsafe.Pointer(C.malloc(4 * 8))
    defer C.free(m)      // 注释释放内存。

    p := (*[4]int)(m)    // 转换为数组指针。
    for i := 0; i < 4; i++ {
        p[i] = i + 100 
    }

    fmt.Println(p)
}
 
输出:
&[100 101 102 103]



Struct/Enum/Union
对 struct、enum 支持良好，union 会被转换成字节数组。如果没使用 typedef 定义，那么必须添加 struct_、enum_、union_ 前缀。 

struct
/*
    #include <stdlib.h>

    struct Data {
        int x;
    };
    
    typedef struct {
        int x;
    } DataType;
    
    struct Data* testData() {
        return malloc(sizeof(struct Data));
    }
       
    DataType* testDataType() {
        return malloc(sizeof(DataType));
    } 
*/
import "C"

func main() {
    var d *C.struct_Data = C.testData()
    defer C.free(unsafe.Pointer(d))

    var dt *C.DataType = C.testDataType()
    defer C.free(unsafe.Pointer(dt))

    d.x = 100
    dt.x = 200

    fmt.Printf("%#v\n", d)
    fmt.Printf("%#v\n", dt)
}

输出:
&main._Ctype_struct_Data{x:100}
&main._Ctype_DataType{x:200}


enum 
/*
    enum Color { BLACK = 10, RED, BLUE };
    typedef enum { INSERT = 3, DELETE } Mode;
*/
import "C"

func main() {
    var c C.enum_Color = C.RED
    var x uint32 = c
    fmt.Println(c, x)

    var m C.Mode = C.INSERT
    fmt.Println(m)
}
 

union
/*
    #include <stdlib.h>
    union Data {
        char x;
        int y; 
    };

    union Data* test() {
        union Data* p = malloc(sizeof(union Data));
        p->x = 100;
        return p;
    } 
*/
import "C"

func main() {
    var d *C.union_Data = C.test()
    defer C.free(unsafe.Pointer(d))

    fmt.Println(d)
}

输出:
&[100 0 0 0]



Export
导出 Go 函数给 C 调用，须使用 "//export" 标记。建议在独立头文件中声明函数原型，避免 "duplicate symbol" 错误。 

main.go
package main

import "fmt"

/*
   #include "test.h"
*/
import "C"

//export hello

func hello() {
    fmt.Println("Hello, World!\n")
}

func main() {
    C.test()
}


test.h
#ifndef __TEST_H__
#define __TEST_H__

extern void hello();
void test();

#endif



test.c
#include <stdio.h>
#include "test.h"

void test() {
    hello();
}




Shared Library
在 cgo 中使用 C 共享库。

test.h
#ifndef __TEST_HEAD__
#define __TEST_HEAD__

int sum(int x, int y);

#endif


test.c
#include <stdio.h>
#include <stdlib.h>
#include "test.h"

int sum(int x, int y)
{
    return x + y + 100;
}

编译成 .so 或 .dylib。

$ gcc -c -fPIC -o test.o test.c
$ gcc -dynamiclib -o libtest.dylib test.o


将共享库和头文件拷贝到 Go 项目目录。 

main.go
package main

/*
   #cgo CFLAGS: -I.
   #cgo LDFLAGS: -L. -ltest
   #include "test.h"
*/
import "C"

func main() {
    println(C.sum(10, 20))
}

输出:
$ go build -o test && ./test
130


编译成功后可用 ldd 或 otool 查看动态库使用状态。 静态库使用方法类似。