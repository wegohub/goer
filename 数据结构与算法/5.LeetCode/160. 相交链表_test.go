package leetcode

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	n := 0
	curA := headA
	for curA.Next != nil {
		n++
		curA = curA.Next
	}

	curB := headB
	for curB.Next != nil {
		n--
		curB = curB.Next
	}

	// 尾节点不相同，不相交
	if curA != curB {
		return nil
	}

	// curA长链表 curB短链表
	if n >= 0 {
		curA = headA
		curB = headB
	} else {
		curA = headB
		curB = headA
	}

	// 长链表先走 |n| 步
	for i := int(math.Abs(float64(n))); i > 0; i-- {
		curA = curA.Next
	}

	for curA != curB {
		curA = curA.Next
		curB = curB.Next
	}

	return curA
}
