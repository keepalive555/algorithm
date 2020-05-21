// 最长回文子序列长度
// 时间复杂度：O(n)
package main

// 状态转移方程：f(x) = arrs[i+1][j-1]

func longestPalindromicSubsequence(s string) int {
	// 二维DP数组
	n := len(s)
	dp := make([][]int, n, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n, n)
	}
	// 扫描
	var i, j int
	for j = n - 1; j >= 0; j-- {
	}
	return ""
}

func main() {
}
