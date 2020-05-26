package main

import (
	"fmt"
)

// 最长上升子序列

func longestIncreasingSubsequence(arrs []int) int {
	// corner case
	if len(arrs) <= 1 {
		return len(arrs)
	}

	// dp数组
	n := len(arrs)
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n, n)
	}

	// 状态转移方程
}
