package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
func eraseOverlapIntervals(intervals [][]int) int {
	// 贪心解法
	// 452题相同思路
	// 按end从小到大排序
	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	// 默认至少有一个区间
	cnt := 1
	// 默认的结束值为整个区间内end最小值（就是第一个，已排序）
	x_end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		// 找到一个最小区间外的才更新
		if intervals[i][0] >= x_end {
			cnt++
			x_end = intervals[i][1]
		}
	}
	// cnt是不重复的区间数量
	// 用总长减去cnt就是要去除的数量
	return len(intervals) - cnt
}

// @lc code=end
