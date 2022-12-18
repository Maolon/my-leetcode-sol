package main

import "strings"

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	ss := strings.Fields(s)
	i, j := 0, len(ss)-1
	for i <= j {
		ss[i], ss[j] = ss[j], ss[i]
		i++
		j--
	}
	return strings.Join(ss, " ")
}

// @lc code=end
