选择排序 
思路分析：在要排序的一组数中，选出最小的一个数与第一个位置的数交换。然后在剩下的数当中再找最小的与第二个位置的数交换，如此循环到倒数第二个数和最后一个数比较为止。

代码实现：
package main

import (
	"fmt"
)

var sli = []int{1, 43, 54, 62, 21, 66, 32, 78, 36, 76, 39}

func selectSort(sli []int) []int {
	//双重循环完成，外层控制轮数，内层控制比较次数
	len := len(sli)
	for i := 0; i < len-1; i++ {
		//先假设最小的值的位置
		k := i
		for j := i + 1; j < len; j++ {
			//sli[k] 是当前已知的最小值
			if sli[k] > sli[j] {
				//比较，发现更小的,记录下最小值的位置；并且在下次比较时采用已知的最小值进行比较。
				k = j
			}
		}
		//已经确定了当前的最小值的位置，保存到$p中。如果发现最小值的位置与当前假设的位置$i不同，则位置互换即可。
		if k != i {
			sli[k], sli[i] = sli[i], sli[k]
		}
	}
	return sli
}

func main() {
	res := selectSort(sli)
	fmt.Println(res)
}
