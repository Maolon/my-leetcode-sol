package main

import "sort"

/*
 * @lc app=leetcode id=973 lang=golang
 *
 * [973] K Closest Points to Origin
 */

// @lc code=start
func kClosest(ps [][]int, k int) [][]int {
	sort.Slice(ps, func(i, j int) bool {
		return ps[i][0]*ps[i][0]+ps[i][1]*ps[i][1] < ps[j][0]*ps[j][0]+ps[j][1]*ps[j][1]
	})

	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		ans[i] = ps[i]
	}
	return ans
}

// @lc code=end
