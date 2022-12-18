package main

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := &ListNode{Val: 0, Next: head}
	pre := newHead

	// 一直循环到 left开始
	for count := 0; pre.Next != nil && count < left-1; count++ {
		pre = pre.Next
	}

	if pre.Next == nil {
		return pre
	}
	cur := pre.Next
	// 同一位置交换 1=pre 2=cur 1->2->3->4 变成 1->3->2->4
	for i := left; i < right; i++ {
		tmp := pre.Next
		pre.Next = cur.Next
		cur.Next = cur.Next.Next
		pre.Next.Next = tmp
	}
	return newHead.Next
}

// @lc code=end
