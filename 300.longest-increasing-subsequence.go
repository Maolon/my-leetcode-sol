package main

import "sort"

/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := []int{}
	// 维护一个tail数组 这里叫dp
	//
	for _, num := range nums {
		// searchints 会进行二分查找
		// 如果num不在数组里，则返回其应该插入的位置
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			// 如果tail[i-1] < num <= tail[i]
			// 则将tail[i]更新为更小的num
			dp[i] = num
		}
	}
	return len(dp)
}

// O(n^2)
// 标准dp
func lenOfLISSlow(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i]++
		res = max(dp[i], res)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
