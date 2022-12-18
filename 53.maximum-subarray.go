package main

/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
func maxSubArray(nums []int) int {
	lenP := len(nums)
	if lenP == 0 {
		return lenP
	}

	if lenP == 1 {
		return nums[0]
	}

	dp, res := make([]int, lenP), nums[0]
	dp[0] = nums[0]
	// 状态转移方程是
	// dp[i] = nums[i] + dp[i-1] (dp[i-1] > 0)，dp[i] = nums[i] (dp[i-1] ≤ 0)
	for i := 1; i < lenP; i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxSubArrayInPlace(ar []int) int {
	l := len(ar)
	max := ar[0]

	for i := 1; i < l; i++ {
		if 0 < ar[i-1] {
			ar[i] += ar[i-1]
		}

		if ar[i] > max {
			max = ar[i]
		}
	}
	return max
}

// @lc code=end
