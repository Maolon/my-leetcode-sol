package main

import "sort"

/*
 * @lc app=leetcode id=452 lang=golang
 *
 * [452] Minimum Number of Arrows to Burst Balloons
 */

// @lc code=start

func findMinArrowShots(points [][]int) int {
	// 贪心算法
	// 435相同思路，见435题
	sort.Slice(points, func(x, y int) bool {
		return points[x][1] < points[y][1]
	})

	cnt := 1
	x_end := points[0][1]

	for i := 0; i < len(points); i++ {
		if points[i][0] > x_end {
			cnt++
			x_end = points[i][1]
		}
	}
	return cnt
}

// @lc code=end
