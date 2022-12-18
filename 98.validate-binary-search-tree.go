/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	//  minValue of int is -1*2e63, maxValue of int is 1*2e64 - 1
	//  also from math library
	// 	MaxInt    = 1<<(intSize-1) - 1
	// 	MinInt    = -1 << (intSize - 1)
    return isValid(root, -1<<63, 1<<63-1)
}

func isValid(tree *TreeNode, min, max int) bool {
	if tree == nil {
		return true
	}
	if tree.Val <= min || tree.Val >= max {
		return false
	}
	return isValid(tree.Left, min, tree.Val) && isValid(tree.Right, tree.Val, max)
}
// @lc code=end

