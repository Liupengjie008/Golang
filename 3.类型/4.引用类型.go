Golang的引用类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。

内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译 成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。

a := []int{0, 0, 0}		// 提供初始化表达式。
a[1] = 10

b := make([]int, 3)		// makeslice
b[1] = 10

c := new([]int)			
c[1] = 10				// Error: invalid operation: c[1] (index of type *[]int)

