package main

/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	// 由于首尾是相邻的，所以需要对比 [0，n-1]、[1，n] 这两个区间的最大值
	return max(robInRange(nums, 0, n-2), robInRange(nums, 1, n-1))
}

func robInRange(nums []int, start, end int) int {
	preMax := nums[start]
	curMax := max(nums[start+1], preMax)

	for i := start + 2; i <= end; i++ {
		tmp := curMax
		curMax = max(curMax, nums[i]+preMax)
		preMax = tmp
	}
	return curMax
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
