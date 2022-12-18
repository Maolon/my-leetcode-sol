/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
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

 // bfs
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
		return [][]int{}
	}

	q := []*TreeNode{root}
	res := make([][]int,0)

	for len(q)>0 {
		l := len(q)
		tmp := make([]int,0,l)

		for i := 0; i<l ;i++{
			if q[i].Left != nil {
				q = append(q,q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q,q[i].Right)
			}
			tmp = append(tmp,q[i].Val)
		}
		q = q[l:]
		res = append(res,tmp)
	}
	
	return res
}
// @lc code=end

