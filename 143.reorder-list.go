package main

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	p1 := head
	p2 := head
	// p2 双步走 查找中点
	for p2.Next != nil && p2.Next.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
	}

	mid := p1
	midCur := mid.Next
	for midCur.Next != nil {
		tmp := midCur
		mid.Next = midCur.Next
		midCur.Next = midCur.Next.Next
		mid.Next.Next = tmp
	}

	p1 := head
	p2 := mid.Next

}

// @lc code=end
