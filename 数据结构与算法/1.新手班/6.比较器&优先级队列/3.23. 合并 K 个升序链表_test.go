package class05

import (
	. "backend/utils/algo"
	"container/heap"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	lists := []*ListNode{
		ArrToLink([]int{1, 4, 5}),
		ArrToLink([]int{1, 3, 4}),
		ArrToLink([]int{2, 6}),
	}
	ans := mergeKLists(lists)
	PrintLink(ans)
	return "Hello World!", nil
}

func mergeKLists(lists []*ListNode) *ListNode {
	h := &PriorityQueue{}
	// 先将每个链表的头部加入小根堆
	for _, item := range lists {
		if item != nil {
			heap.Push(h, item)
		}
	}

	var (
		ans *ListNode
		pre *ListNode
	)
	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		if pre == nil {
			ans = node
			pre = node
		} else {
			pre.Next = node
			pre = node
		}
		// 将node下一个节点加入小根堆
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return ans
}

type PriorityQueue []*ListNode

func (cls PriorityQueue) Len() int {
	return len(cls)
}

func (cls PriorityQueue) Less(i, j int) bool {
	return cls[i].Val < cls[j].Val
}

func (cls PriorityQueue) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (obj *PriorityQueue) Push(node interface{}) {
	*obj = append(*obj, node.(*ListNode))
}

func (obj *PriorityQueue) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}
