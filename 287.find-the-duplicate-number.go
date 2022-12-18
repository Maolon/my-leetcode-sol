package main

/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start
// 快慢指针
// 把整个数组当成linkedlist
// 数组的值是next的index
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]

	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	findPtr := 0

	for findPtr != slow {
		findPtr = nums[findPtr]
		slow = nums[slow]
	}
	return findPtr
}

// b search
// 因为数组范围在1 - n之间 数组长度为n+1
// 所以判断一个区间比如 1 - mid 内如果没有重复 小于或等于mid的数量的数字数量最多为mid
// count 小于或等于mid的数量， 如果大于mid则重复数存在于[1-mid]这个区间 否则存在于（mid-n]这个区间内
func findDuplicateBSearch(nums []int) int {
	if len(nums) < 2 {
		return -1
	}
	l, h := 1, len(nums)
	for l < h {
		count := 0
		mid := (l + h) / 2
		for _, n := range nums {
			if n <= mid {
				count++
			}
		}
		if count > mid {
			h = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// @lc code=end
