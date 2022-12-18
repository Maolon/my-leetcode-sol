package main

/*
* @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	findPathSum(root, targetSum, []int{}, &res)
	return res
}

func findPathSum(node *TreeNode, tgt int, path []int, res *[][]int) {
	if node == nil {
		return
	}
	tgt -= node.Val
	path = append(path, node.Val)

	if tgt == 0 && node.Left == nil && node.Right == nil {
		*res = append(*res, path)
	}
	c := make([]int, len(path))
	d := make([]int, len(path))
	copy(c, path)
	copy(d, path)
	findPathSum(node.Left, tgt, c, res)
	findPathSum(node.Right, tgt, d, res)
}

// @lc code=end
