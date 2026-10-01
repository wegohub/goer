package leetcode

import "container/list"

// 1 1 1 1 1 85 96 114 142 221 226 337 338 448 494 543 560 581 617 621 647 739 763
// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		var pre *Node
		for i := 0; i < size; i++ {
			cur := queue.Front().Value.(*Node)
			queue.Remove(queue.Front())

			if cur.Left != nil {
				queue.PushBack(cur.Left)
			}

			if cur.Right != nil {
				queue.PushBack(cur.Right)
			}

			// 连线
			if pre != nil {
				pre.Next = cur
			}
			pre = cur
		}
		// 为下一层做准备
		pre = nil
	}
	return root
}
