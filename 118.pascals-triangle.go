package main

/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

// @lc code=start
func generate(numRows int) [][]int {
	t := make([][]int, numRows)
	t[0] = []int{1}
	if numRows == 1 {
		return t
	}
	t[1] = []int{1, 1}
	if numRows == 2 {
		return t
	}
	for i := 2; i < numRows; i++ {
		len := i + 1
		t[i] = make([]int, len)
		t[i][0], t[i][len-1] = 1, 1
		for j := 1; j < len-1; j++ {
			t[i][j] = t[i-1][j-1] + t[i-1][j]
		}
	}

	return t
}

// @lc code=end
