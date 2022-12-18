package main

import "sort"

/*
 * @lc app=leetcode id=719 lang=golang
 *
 * [719] Find K-th Smallest Pair Distance
 */

// @lc code=start
func smallestDistancePair(nums []int, k int) int {
	// 先确保数列是递增的
	sort.Ints(nums)
	// 两数差距范围是 0 和 最后一个减去第一个
	// 可以在这个范围内搜索第k小的数
	// 二分搜索相当于指定一个差距的值 然后用另一个function判断有多少个符合这个差距的组合
	lo, hi := 0, nums[len(nums)-1]-nums[0]
	for lo < hi {
		mid := lo + (hi-lo)/2
		tmp := findDistanceCnt(nums, mid)

		if tmp >= k {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// 因为已经数组排序 可以直接进行滑动指针
func findDistanceCnt(nums []int, tgt int) int {
	cnt, i := 0, 0
	for j := 1; j < len(nums); j++ {
		// 确保i < j
		// 那么只要 nums[j] - nums[i] 大于指定差值的就不是我们想要的
		// 退出循环后用 j - i就得出所有符合条件的值的数量
		// 循环整个数组
		for nums[j]-nums[i] > tgt && i < j {
			i++
		}
		cnt += (j - i)
	}
	return cnt
}

// @lc code=end
