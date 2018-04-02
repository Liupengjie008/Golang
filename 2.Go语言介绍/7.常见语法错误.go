
8. map 未初始化 直接使用: Using "nil" Slices and Maps------interfaces, functions, pointers, maps, slices, and channels

错误：



报错信息：



正确：

9. map 只有 len操作， 没有 cap 操作: Map Capacity


错误：



报错信息：



正确：





11. array 是值类型, 作为参数其值不会被改变, 形参复制了一份数据给实参; 如果确实需要改变, 需要使用数组指针 或者 slice切片 作为形参: Array Function Arguments


错误：



报错信息：



正确：


11. for range 遍历返回两个参数k, v, 不是一个参数: Unexpected Values in Slice and Array "range" Clauses

12. 慎用 多维数组、多维slice，需要分步骤make才能完成: Slices and Arrays Are One-Dimensional

13. 请判断 map 中 key 是否存在， 否则可能造成未知错误, 因为key不存在取出来也是有值的, 没有异常！哈哈，GO 就是这么逆天！不知道没有完善的异常处理机制是不是GO最大的败笔 : Accessing Non-Existing Map Keys

14. string是值类型，不可迭代，如果需要迭代可以转化为 []byte 进一步处理: Strings Are Immutable

15. string 可以用下标引用，不能用下表进行修改，string存储使用 UTF8，string的处理建议使用 []rune 转化为Unicode之后进行，或者使用GO的UTF8包中的接口: Strings and Index Operator

16. 请检查你的输入string是否为utf8格式, 然后继续处理，这样可以防止意外的错误: Strings Are Not Always UTF8 Text

17.  获取字符串长度请最好使用 UTF8包的函数 统计Unicode字符数，GO默认统计的是UTF8编码的字符串字节长度，即实际内存中存储的长度: String Length

18. array、slice、map等多行赋值需要添加 , 否则编译失败: Missing Comma In Multi-Line Slice, Array, and Map Literals

19. log.Fatal and log.Panic Do More Than Log 直接宣告Panic 退出: log.Fatal and log.Panic Do More Than Log

20. GO 内置的数据结构， 比如 map，非线程安全，需要自己处理: Built-in Data Structure Operations Are Not Synchronized

21. 迭代 string 会自动转化为 Unicode 字符输出, 请保证输入的string为UTF8格式: Iteration Values For Strings in "range" Clauses

22. map 的输出是 无序 的: Iterating Through a Map Using a "for range" Clause

23. switch-case 默认有break， 如果需要取消 ， 请使用  fallthrough: Fallthrough Behavior in "switch" Statements

24. GO 语言 没有 ++/-- 运算符: Increments and Decrements

25. GO 语言中, ^ 既是按位取反操作符，也是异或的操作符，没有 ~ 操作符: Bitwise NOT Operator

26.  注意按位操作的顺序，请使用括号明确表示: Operator Precedence Differences

27. struct 小写变量不会被序列化: Unexported Structure Fields Are Not Encoded

28. 主线程不会主动等待所有 Goroutine 完成: App Exits With Active Goroutines

29. 向无缓存的Channel发送消息，只要目标接收者准备好就会立即返回（与此相反，从无缓存buffer接收数据则会一直阻塞直到有数据）: Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready

30. 向已关闭的Channel发送会引起Panic(从一个关闭的channel接收是安全的。在接收状态下的 ok的返回值将被设置为 false)，就是说，不要直接粗暴的关闭channel，最好需要在够routine中 select close信号: Sending to an Closed Channel Causes a Panic

31. 使用"nil" Channels是不对的，需要先用make进行初始化，否则会抛出 deadlock 异常: Using "nil" Channels

32. 传值方法的接收者无法修改原有的值: Methods with Value Receivers Can't Change the Original Value

33. 正确关闭HTTP的响应，防止内存泄露或者PANIC：Closing HTTP Response Body

34. 关闭HTTP的连接(req.Close = true 或者 req.Header.Add("Connection", "close") 或者 你也可以取消http的全局连接复用): Closing HTTP Connections

35. 可以使用 == 、reflect.deepEquals 比较Structs, Arrays, Slices, and Maps：Comparing Structs, Arrays, Slices, and Maps

36. 从Panic中恢复（recover()的调用仅当它在defer函数中被直接调用时才有效）：Recovering From a Panic

37. 在Slice, Array, and Map "range"语句中更新引用元素的值是无效的（在“range”语句中生成的数据的值是真实集合元素的拷贝。它们不是原有元素的引用，如果你需要更新原有集合中的数据，使用索引操作符来获得数据）：Updating and Referencing Item Values in Slice, Array, and Map "range" Clauses

38. slice可以直接返回，所以如果不是要对原来的slice直接操做，请copy之后再返回："Hidden" Data in Slices

39. Slice的数据“毁坏”（slice会被直接引用，请注意数据保护）：Slice Data "Corruption"

40. "走味的"Slices（同上）："Stale" Slices

41. 当你通过把一个现有（非interface）的类型定义为一个新的类型时，新的类型不会继承现有类型的方法，如果你确实需要原有类型的方法，你可以定义一个新的struct类型，用匿名方式把原有类型嵌入其中：Type Declarations and Methods

42. 从"for switch"和"for select"代码块中跳出，如果无法使用“return”声明的话，那就为外部循环定义一个标签是另一个好的选择：Breaking Out of "for switch" and "for select" Code Blocks

43. "for"声明中的迭代变量和闭包：Iteration Variables and Closures in "for" Statements

44. Defer函数调用参数的求值，被defer的函数的参数会在defer声明时求值（而不是在函数实际执行时）：Deferred Function Call Argument Evaluation

45. 被Defer的函数调用执行，在函数之内执行，而不是代码块内执行：Deferred Function Call Execution

46. 失败的类型断言：Failed Type Assertions

47. 阻塞的Goroutine和资源泄露：Blocked Goroutines and Resource Leaks

48. 使用指针接收方法的值的实例, 并不是所有的变量是可取址的。Map的元素就不是。通过interface引用的变量也不是：Using Pointer Receiver Methods On Value Instances

49. 更新Map的值， 如果你有一个struct值的map，你无法更新单个的struct值： Updating Map Value Fields

50. "nil" Interfaces和"nil" Interfaces的值（返回 interface时请注意 nil）: "nil" Interfaces and "nil" Interfaces Values

51. 栈和堆变量位置不确定：Stack and Heap Variables

52. GOMAXPROCS, 并发, 和并行（默认情况下，Go仅使用一个执行上下文/OS线程（在当前的版本）。这个数量可以通过设置 GOMAXPROCS来提高）： GOMAXPROCS, Concurrency, and Parallelism

53. 读写操作的重排顺序： Read and Write Operation Reordering

54. 优先调度（有可能会出现这种情况，一个无耻的goroutine阻止其他goroutine运行。当你有一个不让调度器运行的 for循环时，这就会发生）：Preemptive Scheduling 








