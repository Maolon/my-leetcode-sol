package main

/*
 * @lc app=leetcode id=538 lang=golang
 *
 * [538] Convert BST to Greater Tree
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
func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum := 0
	convertBSTToGreater(root, &sum)
	return root
}

func convertBSTToGreater(n *TreeNode, sum *int) {
	if n == nil {
		return
	}
	convertBSTToGreater(n.Right, sum)
	n.Val += *sum
	*sum = n.Val
	convertBSTToGreater(n.Left, sum)
}

// @lc code=end
