package main

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	min, max, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			min, max = max, min
		}
		max = maxInt(nums[i], max*nums[i])
		min = minInt(nums[i], min*nums[i])
		res = maxInt(res, max)
	}
	return res
}

func maxInt(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

func minInt(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end
