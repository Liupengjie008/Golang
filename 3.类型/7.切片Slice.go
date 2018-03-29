需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
2. 切片的长度可以改变，因此，切片是一个可变的数组。
3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。 
4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= (array)，其中array是slice引用的数组。
5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
6. 如果 slice == nil，那么 len、cap 结果都等于 0。


切片初始化
     var slice []int = arr[start:end]              包含start到end之间的元素，但不包含end 
     var slice []int = arr[0:end]                   可以简写为 var slice []int=arr[:end]
     var slice []int = arr[start:len(arr)]         可以简写为 var slice[]int = arr[start:]
     var slice []int = arr[0, len(arr)]             可以简写为 var slice[]int = arr[:]
     如果要切片最后一个元素去掉，可以这么写： Slice = slice[:len(slice)-1]

通过make来创建切片
    var slice []type = make([]type, len)
    slice  := make([]type, len)
	slice  := make([]type, len, cap)

读写操作实际目标是底层数组，只需注意索引号的差别。

	data := [...]int{0, 1, 2, 3, 4, 5}
	
	s := data[2:4]
	s[0] += 100
	s[1] += 200
	
	fmt.Println(s)
	fmt.Println(data)

输出:
	[102 203]
	[0 1 102 203 4 5]
	
可直接创建 slice 对象，自动分配底层数组。
 
	s1 := []int{0, 1, 2, 3, 8: 100}  	// 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))
	
	s2 := make([]int, 6, 8)    			// 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))
	
	s3 := make([]int, 6)  				// 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
	
输出:
	[0 1 2 3 0 0 0 0 100]   9  9 
	[0 0 0 0 0 0] 			6  8 
	[0 0 0 0 0 0]           6  6
	
使用 make 动态创建 slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
 
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, 获取底层数组元素指针。 
	*p += 100
	
	fmt.Println(s)
	
	输出:
	[0 1 102 3]
	
	
	至于 [][]T，是指元素类型为 []T 。
	
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	
	
	可直接修改 struct array/slice 成员。
	
	d := [5]struct {
		x int
	}{}
	
	s := d[:]
	
	d[1].x = 10
	s[2].x = 20
	
	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
	
	输出:
	[{0} {10} {20} {0} {0}]
	0x20819c180, 0x20819c180
	
		
用append内置函数操作切片（切片追加）
    slice = append(slice, 10)

    var a = []int{1,2,3}
    var b = []int{4,5,6}
    a = append(a, b…)

	append ：向 slice 尾部添加数据，返回新的 slice 对象。

	s := make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	
	s2 := append(s, 1)
	fmt.Printf("%p\n", &s2)
	
	fmt.Println(s, s2)
	
	输出:
	0x210230000
	0x210230040
	[] [1]
	
	超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]
	
	s = append(s, 100, 200)       // 一次 append 两个值，超出 s.cap 限制。
	
	fmt.Println(s, data)	      // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0])  // 比对底层数组起始指针。
	
	输出:
	[0 1 100 200] [0 1 2 3 4 0 0 0 0 0 0]
	0x20819c180 0x20817c0c0
	
	从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
	通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，以减少内存分配和数据复制开销。或初始化足够长的 len 属性，改用索引号进行操作。及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
	 
	s := make([]int, 0, 1)
	c := cap(s)
	
	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c { 
			fmt.Printf("cap: %d -> %d\n", c, n) c=n
		} 
	}
	
	输出:
	cap: 1 -> 2
	cap: 2 -> 4
	cap: 4 -> 8
	cap: 8 -> 16
	cap: 16 -> 32
	cap: 32 -> 64

切片拷贝    
    s1 := []int{1,2,3,4,5}

    s2 := make([]int, 10)

    copy(s2, s1)

    s3 := []int{1,2,3}

    s3 = append(s3, s2…)

	s3 = append(s3, 4,5,6)

copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。

	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	s := data[8:]
	s2 := data[:5]
	
	copy(s2, s)          // dst:s2, src:s
	
	fmt.Println(s2)
	fmt.Println(data)
	 
	输出:
	[8 9 2 3 4]
	[8 9 2 3 4 5 6 7 8 9]
	
	应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。

for range 遍历切片
    for index, val := range slice {
    }

切片resize（调整大小）
    var a = []int {1,3,4,5}
    b := a[1:2]
    b = b[0:3]



数组和切片的内存布局
     

字符串和切片（string与slice）
    string底层就是一个byte的数组，因此，也可以进行切片操作。
    str := “hello world”
    s1 := str[0:5]
    fmt.Println(s1)

    s2 := str[5:]
    fmt.Println(s2)
    string本身是不可变的，因此要改变string中字符。需要如下操作：
    str := “hello world”
    s := []byte(str)   //中文字符需要用[]rune(str)
    s[0] = ‘o’
    str = string(s)

