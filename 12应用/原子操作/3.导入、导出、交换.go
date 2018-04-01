原子导入值
赋予变量新值，而不管它原来是什么值。

在原子地存储某个值的过程中，任何CPU都不会进行针对同一个值的读或写操作。
原子的值存储操作总会成功，因为它并不会关心被操作值的旧值是什么。
和CAS操作有着明显的区别。

package main

import (
    "fmt"
    "sync/atomic"
)

var val uint32

func main(){
    atomic.StoreUint32(&val, 200)
    fmt.Println(val)
}

输出结果：
200

函数原型：

atomic.StoreUint32(addr *uint32, val uint32)
atomic.StoreUint64(addr *uint64, val uint64)
atomic.StoreInt32(addr *int32, val int32)
atomic.StoreInt64(addr *int64, val int64)
atomic.StoreUintptr(addr *uintptr, val uintptr)
atomic.StorePointer(addr *unsafe.Pointer, val unsafe.Pointer)



原子导出值
导出变量当前的值。

v:= value 为变量v赋值，但要注意，在进行读取value的操作的过程中,其他对此值的读写操作是可以被同时进行的,那么这个读操作很可能会读取到一个只被修改了一半的数据.
so , 我们要使用sync/atomic代码包同样为我们提供了一系列的函数，以Load为前缀(载入)，来确保这样的糟糕事情发生。

package main

import (
    "fmt"
    "sync/atomic"
)

var val uint32 = 100

func main(){
    atomic.LoadUint32(&val)
    fmt.Println(val)
}

输出结果：
100

函数原型：

atomic.LoadUint32(addr *uint32) uint32
atomic.LoadUint64(addr *uint64) uint64
atomic.LoadInt32(addr *int32) int32
atomic.LoadInt64(addr *int64) int64
atomic.LoadUintptr(addr *uintptr) uintptr
atomic.LoadPointer(addr *unsafe.Pointer) unsafe.Pointer





原子交换值
赋予变量新值，同时返回变量的旧值。

与CAS操作不同，原子交换操作不会关心被操作的旧值。
它会直接设置新值
它会返回被操作值的旧值
此类操作比CAS操作的约束更少，同时又比原子载入操作的功能更强

package main

import (
    "fmt"
    "sync/atomic"
)

var old_val uint32 = 10
var new_val uint32 = 100

func main(){
    atomic.SwapUint32(&old_val,new_val )
    fmt.Println(new_val)
    fmt.Println(old_val)
}

输出结果：
100
100

函数原型：

atomic.SwapUint32(addr *uint32,  new uint32) old uint32
atomic.SwapUint64(addr *uint64,  new uint64) old uint64
atomic.SwapInt32(addr *int32,  new int32) old int32
atomic.SwapInt64(addr *int64,  new int64) old int64
atomic.SwapUintptr(addr *uintptr,  new uintptr) old uintptr
atomic.SwapPointer(addr *unsafe.Pointer,  new unsafe.Pointer) old unsafe.Pointer