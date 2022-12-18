package main

/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */

// @lc code=start
func getRow(rowIndex int) []int {
	r := make([]int, rowIndex+1)
	r[0] = 1

	// c^m_n = c^(m-1)_n * (n - m + 1) / n
	for i := 1; i <= rowIndex; i++ {
		r[i] = r[i-1] * (rowIndex - i + 1) / i
	}
	return r
}

// @lc code=end
