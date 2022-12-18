package main

/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	// dp方法
	// sum(p) - sum(n) = target
	// sum(p) + sum(n) + sum(p) - sum(n) = target + sum(p) + sum(n)
	// 2 * sum(p) = target + sum(all nums)
	// 所以寻找目标是  result = (target + sum(all nums)) / 2
	tot := 0
	for _, n := range nums {
		tot += n
	}

	if target > tot || (target+tot)%2 != 0 {
		return 0
	}
	// dp[i]代表 当 target = i时 可能的组合数
	dp := make([]int, target+1)
	dp[0] = 1
	tgt := (target + tot) / 2
	for _, n := range nums {
		for i := tgt; i >= n; i-- {
			//
			dp[i] += dp[i-n]
		}
	}
	return dp[tgt]
}

// @lc code=end
