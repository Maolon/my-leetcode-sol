package main

/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	return []int{searchFstEqlEle(nums, target), searchLstEqlEle(nums, target)}
}

// 寻找第一个跟target相等的数
func searchFstEqlEle(nums []int, tgt int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if tgt > nums[mid] {
			lo = mid + 1
		} else if tgt < nums[mid] {
			hi = mid - 1
		} else {
			if mid == 0 || (nums[mid-1] != tgt) {
				return mid
			}
			// 如果不是第一个 则第一个还在左边
			// 当成 tgt < nums[mid] 判断
			hi = mid - 1
		}
	}
	return -1
}

// 寻找最后一个跟target相等的数
func searchLstEqlEle(nums []int, tgt int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if tgt > nums[mid] {
			lo = mid + 1
		} else if tgt < nums[mid] {
			hi = mid - 1
		} else {
			if mid == len(nums)-1 || (nums[mid+1] != tgt) {
				return mid
			}
			// 如果不是最后一个 则最后一个还在右边
			// 当成 tgt > nums[mid] 判断
			lo = mid + 1
		}
	}
	return -1
}

// @lc code=end
