package main

/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start

// 因为只有简单的0，1，2三个数字
// 也就是说 只要找到一个2 前面就会有1，0
// 找到一个 1 前面就会有 0
// 建立 z (zero) o (one) 两个前置pointer
// 然后iterate数组 按上面的思路填充数字
func sortColors(nums []int) {
	z, o := 0, 0
	for i, n := range nums {
		// 默认当前为2 如果不是会在后面的循环内被覆盖
		nums[i] = 2
		if n <= 1 {
			nums[o] = 1
			o++
		}
		if n == 0 {
			nums[z] = 0
			z++
		}
	}

}

// @lc code=end
