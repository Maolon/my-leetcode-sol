/*
 * @lc app=leetcode id=96 lang=golang
 *
 * [96] Unique Binary Search Trees
 */

// @lc code=start
func numTrees(n int) int {
	dp := make([]int,n+1)
	dp[0] , dp[1] = 1,1
    
	// dp[n] 代表 1-n 个数能组成多少个不同的二叉排序树
	// F(i,n) 代表以 i 为根节点，1-n 个数组成的二叉排序树的不同的个数
	// F(i,n) = dp[i-1] * dp[n-i] <- 相当于 从 F（i-1,n-i）(i-1)（左节点） 和 F(i+1,n-i) (n-i)（右节点） 两个树不同个数相乘的总和
	// ⬆️ 因为是二叉树，所以我们不需要考虑顺序问题了 直接取 i-1 和 n-i （就是直接取个数）两个树的不同个数相乘
	// dp[n] = F(1,n) + F(2,n) + F(3,n) + …… + F(n,n)
  
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// @lc code=end

