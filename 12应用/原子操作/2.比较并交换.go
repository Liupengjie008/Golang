Compare And Swap 简称CAS （比较并交换）
先比较变量的值是否等于给定旧值，等于旧值的情况下才赋予新值，最后返回新值是否设置成功。

使用锁的做法趋于悲观
我们总假设会有并发的操作要修改被操作的值，并使用锁将相关操作放入临界区中加以保护。
使用CAS操作的做法趋于乐观 
总是假设被操作值未曾被改变（即与旧值相等），并一旦确认这个假设的真实性就立即进行值替换。

package main 

import(
    "fmt"
    "sync"
    "sync/atomic"
)

func main(){
    var sum uint32 = 100
    var wg sync.WaitGroup
    for i := uint32(0); i < 100; i++ {
        wg.Add(1)
        go func(t uint32) {
            defer wg.Done()
            atomic.CompareAndSwapUint32(&sum, 100, sum+1)
        }(i)
    }
    wg.Wait()
    fmt.Println(sum)
}

输出结果：
101

可以看到sum的值只改变了一次，只有当sum值为100的时候，CAS才将sum的值修改为了sum+1。

函数原型：

atomic.CompareAndSwapUint32(addr *uint32, old, new uint32) bool
atomic.CompareAndSwapUint64(addr *uint64, old, new uint64) bool
atomic.CompareAndSwapInt32(addr *int32, old, new int32) bool
atomic.CompareAndSwapInt64(addr *int64, old, new int64) bool
atomic.CompareAndSwapUintptr(addr *uintptr, old, new uintptr) bool
atomic.CompareAndSwapPointer(addr *unsafe.Pointer, old, new unsafe.Pointer) bool