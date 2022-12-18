/*
 * @lc app=leetcode id=99 lang=golang
 *
 * [99] Recover Binary Search Tree
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
func recoverTree(root *TreeNode) {
	var prev, tgt1, tgt2 *TreeNode
	_, tgt1, tgt2 = inOrderTraverse(root, prev, tgt1, tgt2)
	if tgt1 != nil && tgt2 != nil {
		tgt1.Val, tgt2.Val = tgt2.Val, tgt1.Val
	}
}

func inOrderTraverse(root, prev, tgt1, tgt2 *TreeNode) (*TreeNode, *TreeNode, *TreeNode) {
	if root == nil {
		return prev, tgt1, tgt2
	}

	prev, tgt1, tgt2 = inOrderTraverse(root.Left, prev, tgt1, tgt2)
	// prev 是左子树返回的内部节点
	if prev != nil && prev.Val > root.Val {
		if tgt1 == nil {
			tgt1 = prev
		}
		tgt2 = root
	}
	prev = root
	prev, tgt1, tgt2 = inOrderTraverse(root.Right, prev, tgt1, tgt2)
	return prev, tgt1, tgt2
}

// @lc code=end

