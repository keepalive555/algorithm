// 快速排序
// 不稳定的排序算法
// 平均时间复杂度：O(nlogn)
// 最好时间复杂度：O(nlogn)
// 最差时间复杂度：O(n^2)
package main

import (
	"fmt"
)

func QuickSort(nums []int, low, high int) {
	// 递归终止
	if low >= high {
		return
	}

	i := low
	j := high
	privot := nums[i] // 基准

	for i < j {
		// 从右向左扫描
		for nums[j] >= privot && i < j {
			j--
		}
		nums[i] = nums[j]
		// 从左向右扫描
		for nums[i] <= privot && i < j {
			i++
		}
		nums[j] = nums[i]
	}
	nums[j] = privot
	// 递归QuickSort左区间
	QuickSort(nums, low, i-1)
	// 递归QuickSort右区间
	QuickSort(nums, i+1, high)

	return
}

func main() {
	nums := []int{1, 8, 9, 2, 4, 5, 3, 7, 6}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
