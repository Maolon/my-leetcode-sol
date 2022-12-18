package main

import "math"

/*
 * @lc app=leetcode id=174 lang=golang
 *
 * [174] Dungeon Game
 */

// @lc code=start
func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	n := len(dungeon[0])
	// dp代表dp[i][j]时 到达右下角所需的最小血量
	// 当前的位置到达右下角所需的最小血量可以由 min(dp[i+1][j],dp[i][j+1]) - dp[i][j]得出
	// 当然他可能小于0, 所以此时dp[i][j] = 1 即保证骑士不死的最小血量
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//开始填充初始值 (实际应该从m-1,n-1开始)
	if dungeon[m-1][n-1] > 0 {
		// 如果最后一个格子时补血，此时补不补都无所谓
		// 所以只要保证来到此格的时候血量为1，就是最小就行了
		dp[m-1][n-1] = 1
	} else {
		// 如果最后一格时扣血
		// 那么就要比扣的血 > 1
		dp[m-1][n-1] = -dungeon[m-1][n-1] + 1
	}
	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			// 用最大值填充没用到的两边
			if i == m || j == n {
				dp[i][j] = math.MaxInt
				continue
			}
			// 跳过初始值
			if i == m-1 && j == n-1 {
				continue
			}
			res := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			if res <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = res
			}
		}
	}
	return dp[0][0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
