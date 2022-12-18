package main

/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input Array Is Sorted
 */

// @lc code=start
func twoSum(nums []int, tgt int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] == tgt {
			break
		}

		if nums[i]+nums[j] < tgt {
			i++
		} else {
			j--
		}
	}
	return []int{i + 1, j + 1}
}

// @lc code=end
