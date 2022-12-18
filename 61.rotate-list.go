package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	len := 0
	newHead := &ListNode{Val: 0, Next: head}
	cur := newHead

	for cur.Next != nil {
		len++
		cur = cur.Next
	}

	// if k = lenth of list then no modification shall apply
	if (k % len) == 0 {
		return head
	}

	// at this moment cur is at the end of list
	// link the back to head
	cur.Next = head
	// reused cur as the iterator head
	cur = newHead

	for i := len - k%len; i > 0; i-- {
		cur = cur.Next
	}
	res := &ListNode{Val: 0, Next: cur.Next}
	cur.Next = nil
	return res.Next
}

// @lc code=end
