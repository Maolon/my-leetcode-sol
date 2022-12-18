package main

/*
 * @lc app=leetcode id=532 lang=golang
 *
 * [532] K-diff Pairs in an Array
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	cnt := 0
	m := make(map[int]int, len(nums))

	for _, n := range nums {
		m[n]++
	}

	for key := range m {
		// k = 0 是特殊情况 需要单独判断
		// 所以当map的值 > 1的时候 也就是有两个一样的数
		// cnt++
		if k == 0 && m[key] > 1 {
			cnt++
			continue
		}
		// map里存在 key + k 的值
		if m[key+k] > 0 && k > 0 {
			cnt++
		}
	}

	return cnt
}

// @lc code=end
