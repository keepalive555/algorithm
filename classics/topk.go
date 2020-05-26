package main

import (
	"fmt"
	"sort"
)

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func parent(i int) int {
	return (i - 1) / 2
}

// 大顶堆
func HeapAdjust(arrs []int, i, length int) {
	value := arrs[i]
	child := left(i)
	for child < length {
		// 存在右孩子，并且右孩子值比左孩子大
		if child+1 < length && arrs[child+1] > arrs[child] {
			child++
		}
		// 若父节点大于等于子节点中较大值
		if value >= arrs[child] {
			break
		}
		// 孩子节点填补至父节点位置
		arrs[i] = arrs[child]
		i = child
		child = left(i)
	}
	arrs[i] = value
}

func TopK(arrs []int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	n := len(arrs)
	if k >= n {
		return arrs
	}
	// 使用堆
	// k/2-1，非叶子节点，叶子节点无子节点，无需HeapAdjust
	for j := k/2 - 1; j >= 0; j-- {
		HeapAdjust(arrs, j, k)
	}
	// 求topk
	for i := k; i < len(arrs); i++ {
		if arrs[i] < arrs[0] {
			arrs[i], arrs[0] = arrs[0], arrs[i]
			HeapAdjust(arrs, 0, k)
		}
	}
	return arrs[:k]
}

func main() {
	arrs := []int{999, 1231, 123, 12, 2, 3, 488, 2929, 3234, 111, 33, 44, 55}
	res := TopK(arrs, 5)
	fmt.Printf("top k result = %+v\n", res)
	sort.Ints(arrs)
	fmt.Printf("sorted array = %+v\n", res)
}
