package main

/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
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

func inorderTraversal(root *TreeNode) []int {
	var res []int
	traversal(root, &res)
	return res
}

// 在go中 array的传递是值传递
// 也就是说如果 output 为 output []int
// 他会从inordertraversal拿到res的copy
// 而我们需要output的引用
func traversal(r *TreeNode, output *[]int) {
	if r == nil {
		return
	}
	traversal(r.Left, output)
	// 改output指针的值
	// append拿到的是指针的值
	*output = append(*output, r.Val)
	traversal(r.Right, output)
}

// @lc code=end
