package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	head := ArrToLink([]int{1, 2, 3, 4, 5})
	PrintLink(head)
	ans := reverseBetween(head, 1, 2)
	PrintLink(ans)
	return "Hello World!", nil
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	cur := head
	var find *ListNode    // 找到left位置的节点
	var findPre *ListNode // left位置节点的上一个节点
	var end *ListNode     // right+1 位置
	var pre *ListNode
	index := 1
	for cur != nil {
		// 找到反转的起点
		if index == left {
			findPre = pre
			find = cur
			pre = nil
		}

		// 来到末尾位置
		if index == right+1 {
			end = cur
			break
		}

		// 执行反转逻辑
		if find != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		} else { // 执行寻找left逻辑
			pre = cur
			cur = cur.Next
		}

		index++
	}

	// 将链表串起来
	// left == 1
	if findPre == nil {
		head = pre
	} else {
		findPre.Next = pre
	}

	// end == nil => right == n
	if end != nil {
		find.Next = end
	}

	return head
}
