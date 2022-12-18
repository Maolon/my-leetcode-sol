package main

import "math"

/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// buy 为第i步buy的最大利润
	// sell 为第i步sell的最大利润
	buy, sell := make([]int, len(prices)), make([]int, len(prices))
	for i, _ := range buy {
		buy[i] = math.MinInt32
	}
	//假设第一步买了的价格
	buy[0] = -prices[0]
	// 第一个买 或者 第一个没买 第二个买了
	buy[1] = max(buy[0], -prices[1])
	// sell0 为 0 因为第一个不可能卖
	sell[1] = max(sell[0], buy[0]+prices[1])

	for i := 2; i < len(buy); i++ {
		// sell完的后一天是cooldown 所以是i-2，或者不buy 那么就是延续上一个i-1
		buy[i] = max(sell[i-2]-prices[i], buy[i-1])
		// buy完可以直接sell，所以是buy-1+现在的price,或者不sell，延续上一个的结果
		sell[i] = max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[len(sell)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

// @lc code=end
