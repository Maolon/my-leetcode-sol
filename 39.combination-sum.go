package main

import "sort"

/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	len := len(candidates)
	if len == 0 {
		return [][]int{}
	}

	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinations(candidates, target, 0, c, &res)
	return res
}

func findCombinations(nums []int, tgt int, index int, c []int, res *[][]int) {
	if tgt <= 0 {
		if tgt == 0 {
			newC := make([]int, len(c))
			copy(newC, c)
			*res = append(*res, newC)
		}
		return
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > tgt {
			break
		}
		c = append(c, nums[i])
		findCombinations(nums, tgt-nums[i], i, c, res)
		c = c[:len(c)-1]
	}

}

// @lc code=end
