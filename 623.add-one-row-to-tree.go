package main

/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root, Right: nil}
	}
	addRowToTree(root, val, depth, 1)
	return root
}

func addRowToTree(n *TreeNode, val, d, lvl int) {
	//在要加上去的上一层开始插入
	if lvl == d-1 {
		n.Left = &TreeNode{Val: val, Left: n.Left, Right: nil}
		n.Right = &TreeNode{Val: val, Left: nil, Right: n.Right}
		return
	}
	if n.Left != nil {
		addRowToTree(n.Left, val, d, lvl+1)
	}
	if n.Right != nil {
		addRowToTree(n.Right, val, d, lvl+1)
	}
}

// @lc code=end
