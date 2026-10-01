package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	long := l1
	short := l2
	if getListLen(l2) > getListLen(l1) {
		long = l2
		short = l1
	}

	// 将结果压入长链表
	ans := long
	pre := long

	// 进位
	carry := 0

	// 长短都有
	for short != nil {
		total := long.Val + short.Val + carry
		carry = total / 10
		long.Val = total % 10
		short = short.Next
		pre = long
		long = long.Next
	}

	// 只有长链表
	for long != nil {
		total := long.Val + carry
		carry = total / 10
		long.Val = total % 10
		pre = long
		long = long.Next
	}

	// 进位我不为0
	if carry > 0 && pre != nil {
		pre.Next = &ListNode{Val: carry}
	}
	return ans
}

func getListLen(head *ListNode) int {
	ans := 0
	for head != nil {
		ans++
		head = head.Next
	}
	return ans
}
