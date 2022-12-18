package main

// type TreeNode struct {
// 	    Val int
// 	     Left *TreeNode
// 	     Right *TreeNode
//  }

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
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
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{} // generate a null tree
	}
	return buildTree(1, n)
}

func buildTree(start, end int) []*TreeNode {
	tree := []*TreeNode{}
	// 结束条件
	if start > end {
		append(tree, nil)
		return tree
	}
	for i := start; i <= end; i++ {
		left := buildTree(start, i-1)
		right := buildTree(i+1, end)
		for _, l := range left {
			for _, r := range right {
				tree = append(tree, &TreeNode{i, l, r})
			}
		}
	}
	return tree
}

// @lc code=end
