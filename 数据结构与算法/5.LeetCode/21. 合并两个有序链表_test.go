package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var (
		ans *ListNode
		pre *ListNode
	)
	if list1.Val < list2.Val {
		ans = list1
		pre = list1
		list1 = list1.Next
	} else {
		ans = list2
		pre = list2
		list2 = list2.Next
	}

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			pre.Next = list1
			pre = list1
			list1 = list1.Next
		} else {
			pre.Next = list2
			pre = list2
			list2 = list2.Next
		}
	}

	if list1 != nil {
		pre.Next = list1
	}

	if list2 != nil {
		pre.Next = list2
	}

	return ans
}
