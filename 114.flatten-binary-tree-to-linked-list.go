package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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
func flatten(root *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return
	}
	flatten(root.Right)
	flatten(root.Left)
	// 存一下Right
	curRight := root.Right
	// 把左tree挪到现在的右tree的位置
	root.Right = root.Left
	// 清空左tree
	root.Left = nil
	// 一直查找到右tree的底
	for root.Right != nil {
		root = root.Right
	}
	// 把现在的Right挪到最后
	root.Right = curRight
}

// @lc code=end
