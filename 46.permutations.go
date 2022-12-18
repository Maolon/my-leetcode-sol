package main

/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := [][]int{}
	dfsPermute(nums, []int{}, &res)
	return res
}

func dfsPermute(nums []int, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	if len(nums) == 0 {
		*res = append(*res, b)
	}
	for i := 0; i < len(nums); i++ {
		available := []int{}
		// 两步 available 的目的是从nums备选里去除当前的i(i只能用一次)
		available = append(available, nums[:i]...)
		available = append(available, nums[i+1:]...)
		dfsPermute(available, append(b, nums[i]), res)
	}
}

// @lc code=end
