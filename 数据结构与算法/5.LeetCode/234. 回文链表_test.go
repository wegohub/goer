package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func isPalindrome(head *ListNode) bool {
	// 获取到链表上中点
	// 只有1个点
	if head == nil || head.Next == nil {
		return true
	}
	// 只有2个点
	if head.Next.Next == nil {
		if head.Val == head.Next.Val {
			return true
		} else {
			return false
		}
	}

	// > 2个点
	slow := head.Next
	fast := head.Next.Next

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	start := slow.Next
	slow.Next = nil

	// 从start开始反转链表
	var pre *ListNode // 反转后半部分的链表新头部
	cur := start
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	l1 := head
	l2 := pre
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	// 将后半部分链表调回来
	cur = pre
	pre = nil
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	// 将原链表连起来
	slow.Next = start

	return true
}
