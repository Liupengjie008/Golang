插入排序
思路分析：在要排序的一组数中，假设前面的数已经是排好顺序的，现在要把第n个数插到前面的有序数中，使得这n个数也是排好顺序的。如此反复循环，直到全部排好顺序。

代码实现：
package main

import (
	"fmt"
)

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func insertSort(sli []int) []int {
	len := len(sli)
	for i := 0; i < len; i++ {
		tmp := sli[i]
		//内层循环控制，比较并插入
		for j := i - 1; j >= 0; j-- {
			if tmp < sli[j] {
				//发现插入的元素要小，交换位置，将后边的元素与前面的元素互换
				sli[j+1], sli[j] = sli[j], tmp
			} else {
				//如果碰到不需要移动的元素，由于是已经排序好是数组，则前面的就不需要再次比较了。
				break
			}
		}
	}
	return sli
}

func main() {
	res := insertSort(sli)
	fmt.Println(res)
}
