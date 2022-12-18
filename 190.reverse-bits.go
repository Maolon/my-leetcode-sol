package main

/*
 * @lc app=leetcode id=190 lang=golang
 *
 * [190] Reverse Bits
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var res uint32
	// 总32位bits
	for i := 0; i < 32; i++ {
		res = res<<1 | num&1
		num >>= 1
	}

	return res
}

// @lc code=end
