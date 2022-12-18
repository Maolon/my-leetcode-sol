package main

/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	res1, res2 := dfsRob(root)
	return max(res1, res2)
}

func dfsRob(n *TreeNode) (a, b int) {
	if n == nil {
		return 0, 0
	}
	nl, rl := dfsRob(n.Left)
	nr, rr := dfsRob(n.Right)
	// 当前不打劫
	nop := max(nl, rl) + max(nr, rr)
	// 当前打劫
	robb := n.Val + nr + nl

	return nop, robb
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
