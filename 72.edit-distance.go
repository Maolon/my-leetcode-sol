package main

/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	// 遍历法
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// base case 相当于word1 结束 而 word2没结束时剩余的次数
	for i := 0; i < m; i++ {
		dp[i][n] = m - i
	}
	// 同上 同理
	for j := 0; j < n; j++ {
		dp[m][j] = n - j
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word1[i] == word2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				// 删除  相当于向s1内删除当前值 s2不动
				dp[i][j] = 1 + dp[i+1][j]
				// 插入 相当于向s1内插入s2的j上得值
				dp[i][j] = min(dp[i][j], 1+dp[i][j+1])
				// 替换 相当于s1[i]替换成s2[j]
				dp[i][j] = min(dp[i][j], 1+dp[i+1][j+1])
			}
		}
	}
	return dp[0][0]
}

// recursive 方法 但是 time limit exceed
func dp(s1 string, i int, s2 string, j int) int {
	if i == -1 {
		// 当 i走完的时候 如果j还没走完
		// 相当于插入所有j剩下的值
		return j + 1
	}
	if j == -1 {
		// 同上 同理
		return i + 1
	}

	if s1[i] == s2[j] {
		return dp(s1, i-1, s2, j-1)
	}
	//三种子情况
	// 1. 插入 dp(s1,i,s2,j-1) + 1 相当于向s1内插入s2的j上得值
	// 2. 删除 dp(s1,i-1,s2,j)+ 1 相当于向s1内删除当前值 s2不动
	// 3. 替换 dp(s1,i-1,s2,j-1) + 1 相当于s1[i]替换成s2[j]
	return min(dp(s1, i, s2, j-1)+1,
		min(dp(s1, i-1, s2, j)+1,
			dp(s1, i-1, s2, j-1)+1))
}

// recursive 优化了时间 但是 time limit exceed
func dpOptimized(s1 string, i int, s2 string, j int, memo *[][]int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}

	if s1[i] == s2[j] {
		(*memo)[i][j] = dp(s1, i-1, s2, j-1)
	} else {
		(*memo)[i][j] = min(dp(s1, i, s2, j-1)+1,
			min(dp(s1, i-1, s2, j)+1,
				dp(s1, i-1, s2, j-1)+1))
	}
	return (*memo)[i][j]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
