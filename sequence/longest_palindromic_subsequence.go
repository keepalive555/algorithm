// 最长回文子序列长度
// 时间复杂度：O(n)
package main

import (
	"fmt"
)

// 状态转移方程：f(x) = arrs[i+1][j-1]
func longestPalindromicSubsequence(s string) int {
	// 二维DP数组
	n := len(s)
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n, n)
	}
	// 字节数据
	arrs := []byte(s)
	// 扫描
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if arrs[i] == arrs[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	s := "a12b34c56d7c89b0a"
	r := longestPalindromicSubsequence(s)
	fmt.Println(r)
}
