/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(n int) []int {
	b := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 0: 0000 = 0
		// 1: 0001 = 1
		// 2: 0010 = 1
		// 3: 0011 = 2
		// 4: 0100 = 1
		// 5: 0101 = 2
		// 6: 0110 = 2
		// 7: 0111 = 3
		// 8: 1000 = 1
		// 利用的是清除最高位一下的1之后+1
		// 比如在有进位的情况下 i = 8, 8 & 7 = 0, 取b[0], 加1 = 1
		// 比如非进位条件下 i = 7, 7 & 6 = 0111 & 0110 = 0110, 取b[6] = 2, 加1 = 3
		b[i] += b[i&(i-1)] + 1
	}
	return b
}

// @lc code=end

