package main

import "sort"

/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	dfsSubsetWithDup(nums, []int{}, &res)
	return res
}

func dfsSubsetWithDup(nums []int, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*res = append(*res, b)
	for i := 0; i < len(nums); i++ {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			dfsSubsetWithDup(nums[i+1:], append(b, nums[i]), res)
		}
	}
}

// @lc code=end
