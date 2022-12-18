package main

/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	end := len(nums) - 1

	cur, lst := 0, 0

	for lst < end {
		for nums[lst] == nums[cur] {
			cur++
			if cur == end+1 {
				return lst + 1
			}
		}
		nums[lst+1] = nums[cur]
		lst++
	}
	return lst + 1
}

// @lc code=end
