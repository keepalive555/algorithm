// 冒泡排序
// 稳定排序
package main

import (
	"fmt"
)

// 比较相邻的两个元素
func BubbleSort(arrs []int) {
	n := len(arrs)
	if n < 2 {
		return
	}

	var i, j int
	for i = 0; i < n; i++ {
		// 临界值处理n-i-1最后一个元素不用进行比较
		for j = 0; j < n-i-1; j++ {
			if arrs[j] > arrs[j+1] {
				// 向后交换大元素
				arrs[j], arrs[j+1] = arrs[j+1], arrs[j]
			}
		}
	}
}

func main() {
	nums := []int{9, 8, 2, 1, 5, 4, 3, 7, 6, 0}
	BubbleSort(nums)
	fmt.Println(nums)
}
