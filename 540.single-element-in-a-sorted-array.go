package main

/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start
func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-lo)/2
		// 如果没有遇到缺的那个数字 那每个数字应该都是2位
		// 也就是说 当idx为偶数时 nums[idx] == nums[idx+1]
		// 当出现 nums[idx] !== nums[idx+1]时
		// 说明在这个idx之前已经出现了那个单独的数 向左查找 否 向右查找
		if mid%2 == 0 {
			if nums[mid] == nums[mid+1] {
				lo = mid + 1
			} else {
				hi = mid
			}
		} else {
			// 奇数同理
			if nums[mid] == nums[mid-1] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}

	}
	return nums[lo]
}

// @lc code=end
