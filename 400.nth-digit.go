package main

import (
	"math"
)

/*
 * @lc app=leetcode id=400 lang=golang
 *
 * [400] Nth Digit
 */

// @lc code=start
func findNthDigit(n int) int {
	// [1-9]， 每个数有一位数字， 这样的数有9个 bits = 1 ,9 = math.pow10(bits-1) * bits
	// [10-99], 每个数有两位数字， 这样的数有99-9=90个
	// [100-999], 每个数有三位数字，这样的数有999-99=900个
	// n 不断减去 bits 从 1 开始的数字总数,求出 n 所在的数字是几位数即 bits
	// 计算 n 所在的数字 num，等于初始值加上 (n - 1) / bits
	// 计算 n 所在这个数字的第几位 digitIdx 等于 (n - 1) % bits
	if n <= 9 {
		return n
	}
	//因为 n > 9 所以默认bits已经有1了
	bits := 1
	// 因为默认bits已经为1， 为1的时候总共有9个数
	for n > 9*int(math.Pow10(bits-1))*bits {
		// 为了查找到底在哪个数字段里，减去之前已经确定不在的范围
		n -= 9 * int(math.Pow10(bits-1)) * bits
		bits++
	}
	// 默认是当前数组段的第一位
	idx := n - 1
	start := int(math.Pow10(bits - 1))
	num := start + idx/bits
	digitIdx := idx % bits
	// 求解具体的位数所在的数
	// int(math.Pow10(bits-digitIdx-1)) 算出具体factor 然后用num / 这个数把范围限制在 1 - 10 里
	// 因为是10进制 所以最后还要%10来得到正确的位数， 即当答案为10的时候实际答案应该是0，这样范围就是个位数 0 - 9 了
	return num / int(math.Pow10(bits-digitIdx-1)) % 10
}

// @lc code=end
