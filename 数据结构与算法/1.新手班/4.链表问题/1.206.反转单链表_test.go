package class04

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	testTimes := 1
	for i := 0; i < testTimes; i++ {
		arr := GenArray(20, 100)
		head := ArrToLink(arr)
		ans := reverseList(head)

		// 检查长度
		ansLen := len(LinkToArr(ans))
		if ansLen != len(arr) {
			fmt.Println("err: ", arr)
			return "Error Len", nil
		}

		// 检查值
		cur := ans
		index := len(arr) - 1
		for cur != nil {
			if cur.Val != arr[index] {
				fmt.Println("err: ", arr)
				return "Error", nil
			}
			cur = cur.Next
			index--
		}
	}
	return "success!", nil
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		next := head.Next
		head.Next = pre
		pre = head
		head = next
	}
	return pre
}
