package main

import "math"

/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	num3, stack := math.MinInt32, []int{}
	// 假设序列为 [1,4,2]
	for i := len(nums) - 1; i >= 0; i-- {
		// 这里确保的是 1 < 4的部分
		if nums[i] < num3 {
			return true
		}
		// 推出所有比当前少的值 因为我们要的是 [1,4,2] 这样的序列
		// 这里准备的是 4 > 2的部分，确保找到非递减的值后就默认这个值为nums3
		for len(stack) != 0 && nums[i] > stack[len(stack)-1] {
			//取最后一位
			num3 = stack[len(stack)-1]
			//pop最后一位
			stack = stack[:len(stack)-1]
		}
		//推入栈
		stack = append(stack, nums[i])

	}

	return false
}

// @lc code=end
