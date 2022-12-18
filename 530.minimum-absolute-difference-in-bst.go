package main

import "math"

/*
 * @lc app=leetcode id=530 lang=golang
 *
 * [530] Minimum Absolute Difference in BST
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

// 与783题完全相同
func getMinimumDifference(root *TreeNode) int {
	res, nodes := math.MaxInt16, -1
	//利用中序遍历查找所有的数
	dfsBST(root, &res, &nodes)
	return res
}

func dfsBST(n *TreeNode, res, pre *int) {
	if n == nil {
		return
	}

	dfsBST(n.Left, res, pre)
	if *pre != -1 {
		*res = min(*res, abs(n.Val-*pre))
	}
	*pre = n.Val
	dfsBST(n.Right, res, pre)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
