// K路归并排序
package main

import (
	"fmt"
)

const (
	MinInt int = -(1<<31 - 1)
	MaxInt int = 1 << 31
)

type LoserTree struct {
	k  int   // K路归并排序
	ls []int // 内部节点（非叶子节点），节点个数：K-1
	b  []int // 叶子节点，加站桩节点，总个数：K+1
}

func NewLoserTree(k int) *LoserTree {
	loserTree := &LoserTree{
		k:  k,
		ls: make([]int, k, k),
		b:  make([]int, k+1, k+1),
	}
	// 初始化败者树
	for i := 0; i < k; i++ {
		loserTree.ls[i] = k
	}
	// 占位
	loserTree.b[k] = MinInt
	return loserTree
}

func (ls *LoserTree) Set(i int, value int) {
	if i >= ls.k {
		return
	}
	// 要比之前节点大
	if value > ls.b[i] {
		ls.b[i] = value
	}
}

func (ls *LoserTree) Get(i int) (int, bool) {
	if i >= ls.k {
		return 0, false
	}
	return ls.b[i], true
}

func (ls *LoserTree) Top() int {
	return ls.ls[0]
}

func (ls *LoserTree) Adjust(i int) {
	p := (i + ls.k) / 2
	for p > 0 {
		if ls.b[i] > ls.b[ls.ls[p]] {
			ls.ls[p], i = i, ls.ls[p]
		}
		p /= 2
	}
	ls.ls[0] = i
}

func KMergeSort(arrList [][]int) []int {
	// 临界条件判断
	k := len(arrList)
	if k < 0 {
		return nil
	}
	if k == 1 {
		return arrList[0]
	}
	// K路归并排序
	posList := make([]int, k, k)
	ls := NewLoserTree(k)
	for i := 0; i < k; i++ {
		ls.Set(i, arrList[i][posList[i]])
	}
	// 调整败者树
	for i := 0; i < k; i++ {
		ls.Adjust(i)
	}

	// 记录各数组完成进度
	finishList := make([]bool, k, k)
	isFinish := func() bool {
		for _, v := range finishList {
			if v == false {
				return false
			}
		}
		return true
	}

	size := 0
	for _, arr := range arrList {
		size += len(arr)
	}
	res := make([]int, 0, size)

	for !isFinish() {
		idx := ls.Top()
		posList[idx]++
		v, _ := ls.Get(idx)
		res = append(res, v)

		if posList[idx] < len(arrList[idx]) {
			ls.Set(idx, arrList[idx][posList[idx]])
		} else {
			ls.Set(idx, MaxInt)
			finishList[idx] = true
		}
		ls.Adjust(idx)
	}
	return res
}

func main() {
	arrsList := [][]int{
		[]int{1, 3, 5, 7, 9},
		[]int{2, 4, 6, 8, 10},
		[]int{1, 1, 2, 3},
		[]int{6, 11, 24, 50},
	}
	res := KMergeSort(arrsList)
	fmt.Println(res)
}
