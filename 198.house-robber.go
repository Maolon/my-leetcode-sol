package main

/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	odd, even := 0, 0
	for idx, val := range nums {
		// 正常情况下是按奇数 偶数拿就行了
		// 但是特殊情况比如 [2,1,1,2]
		// 如果直接拿就只有3 而跳着拿 奇数拿偶数 或者偶数拿奇数
		// index = 0 和 index = 3时 最大值为 4 > 3
		// 所以考虑到这种情况 需要多比较一个跳拿的情况
		if idx%2 == 0 {
			even = max(even+val, odd)
		} else {
			odd = max(odd+val, even)
		}
	}
	return max(odd, even)
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
