package main

/*
 * @lc app=leetcode id=516 lang=golang
 *
 * [516] Longest Palindromic Subsequence
 */

// @lc code=start
func longestPalindromeSubseq(s string) int {
	for len(s) == 0 {
		return 0
	}
	dp := make([][]int, len(s)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
		// 每个字符都是自己的回文串
		dp[i][i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				// 两个相等 把两个字符同时加入回文串
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][len(s)-1]
}

func lPSWithDpSpaceCompress(s string) int {
	// 我们可以注意到 正常解法使用了二维数组
	// 但实际使用到的仅仅是 dp[i][j] dp[i+1][j-1] dp[i+1][j] dp[i][j-1]这四个状态
	//                  j
	// 且读取顺序是 i [3,目标]
	//              [1, 2]
	// 也就是dp[i+1][j-1]的值会被 3号目标 dp[i][j-1]覆盖，所以我们需要用零时值记住
	for len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s)+1)
	for i := range dp {
		dp[i] = 1
	}
	for i := len(s) - 2; i >= 0; i-- {
		// 会被覆盖的 dp[i+1][j-1]的值
		pre := 0
		for j := i + 1; j < len(s); j++ {
			// 存好下一个pre 也就是会被覆盖的 dp[i+1][j-1]的值
			temp := dp[j]
			if s[i] == s[j] {
				// dp[i][j] = dp[i+1][j-1] + 2
				dp[j] = pre + 2
			} else {
				// dp[i][j] = max(dp[i+1][j], dp[i][j-1])
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = temp
		}
	}
	return dp[len(s)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
