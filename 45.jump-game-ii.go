package main

/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	lenN := len(nums)
	if lenN == 1 {
		return 0
	}
	max := 0
	steps := 0
	nextStep := 0
	for i, n := range nums {
		if n+i > max {
			max = n + i
			if max >= lenN-1 {
				return steps + 1
			}
		}

		if i == nextStep {
			nextStep = max
			steps++
		}
	}
	return steps
}

// @lc code=end
