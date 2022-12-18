// 空间压缩版本， 因为每一位其实只取dp[i-1]的值且无重叠
// 所以可以优化一维数组
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	// 初始状态当j = 0的时候只有一种解法
	// 所以为1
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i]]
			}
		}
	}
	return dp[amount]
}