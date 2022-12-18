package main

/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	for _, r := range matrix {
		lo, hi := 0, len(matrix[0])-1
		for lo <= hi {
			mid := lo + (hi-lo)/2
			if r[mid] > target {
				hi = mid - 1
			} else if r[mid] < target {
				lo = mid + 1
			} else {
				return true
			}
		}
	}
	return false
}

// @lc code=end
