package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// O(1) space O(n) time
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	// preSlow is the pointer before the one need to be removed
	// slow is the one to remove
	// fast is the one to detect reach of the end
	preSlow, slow, fast := dummyHead, head, head

	// nth node from end so make fast reach the end and stop
	for fast != nil {
		if n <= 0 {
			preSlow = slow
			slow = slow.Next
		}
		// wait nth and start moving the slow pointer
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummyHead.Next
}

// @lc code=end
