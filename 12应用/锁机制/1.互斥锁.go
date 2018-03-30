golang中sync包实现了两种锁Mutex （互斥锁）和RWMutex（读写锁），其中RWMutex是基于Mutex实现的，只读锁的实现使用类似引用计数器的功能．

type Mutex

    func (m *Mutex) Lock()

    func (m *Mutex) Unlock()

type RWMutex

    func (rw *RWMutex) Lock()

    func (rw *RWMutex) RLock()

    func (rw *RWMutex) RLocker() Locker

    func (rw *RWMutex) RUnlock()

    func (rw *RWMutex) Unlock()

1、互斥锁     （用于基本上全是写入操作的应用）

    go mutex是互斥锁，只有Lock和Unlock两个方法，在这两个方法之间的代码不能被多个goroutins同时调用到。

其中Mutex为互斥锁，Lock()加锁，Unlock()解锁，使用Lock()加锁后，便不能再次对其进行加锁，直到利用Unlock()解锁对其解锁后，才能再次加锁．适用于读写不确定场景，即读写次数没有明显的区别，并且只允许只有一个读或者写的场景，所以该锁叶叫做全局锁。

func (m *Mutex) Unlock()用于解锁m，如果在使用Unlock()前未加锁，就会引起一个运行错误．已经锁定的Mutex并不与特定的goroutine相关联，这样可以利用一个goroutine对其加锁，再利用其他goroutine对其解锁。

互斥锁只能锁定一次，当在解锁之前再次进行加锁，便会死锁状态，如果在加锁前解锁，便会报错“panic: sync: unlock of unlocked mutex”