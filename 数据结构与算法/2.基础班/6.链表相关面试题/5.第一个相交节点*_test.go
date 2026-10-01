package class06

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 两个单链表相交的一系列问题
// head1和head2可能有环，给定两个单链表的头节点head1和head2, 这个两个链表有可能相交，有可能不相交。请实现一个函数，如果两个链表相交，返回相交的第一个节点，如果不相交返回nil.
// 要求：如果链表1的长度为N, 链表2的长度为M, 时间复杂度要达到O(N+M), 空间复杂度要求O(1)
func getLoopNode(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil || head2 == nil {
		return nil
	}
	ring1 := getSingleLinkLoopNode(head1)
	ring2 := getSingleLinkLoopNode(head2)

	// 一个有环，一个无环，不可能相交
	if (ring1 == nil && ring2 != nil) || (ring1 != nil && ring2 == nil) {
		return nil
	}

	// 都无环
	if ring1 == nil && ring2 == nil {
		return getNoRingLinkLoopNode(head1, head2)
	}

	// 都有环
	// 1. 入环节点都相同
	if ring1 == ring2 {
		return getRingLinkLoopNode(head1, head2, ring1)
	}

	// 2. 入环节点不同
	// 看链表1走一圈能不能遇到链表2
	current := ring1.Next
	for current != ring1 {
		// 遇到链表2直接返回
		if current == ring2 {
			return ring1
		}
		current = current.Next
	}
	// 走一圈都没有遇到，两个链表不相交
	return nil
}

// 判断单链表是否存在环，存在返回第一个入环节点, 不存在返回nil
func getSingleLinkLoopNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			fast = head
			break
		}
	}
	// 无环
	if fast.Next == nil || fast.Next.Next == nil {
		return nil
	}

	// 有环
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}

	return fast
}

// 获取两个无环链表的第一个相交节点
func getNoRingLinkLoopNode(head1, head2 *ListNode) *ListNode {
	length := 0
	current1 := head1
	for current1.Next != nil {
		length++
		current1 = current1.Next
	}
	current2 := head2
	for current2.Next != nil {
		length--
		current2 = current2.Next
	}
	// 尾节点不相同，不相交
	if current1 != current2 {
		return nil
	}
	// 长链表
	longLink := head1
	// 短链表
	sortLink := head2
	if length < 0 {
		longLink = head2
		sortLink = head1
	}

	// 长链表先走差值
	for i := 0; i < int(math.Abs(float64(length))); i++ {
		longLink = longLink.Next
	}

	// 一直走到相交
	for longLink != sortLink {
		longLink = longLink.Next
		sortLink = sortLink.Next
	}

	return longLink
}

// 获取两个有环链表的第一个相交节点
func getRingLinkLoopNode(head1, head2, endNode *ListNode) *ListNode {
	length := 0
	current1 := head1
	for current1 != endNode {
		length++
		current1 = current1.Next
	}
	current2 := head2
	for current2 != endNode {
		length--
		current2 = current2.Next
	}
	longLink := head1
	sortLink := head2
	if length < 0 {
		longLink = head2
		sortLink = head1
	}

	// 长链表走差值步
	for i := 0; i < int(math.Abs(float64(length))); i++ {
		longLink = longLink.Next
	}

	for longLink != sortLink {
		longLink = longLink.Next
		sortLink = sortLink.Next
	}

	return longLink
}
