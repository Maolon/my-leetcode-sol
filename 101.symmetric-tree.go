/*
 * @lc app=leetcode id=101 lang=golang
 *
 * [101] Symmetric Tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isMirror(root.Left,root.Right)
}

func isMirror(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	// only one side nil and the other side not nil
	if l == nil || r == nil {
		return false
	}

	return l.Val == r.Val && isMirror(l.Left,r.Right) && isMirror(l.Right,r.Left)

}

// @lc code=end

