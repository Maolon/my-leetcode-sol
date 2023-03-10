package main

/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	m := map[int]bool{}
	res := []int{}
	for _, n := range nums1 {
		m[n] = true
	}

	for _, n := range nums2 {
		if m[n] {
			delete(m, n)
			res = append(res, n)
		}
	}
	return res
}

// @lc code=end
