package main

import "math"

/*
 * @lc app=leetcode id=124 lang=golang
 *
 * [124] Binary Tree Maximum Path Sum
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
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}

func getPathSum(n *TreeNode, maxSum *int) int {
	if n == nil {
		return math.MinInt32
	}
	l := getPathSum(n.Left, maxSum)
	r := getPathSum(n.Right, maxSum)
	// 总共比较四种状态谁最大
	// 左branch + 当前val， 右branch + 当前val， 左+右+当前val， 当前记录的总最大值
	curMax := max(max(l+n.Val, r+n.Val), n.Val)
	*maxSum = max(*maxSum, max(curMax, l+r+n.Val))
	return curMax
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// @lc code=end
