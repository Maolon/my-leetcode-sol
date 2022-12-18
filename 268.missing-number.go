package main

/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	xor, i := 0, 0
	for i = 0; i < len(nums); i++ {
		// 利用 xor 相同数字为0的特性，
		// 肯定只有缺的那个不会被xor掉
		xor = xor ^ i ^ nums[i]
	}
	// 最后一位还没被xor过 xor完就是缺的那个
	return xor ^ i
}

// @lc code=end
