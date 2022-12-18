package main

/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	res := []int{}
	for _, n := range nums1 {
		m[n]++
	}

	for _, n := range nums2 {
		if m[n] > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}

// @lc code=end
