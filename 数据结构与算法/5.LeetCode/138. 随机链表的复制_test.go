package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	// 先将原链表的每个节点复制，反正下一个节点的位置
	cur := head
	for cur != nil {
		next := cur.Next
		copyNode := &Node{Val: cur.Val}
		cur.Next = copyNode
		copyNode.Next = next
		cur = next
	}

	// 设置random指针
	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	// 分离链表
	ans := head.Next
	cur = head
	for cur != nil {
		next := cur.Next.Next
		copyNode := cur.Next
		cur.Next = next
		if next != nil {
			copyNode.Next = next.Next
		}
		cur = next
	}
	return ans
}
