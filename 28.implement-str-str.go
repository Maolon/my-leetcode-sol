package main

import "strings"

func strStr(haystack string, needle string) int {
	return KMP(needle, haystack)
}

// KMP算法
func KMP(pat string, txt string) int {
	M := len(pat)
	N := len(txt)

	// 构建状态转移表 dp[状态][字符]
	dp := make([][]int, M)

	for i := range dp {
		dp[i] = make([]int, 256)
	}
	// 初始化第一个转移，目标为第二个状态 也就是 1
	dp[0][pat[0]] = 1
	// 初始化影子状态 （前缀一样的最长重启位置）
	x := 0
	for j := 1; j < M; j++ {
		for c := 0; c < 256; c++ {
			dp[j][c] = dp[x][c]
		}
		dp[j][pat[j]] = j + 1
		// 更新影子状态
		x = dp[x][pat[j]]
	}
	// 根据状态转移表搜索
	j := 0
	for i := 0; i < N; i++ {
		j = dp[j][txt[i]]
		// j 到达pat.length
		if j == M {
			return i - M + 1
		}
	}
	return -1
}

// 作弊解法
func strStrByGoDefault(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// Slice解法
func strStrSlice(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
