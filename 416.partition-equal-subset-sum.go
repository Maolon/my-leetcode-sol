package main

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	// 相当于一个背包里放入值，值的总和等于sum/2(因为正好等于两半)
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%2 != 0 {
		return false
	}
	half := sum / 2
	// dp[i]为 当nums[i]时 能不能填满 大小为len(dp)的背包
	// 状态转移方程为 F(i,C) = F(i - 1,C) || F(i - 1, C - w[i])。
	// 当 i - 1 个物品就可以填满 C ，这种情况满足题意。
	// 同时如果 i - 1 个物品不能填满背包，加上第 i 个物品以后恰好可以填满这个背包
	dp := make([]bool, half+1)
	// 当背包大小为i的时候 nums[0]可以填满背包
	for i := 0; i <= half; i++ {
		dp[i] = (nums[0] == i)
	}

	for i := 1; i < len(nums); i++ {
		for j := half; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[half]
}

// @lc code=end
