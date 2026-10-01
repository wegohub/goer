package class06

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5

	copyHead := copyRandomList(node1)
	cur := copyHead
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}

	return "Hello World!", nil
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
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
