package main

/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
func buildTree(inorder []int, postorder []int) *TreeNode {
	postLen := len(postorder)
	if postLen == 0 {
		return nil
	}
	root := postorder[postLen-1]
	tree := &TreeNode{Val: root}
	postorder = postorder[:postLen-1]
	for pos, node := range inorder {
		if node == root {
			//  pos = 1 inorder[9] len(inorder[:pos] = 3 postorder[15,7,20]
			tree.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			// pos = 1 inorder[15,20,7] len(inorder[:pos] = 1 postorder[15,7,20]
			tree.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return tree
}

// @lc code=end
