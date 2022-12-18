package main

/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
	p := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			p += prices[i+1] - prices[i]
		}
	}
	return p
}

// @lc code=end
