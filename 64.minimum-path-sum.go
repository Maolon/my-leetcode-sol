package main

import "math"

/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// 初始化col 0 上面所有i值，因为是path
	// 所以是累加上一个的值
	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}
	// 初始化row 0上的所有j值 同上
	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}

func minPathSumRecursion(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	return dp(m-1, n-1, &grid)
}

// 递归方法 超时
func dp(i, j int, grid *[][]int) int {
	if i == 0 && j == 0 {
		return (*grid)[0][0]
	}
	if i < 0 || j < 0 {
		return math.MaxInt32
	}
	return min(dp(i, j-1, grid), dp(i-1, j, grid)) + (*grid)[i][j]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
