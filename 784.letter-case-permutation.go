package main

import (
	"strings"
)

/*
 * @lc app=leetcode id=784 lang=golang
 *
 * [784] Letter Case Permutation
 */

// @lc code=start
func letterCasePermutation(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	res := []string{}
	dfsLetterCasePermutation(strings.Split(strings.ToLower(s), ""), "", &res, len(s))
	return res
}

func dfsLetterCasePermutation(ss []string, path string, res *[]string, tot int) {
	if len(path) == tot {
		*res = append(*res, path)
	}
	for i := 0; i < len(ss); i++ {
		data := []byte(ss[i])
		if data[0] >= 'a' && data[0] <= 'z' {
			dfsLetterCasePermutation(ss[i+1:], path+strings.ToUpper(ss[i]), res, tot)
		}
		dfsLetterCasePermutation(ss[i+1:], path+ss[i], res, tot)

	}
}

// @lc code=end
