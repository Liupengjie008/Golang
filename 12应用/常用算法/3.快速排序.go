快速排序  
思路分析：选择一个基准元素，通常选择第一个元素或者最后一个元素。通过一趟扫描，将待排序列分成两部分，一部分比基准元素小，一部分大于等于基准元素。此时基准元素在其排好序后的正确位置，然后再用同样的方法递归地排序划分的两部分。

代码实现：

package main

import (
	"fmt"
)

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func quickSort(sli []int) []int {
	//先判断是否需要继续进行
	len := len(sli)
	if len <= 1 {
		return sli
	}
	//选择第一个元素作为基准
	base_num := sli[0]
	//遍历除了标尺外的所有元素，按照大小关系放入两个数组内
	//初始化两个数组
	left_sli := []int{}  //小于基准的
	right_sli := []int{} //大于基准的
	for i := 1; i < len; i++ {
		if base_num > sli[i] {
			//放入左边数组
			left_sli = append(left_sli, sli[i])
		} else {
			//放入右边
			right_sli = append(right_sli, sli[i])
		}
	}

	//再分别对左边和右边的数组进行相同的排序处理方式递归调用这个函数
	left_sli = quickSort(left_sli)
	right_sli = quickSort(right_sli)

	//合并
	left_sli = append(left_sli, base_num)
	return append(left_sli, right_sli...)
}
func main() {
	res := quickSort(sli)
	fmt.Println(res)
}
