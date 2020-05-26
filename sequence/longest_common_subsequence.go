package main

import (
	"fmt"
)

// 最长公共子序列
// 解题思路：假设字符x，在子序列a中出现过，子序列b中出现过，则字符x一定在最长公共子序列中

func longestCommonSubsequence(a, b string) int {
	// corner case
	if a == "" || b == "" {
		return 0
	}
	arrs1 := []byte(a)
	arrs2 := []byte(b)
	n1 := len(a)
	n2 := len(b)

	// dp数组
	dp := make([][]int, n1+1, n1+1)
	for i := 0; i <= n1; i++ {
		dp[i] = make([]int, n2+1, n2+1)
	}
	// 动态规划
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			if arrs1[i-1] == arrs2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[n1][n2]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	a := "m1y3w5or8k"
	b := "m2y4wo6r779d"
	r := longestCommonSubsequence(a, b)
	fmt.Println(r)
}
