package class06

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 对数器验证
	testTimes := 10000
	for i := 0; i < testTimes; i++ {
		arr := GenArray(100, 1000)
		link := ArrToLink(arr)
		// 中点/上中点
		topPoint := (len(arr) - 1) / 2
		topPointLink := TopMidpoint(link)
		if topPointLink == nil && topPoint != 0 {
			fmt.Println("中点/上中点err: ", arr)
			break
		}
		if topPointLink != nil && topPointLink.Val != arr[topPoint] {
			fmt.Println("中点/上中点err: ", arr)
			break
		}

		// 中点/下中点
		lowPoint := len(arr) / 2
		lowPointLink := LowerMidpoint(link)
		if lowPointLink == nil && lowPoint != 0 {
			fmt.Println("中点/下中点err: ", arr)
			break
		}
		if lowPointLink != nil && lowPointLink.Val != arr[lowPoint] {
			fmt.Println("中点/下中点err: ", arr)
			break
		}

		// 中点/上中点 的前一个
		topPointFront := ((len(arr) - 1) / 2) - 1
		if topPointFront < 0 {
			topPointFront = 0
		}
		topPointFrontLink := ThePreviousOneAtTheTopMidpoint(link)
		if topPointFrontLink == nil && topPointFront != 0 {
			fmt.Println("中点/上中点 的前一个err: ", arr)
			break
		}
		if topPointFrontLink != nil && topPointFrontLink.Val != arr[topPointFront] {
			fmt.Println("中点/上中点 的前一个err: ", arr)
			break
		}

		// 中点/下中点 的前一个
		lowPointFront := (len(arr) / 2) - 1
		if lowPointFront < 0 {
			lowPointFront = 0
		}
		lowPointFrontLink := ThePreviousOneAtTheLowerMidpoint(link)
		if lowPointFrontLink == nil && lowPointFront != 0 {
			fmt.Println("中点/下中点 的前一个err: ", arr)
			break
		}
		if lowPointFrontLink != nil && lowPointFrontLink.Val != arr[lowPointFront] {
			fmt.Println("中点/下中点 的前一个err: ", arr)
			break
		}

	}

	return "Success!", nil
}

// 快慢指针灵魂四问

// 1) 输入链表头节点，奇数长度返回中点，偶数长度返回上中点
func TopMidpoint(head *ListNode) *ListNode {
	// 只有两个点
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	// >2个点
	slow := head.Next
	fast := head.Next.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 2) 输入链表头节点，奇数长度返回中点，偶数长度返回下中点
func LowerMidpoint(head *ListNode) *ListNode {
	// 只有一个点
	if head == nil || head.Next == nil {
		return head
	}
	// >1个点
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 3) 输入链表头节点，奇数长度返回中点前一个，偶数长度返回上中点前一个
func ThePreviousOneAtTheTopMidpoint(head *ListNode) *ListNode {
	// 只有一个点
	if head == nil || head.Next == nil {
		return nil
	}
	// 只有两个点
	if head.Next.Next == nil {
		return head
	}
	// >2 个点
	pre := head
	slow := head.Next
	fast := head.Next.Next

	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	return pre
}

// 4) 输入链表头节点，奇数长度返回中点前一个，偶数长度返回下中点前一个
func ThePreviousOneAtTheLowerMidpoint(head *ListNode) *ListNode {
	// 只有一个数
	if head == nil || head.Next == nil {
		return nil
	}
	// >1个数
	pre := head
	slow := head.Next
	fast := head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	return pre
}
