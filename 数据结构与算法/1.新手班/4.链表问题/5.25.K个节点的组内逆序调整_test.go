package class04

import (
	. "backend/utils/algo"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	head := ArrToLink([]int{1, 2, 3, 4})

	ans := reverseKGroup(head, 2)

	PrintLink(ans)

	return "Hello World!", nil
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 先看第一组够不够
	end := kEndNode(head, k)
	if end == nil {
		return head
	}

	start := head
	nextStart := end.Next
	ans := reverseListNode(head, nextStart)

	for nextStart != nil {
		lastStart := start
		start = nextStart
		end = kEndNode(start, k)
		// 不够k个了，上一组的头连当前组的头
		if end == nil {
			nextStart = nil
			lastStart.Next = start
		} else { // 上一组的头连当前组的尾
			nextStart = end.Next
			lastStart.Next = reverseListNode(start, nextStart)
		}
	}

	return ans
}

// 给一个头节点数够k个，返回第k个节点
func kEndNode(head *ListNode, k int) *ListNode {
	var (
		ans   *ListNode
		index = 0
	)
	for head != nil {
		index++
		if index == k {
			ans = head
			break
		}
		head = head.Next
	}
	return ans
}

// 反转单链表,返回新头部
func reverseListNode(head *ListNode, stop *ListNode) *ListNode {
	var pre *ListNode
	for head != stop {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
