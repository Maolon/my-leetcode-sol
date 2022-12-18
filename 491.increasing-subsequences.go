package main

/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Increasing Subsequences
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	dfsFindSubsequences(nums, []int{}, &res)
	return res
}

func dfsFindSubsequences(nums []int, path []int, res *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	if len(path) > 1 {
		*res = append(*res, b)
	}

	for i := 0; i < len(nums); i++ {
		if i == 0 || (i > 0 && nums[i-1] != nums[i]) {
			if len(path) == 0 || b[len(b)-1] <= nums[i] {
				dfsFindSubsequences(nums[i+1:], append(b, nums[i]), res)
			}

		}
	}
}

// @lc code=end
