package class04

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	head := &DoubleLinkNode{Val: 1}
	node1 := &DoubleLinkNode{Val: 2}
	node2 := &DoubleLinkNode{Val: 3}

	head.Next = node1
	node1.Next = node2

	node2.Last = node1
	node1.Last = head

	ans := reverseDoubleList(head)

	cur := ans
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}

	return "Hello World!", nil
}

func reverseDoubleList(head *DoubleLinkNode) *DoubleLinkNode {
	var pre *DoubleLinkNode
	for head != nil {
		next := head.Next
		head.Next = pre
		head.Last = next
		pre = head
		head = next
	}
	return pre
}
