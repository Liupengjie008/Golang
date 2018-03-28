Golang的引用类型包括 slice、map 和 channel。它们有复杂的内部结构，除了申请内存外，还需要初始化相关属性。

内置函数 new 计算类型大小，为其分配零值内存，返回指针。而 make 会被编译器翻译 成具体的创建函数，由其分配内存和初始化成员结构，返回对象而非指针。

a := []int{0, 0, 0}		// 提供初始化表达式。
a[1] = 10

b := make([]int, 3)		// make slice
b[1] = 10

c := new([]int)			
c[1] = 10				// Error: invalid operation: c[1] (index of type *[]int)



引用类型：变量存储的是一个地址，这个地址存储最终的值。内存通常在堆上分配。通过GC回收。
        获取指针类型所指向的值，使用：* 。比如：var *p int, 使用*p获取p指向的值
		指针、slice、map、chan等都是引用类型。
		

new和make的区别

make 用来创建map、slice、channel
new用来创建值类型
	
new 和 make 均是用于分配内存：
	new 用于值类型和用户定义的类型，如自定义结构，make 用于内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针。它也可以被用于基本类型：v := new(int)。
	make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作。new() 是一个函数，不要忘记它的括号。