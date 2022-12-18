package main

/*
 * @lc app=leetcode id=1143 lang=golang
 *
 * [1143] Longest Common Subsequence
 */

// @lc code=start
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	// i = 0, j = 0, res = 0
	// 可以用空间压缩
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			//两个相等，开始记录LCS
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 不相等 三种情况，
				// 1. LCS 在 i + 1 里 就是 text2[j] == text1[i: ]的某字符
				// 2. 或者 在 j + 1 里 就是text1[i] == text2[j:] 的某字符
				// 3. 不考虑 text[i:] == text[j:] 的情况 因为LCS无论如何都短于上两种情况
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
