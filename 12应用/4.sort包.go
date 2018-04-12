import "sort"

sort包提供了排序切片和用户自定义数据集的函数。

func Ints(a []int)
Ints函数将a排序为递增顺序。
func IntsAreSorted(a []int) bool
IntsAreSorted检查a是否已排序为递增顺序。
func SearchInts(a []int, x int) int
SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。

代码实现：
package main

import (
	"fmt"
	"sort"
)

var IntSlice = []int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6}

func main() {
	sort.Ints(IntSlice)
	fmt.Println(IntSlice)
	b := sort.IntsAreSorted(IntSlice)
	fmt.Println(b)
	site := sort.SearchInts(IntSlice, 11)
	fmt.Println(site)
}

运行结果：
[0 1 2 3 4 5 6 7 8 9]
true
10

[]int 递减排序
package main

import (
	"fmt"
	"sort"
)

var IntSlice = []int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6, 10, 12}

func main() {
	sort.Ints(IntSlice)
	len := len(IntSlice)
	num := len / 2

	for i := 0; i < num; i++ {
		IntSlice[i], IntSlice[len-i-1] = IntSlice[len-i-1], IntSlice[i]
	}
	fmt.Println(IntSlice)
}



func Float64s(a []float64)
Float64s函数将a排序为递增顺序。
func Float64sAreSorted(a []float64) bool
Float64sAreSorted检查a是否已排序为递增顺序。
func SearchFloat64s(a []float64, x float64) int
SearchFloat64s在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。



func Strings(a []string)
Strings函数将a排序为递增顺序。
func StringsAreSorted(a []string) bool
StringsAreSorted检查a是否已排序为递增顺序。
func SearchStrings(a []string, x string) int
SearchStrings在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。s