package main

import "math"

/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	if n == 0 {
		return 0
	}
	// 等差数列
	// (1+x)*x/2 = n  x = floor(sqrt(2*n+1/4)-1/2)
	return int(math.Floor(math.Sqrt(float64(2*n)+0.25) - 0.5))
}

// @lc code=end
