// 二分查找
package main

import (
	"fmt"
)

func BinarySearch(nums []int, num int) int {
	// 二分查找
	left := 0
	right := len(nums) - 1

	// 临界值：空数组
	if left >= right {
		return -1
	}

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			// 搜索左区间
			right = mid - 1
		} else {
			// 搜索右区间
			left = mid + 1
		}
	}

	return -1
}

func main() {
	nums := []int{10, 13, 14, 18, 20, 22, 23, 30, 33, 45}
	idx := BinarySearch(nums, 23)
	fmt.Println(idx)
}
