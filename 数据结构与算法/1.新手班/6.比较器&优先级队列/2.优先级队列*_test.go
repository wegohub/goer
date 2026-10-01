package class05

import (
	"container/heap"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	h := &PriorityQueueInt{5, 1, 9, 3, 5, 1, 8, 7, 6, 4}
	heap.Init(h)
	heap.Push(h, 5)
	heap.Push(h, 1)
	heap.Push(h, 9)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 10)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}
	return "Hello World!", nil
}

type PriorityQueueInt []int

func (cls PriorityQueueInt) Len() int {
	return len(cls)
}

func (cls PriorityQueueInt) Less(i, j int) bool {
	return cls[i] > cls[j]
}

func (cls PriorityQueueInt) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (obj *PriorityQueueInt) Push(v interface{}) {
	*obj = append(*obj, v.(int))
}

func (obj *PriorityQueueInt) Pop() interface{} {
	old := *obj
	n := len(*obj)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// PriorityQueue 优先级队列
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // 防止内存泄漏
	*pq = old[0 : n-1]
	return item
}

// 自定义堆结构
type MyHeap struct {
	data []int
	size int
	cap  int
}

func (obj *MyHeap) Push(num int) {
	if obj.size >= obj.cap { // 堆满，扩容到原来大小的2倍
		obj.cap = obj.cap * 2
		newData := make([]int, obj.cap, obj.cap)
		copy(newData, obj.data)
		obj.data = newData
	}
	obj.data[obj.size] = num
	obj.heapInsert(obj.size)
	obj.size++
}

func (obj *MyHeap) Pop() int {
	if obj.size == 0 {
		fmt.Println("堆为空")
		return -1
	}
	ans := obj.data[0]
	obj.data[0] = obj.data[obj.size-1]
	obj.size--
	obj.heapify(0)
	return ans
}

func (obj *MyHeap) Size() int {
	return obj.size
}

func (obj *MyHeap) IsEmpty() bool {
	return obj.size == 0
}

func (obj *MyHeap) Print() {
	fmt.Println(obj.data)
}

// index 从那个位置开始insert(向上看)
func (obj *MyHeap) heapInsert(index int) {
	head := (index - 1) / 2
	for obj.data[head] > obj.data[index] {
		obj.data[head], obj.data[index] = obj.data[index], obj.data[head]
		index = head
		head = (index - 1) / 2
	}
}

// index 从那个位置开始heapify(向下看)
func (obj *MyHeap) heapify(index int) {
	left := 2*index + 1
	for left < obj.size {
		target := left
		// 左右孩子谁小是谁
		if left+1 < obj.size && obj.data[left+1] < obj.data[left] {
			target = left + 1
		}
		// 两个孩子都比父节点大
		if obj.data[target] > obj.data[index] {
			break
		}
		// 交换
		obj.data[index], obj.data[target] = obj.data[target], obj.data[index]
		// index来到大的孩子继续往下看
		index = target
		left = 2*index + 1
	}
}
