package main

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			// dp[i] 指的是 n = i 时最大的乘积
			// 三种情况,取最大值
			// 1. 不在这一步做任何操作 dp[i] = dp[i]
			// 2. 在这一步做数 j * 拆分数i - j，只做两步拆分 （j strictly < i）
			// 3. 这一步的数j * 之前数 i - j 但是这个i-j还能继续拆分 所以是 dp[i-j]
			dp[i] = max(dp[i], j*max(i-j, dp[i-j]))
		}

	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
