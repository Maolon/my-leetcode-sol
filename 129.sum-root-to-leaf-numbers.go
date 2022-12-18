package main

/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	res := 0
	findSumNumbers(root, 0, &res)
	return res
}

func findSumNumbers(n *TreeNode, sum int, res *int) {
	if n == nil {
		return
	}

	sum = sum*10 + n.Val

	if n.Left == nil && n.Right == nil {
		*res += sum
	}

	findSumNumbers(n.Left, sum, res)
	findSumNumbers(n.Right, sum, res)
}

// @lc code=end
