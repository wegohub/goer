package class06

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	head := ArrToLink([]int{1, 5, 3, 2, 7, 5, 4, 4, 6, 5, 8, 9, 7})
	ans := partition(head, 2)
	PrintLink(ans)

	return "Hello World!", nil
}

func partition(head *ListNode, num int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var (
		lessHead  *ListNode // 小于区域的头
		lessTail  *ListNode // 小于区域的尾
		equalHead *ListNode // 等于区域的头
		equalTail *ListNode // 等于区域的尾
		moreHead  *ListNode // 大于区域的头
		moreTail  *ListNode // 大于区域的尾
	)

	cur := head
	for cur != nil {
		if cur.Val < num {
			if lessHead == nil {
				lessHead = cur
				lessTail = cur
			} else {
				lessTail.Next = cur
				lessTail = cur
			}
		} else if cur.Val == num {
			if equalHead == nil {
				equalHead = cur
				equalTail = cur
			} else {
				equalTail.Next = cur
				equalTail = cur
			}
		} else {
			if moreHead == nil {
				moreHead = cur
				moreTail = cur
			} else {
				moreTail.Next = cur
				moreTail = cur
			}
		}
		cur = cur.Next
	}

	// 小于区域不是空的
	if lessHead != nil {
		lessTail.Next = equalHead
		// 如果等于区域是空的，设置等于区域的尾巴为小于区域的尾巴
		if equalHead == nil {
			equalTail = lessTail
		}
		// 等于区域的尾巴连大于区域的头
		equalTail.Next = moreHead
	}

	// 确定返回头
	if lessHead != nil {
		return lessHead
	}
	if equalHead != nil {
		return equalHead
	}
	return moreHead
}
