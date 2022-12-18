package main

/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

// @lc code=start
func hIndex(citations []int) int {
	lo, hi := 0, len(citations)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if len(citations)-mid > citations[mid] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return len(citations) - lo
}

// @lc code=end
