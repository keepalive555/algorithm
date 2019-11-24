// 简单选择排序
// 不稳定的排序算法
// 时间复杂度：O(n^2/2)
package main

import (
	"fmt"
)

// 简单选择排序
// 不是稳定的排序算法，例如：[]int{2, 2, 1, 3, 4}这种例子
func SimpleSelectionSort(arrs []int) {
	n := len(arrs)
	if n < 2 {
		return
	}

	var i, j int
	for i = 0; i < n-1; i++ {
		for j = i + 1; j < n; j++ {
			if arrs[j] < arrs[i] {
				arrs[i], arrs[j] = arrs[j], arrs[i]
			}
		}
	}
}

func main() {
	arrs := []int{1, 9, 8, 6, 7, 5, 4, 3, 2, 0}
	// arrs := []int{1}
	SimpleSelectionSort(arrs)
	fmt.Printf("arrs=%+v\n", arrs)
}
