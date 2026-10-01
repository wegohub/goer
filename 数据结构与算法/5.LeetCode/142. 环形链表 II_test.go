package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func detectCycle(head *ListNode) *ListNode {
	// 有0、1、2个节点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return nil
	}

	slow := head.Next
	fast := head.Next.Next
	for slow != fast {
		if fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}
