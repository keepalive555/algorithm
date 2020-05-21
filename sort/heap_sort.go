// 堆排序
// 不稳定的排序算法
// 参考文档：https://www.cnblogs.com/chengxiao/p/6129630.html
// 时间复杂度：O(n*logn)
package main

import (
	"fmt"
)

// 节点左孩子节点
func left(i int) int {
	return 2*i + 1
}

// 节点右孩子节点
func right(i int) int {
	return 2*i + 2
}

// 节点父节点
func parent(i int) int {
	return (i - 1) / 2
}

// 将以节点i的为根的子树，调整为大顶堆
func HeapAdjust(arrs []int, length int, i int) {
	value := arrs[i]
	child := left(i) // 孩子节点下标
	for child < length {
		// 存在右节点，且右节点的值大于左节点，则child指向右节点
		if child+1 < length && arrs[child] < arrs[child+1] {
			child++
		}
		// 如果父节点值大于等于两个孩子节点，则退出
		if value >= arrs[child] {
			break
		}
		// 交换父节点与子节点值
		arrs[i] = arrs[child]
		i = child
		child = left(i)
	}

	// 将根节点放入最终位置
	arrs[i] = value
}

func HeapSort(arrs []int) {
	// 将数组调整为大顶堆
	n := len(arrs)
	var i, j int
	for i = n/2 - 1; i >= 0; i-- {
		HeapAdjust(arrs, n, i)
	}
	for j = n - 1; j >= 0; j-- {
		arrs[0], arrs[j] = arrs[j], arrs[0]
		HeapAdjust(arrs, j, 0)
	}
}

func main() {
	nums := []int{1, 9, 8, 7, 2, 3, 5, 6, 4, 0}
	HeapSort(nums)
	fmt.Printf("nums=%+v\n", nums)
}
