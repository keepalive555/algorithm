package main

import (
	"fmt"
)

// 归并排序，采用分治的思想
// 排序稳定性：
func MergeSort(arrs []int, low, high int, tmp []int) {
	// low >= high 时说明，数组已被分隔为单元素数组
	if low < high {
		// 切割数组
		mid := (low + high) / 2
		// 递归排序左区间
		MergeSort(arrs, low, mid, tmp)
		// 递归排序右区间
		MergeSort(arrs, mid+1, high, tmp)
		// 合并左右两个有序区间
		MergeArray(arrs, low, mid, high, tmp)
		// fmt.Printf("low=%d, high=%d, arrs=%+v\n", low, high, arrs)
	}

	return
}

// 合并数组，时间复杂度：O(n)
func MergeArray(arrs []int, low, mid, high int, tmp []int) {
	i := low     // 左区间，开始下标
	j := mid + 1 // 右区间，开始下标
	m := mid     // 左区间，结束下标
	n := high    // 右区间，结束下标
	k := 0       // 存放结果tmp的下标

	for i <= m && j <= n {
		// 临界值，相等的处理，左区间的在前
		if arrs[i] <= arrs[j] {
			tmp[k] = arrs[i]
			k++
			i++
		} else {
			tmp[k] = arrs[j]
			k++
			j++
		}
	}

	// 处理剩余区间的元素
	for i <= m {
		tmp[k] = arrs[i]
		k++
		i++
	}

	for j <= n {
		tmp[k] = arrs[j]
		k++
		j++
	}

	// 将tmp中排序好的值放入arrs中
	for i = 0; i < k; i++ {
		arrs[low+i] = tmp[i]
	}

	return
}

func main() {
	nums := []int{9, 0, 1, 2, 3, 6, 7, 5, 4, 8}
	tmp := make([]int, len(nums), 100)
	MergeSort(nums, 0, len(nums)-1, tmp)
	fmt.Printf("nums=%+v\n", nums)
}
