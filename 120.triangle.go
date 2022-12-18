package main

/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(t [][]int) int {
	if t == nil {
		return 0
	}
	for r := len(t) - 2; r >= 0; r-- {
		for c := 0; c < len(t[r]); c++ {
			t[r][c] += min(t[r+1][c+1], t[r+1][c])
		}
	}
	return t[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
