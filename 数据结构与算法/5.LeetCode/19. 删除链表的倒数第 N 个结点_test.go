package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	fast := head
	slow := head
	pre := head
	// fast先走n步
	for i := 0; i < n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}
	if fast == nil {
		return head.Next
	}

	for fast != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next
	}

	pre.Next = pre.Next.Next
	return head
}
