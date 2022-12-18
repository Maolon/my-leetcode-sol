package main

/*
 * @lc app=leetcode id=115 lang=golang
 *
 * [115] Distinct Subsequences
 */

// @lc code=start
func numDistinct(s string, t string) int {
	// i代表 i -> len(s) 内为被匹配文字序列
	// j代表 j -> len(j) 内为被匹配文字序列
	// dp[i][j] 代表 s[i:] 即(i-len(s)) 和 t[j:] 即(j-len(s)) 的匹配个数
	// 当 i == len(s) 时，dp[i][j] = 0， 因为此时被匹配序列为空，所以匹配个数为0
	// 当 j == len(t) 时，dp[i][j] = 1， 因为此时匹配序列为空， 空序列可以匹配任何序列，所以匹配个数为1
	// 转移方程：
	// 当 s[i] == t[j] 时，dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
	// ⬆️ 两种情况
	// 1. s[i] 直接匹配 t[j]，那么 子序列 s[i+1:] 和 子序列t[j+1:] 的匹配个数
	// 2. s[i] 匹配且子序列也匹配 t[j]，那么 子序列 s[i+1:] 和 序列t[j:] 的匹配个数
	// 当 s[i] != t[j] 时，dp[i][j] = dp[i+1][j]
	// ⬆️ s[i] 不匹配 t[j]，那么 子序列 s[i+1:] 和 子序列t[j:] 的匹配个数

	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([][]int, m+1)
	//初始化二维数组
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][n] = 1
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == t[j] {
				dp[i][j] = dp[i+1][j+1] + dp[i+1][j]
			} else {
				dp[i][j] = dp[i+1][j]
			}
		}
	}
	return dp[0][0]
}

// @lc code=end
