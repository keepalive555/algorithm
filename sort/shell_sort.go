package main

import (
	"fmt"
)

// 希尔排序
// 希尔排序是插入排序的升级版本
func ShellSort(arrs []int) {
	var i, j int
	n := len(arrs)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i = gap; i < n; i++ {
			val := arrs[i]
			for j = i - gap; j >= 0 && arrs[j] > val; j -= gap {
				arrs[j+gap] = arrs[j]
			}
			arrs[j+gap] = val
		}
	}
}

func main() {
	arrs := []int{9, 8, 2, 1, 4, 5, 0, 6, 3}
	ShellSort(arrs)
	fmt.Println(arrs)
}
