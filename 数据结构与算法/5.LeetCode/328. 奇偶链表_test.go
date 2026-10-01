package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}

	// 奇头
	oddHead := head
	// 奇尾
	oddTail := head
	// 偶头
	evenHead := head.Next
	// 偶尾
	evenTail := head.Next

	index := 3
	cur := head.Next.Next
	for cur != nil {
		if index&1 == 0 { // 偶数
			evenTail.Next = cur
			evenTail = cur
		} else { // 奇数
			oddTail.Next = cur
			oddTail = cur
		}
		cur = cur.Next
		index++
	}

	// 奇偶连起来
	oddTail.Next = evenHead
	// 末尾一定要断开，不然就死循环了
	// 比如：1 2 3 4 5
	// 迭代结束后： 1 3 5
	//            2 4->5
	// 这道题的考点就在这里
	evenTail.Next = nil

	return oddHead
}

func oddEvenList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	var pre *ListNode
	odd := head
	even := head.Next
	for odd != nil && odd.Next != nil {
		next := odd.Next.Next
		evenNode := odd.Next
		odd.Next = next
		if next != nil {
			evenNode.Next = next.Next
		}
		pre = odd
		odd = next
	}
	if odd != nil {
		odd.Next = even
	} else {
		pre.Next = even
	}
	return head
}
