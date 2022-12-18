package main

import "sort"

/*
 * @lc app=leetcode id=611 lang=golang
 *
 * [611] Valid Triangle Number
 */

// @lc code=start
func triangleNumber(nums []int) int {
	res := 0
	sort.Ints(nums)
	// 优化版的 三循环暴力解法
	// 从O(n^3) -> O(n^2)
	for i := 0; i < len(nums)-2; i++ {
		// k从 j+1 开始 就是 i+2
		k := i + 2
		for j := i + 1; j < len(nums)-1 && nums[i] != 0; j++ {
			for k < len(nums) && nums[i]+nums[j] > nums[k] {
				k++
			}
			res += k - j - 1
		}
	}

	return res
}

// @lc code=end
