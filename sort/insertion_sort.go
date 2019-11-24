package main

import (
	"fmt"
)

// 插入排序
// 类比『摸扑克牌』
// 插入排序对几乎已经排序好的数据操作时，效率高，可以达到线性排序的效率
// 插入排序在倒序时性能最差
func InsertionSort(arrs []int) {
	n := len(arrs)
	if n < 2 {
		return
	}
	var i, j int
	for i = 1; i < n; i++ {
		val := arrs[i]
		for j = i - 1; j >= 0 && arrs[j] > val; j-- {
			arrs[j+1] = arrs[j]
		}
		arrs[j+1] = val
	}
	return
}

func main() {
	arrs := []int{9, 8, 2, 1, 4, 5, 0, 6, 3}
	InsertionSort(arrs)
	fmt.Println(arrs)
}
