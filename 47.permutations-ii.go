package main

import "sort"

/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start

// todo: 可以优化这个解法
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	dfsPermuteUnique(nums, []int{}, &res)
	return res
}

func dfsPermuteUnique(nums []int, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	if len(nums) == 0 {
		*res = append(*res, b)
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			available := []int{}
			available = append(available, nums[:i]...)
			available = append(available, nums[i+1:]...)
			dfsPermuteUnique(available, append(b, nums[i]), res)
		}

	}
}

// @lc code=end
