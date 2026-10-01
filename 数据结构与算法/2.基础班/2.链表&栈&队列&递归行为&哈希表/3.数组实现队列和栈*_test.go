package class02

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println("测试数组栈开始")
	stack := NewArrStack(10)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
	fmt.Println(stack.IsEmpty())
	fmt.Println("测试数组栈结束\n====================")

	fmt.Println("测试数组队列开始")
	queue := NewArrQueue(10)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	for !queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}
	fmt.Println(queue.IsEmpty())
	fmt.Println("测试数组队列结束\n====================")

	return "Hello World!", nil
}

func NewArrStack(maxSize int) *ArrStack {
	return &ArrStack{
		data:    make([]int, maxSize),
		size:    0,
		maxSize: maxSize,
	}
}

// ArrStack 数组实现栈
type ArrStack struct {
	data    []int
	size    int
	maxSize int
}

func (obj *ArrStack) IsEmpty() bool {
	return obj.size == 0
}

func (obj *ArrStack) Size() int {
	return obj.size
}

func (obj *ArrStack) Push(v int) {
	if obj.size == obj.maxSize {
		panic("stack is full")
	}
	obj.data[obj.size] = v
	obj.size++
}

func (obj *ArrStack) Pop() int {
	if obj.IsEmpty() {
		panic("stack is empty")
	}
	obj.size--
	return obj.data[obj.size]
}

func (obj *ArrStack) Peek() int {
	if obj.IsEmpty() {
		panic("stack is empty")
	}
	return obj.data[obj.size-1]
}

func NewArrQueue(maxSize int) *ArrQueue {
	return &ArrQueue{
		maxSize:   maxSize,
		data:      make([]int, maxSize),
		pushIndex: 0,
		popIndex:  0,
		size:      0,
	}
}

// ArrQueue 数组实现队列
type ArrQueue struct {
	maxSize   int
	data      []int
	pushIndex int
	popIndex  int
	size      int
}

func (obj *ArrQueue) IsEmpty() bool {
	return obj.size == 0
}

func (obj *ArrQueue) IsFull() bool {
	return obj.size == obj.maxSize
}

func (obj *ArrQueue) Size() int {
	return obj.size
}

func (obj *ArrQueue) Push(v int) {
	if obj.IsFull() {
		panic("queue is full")
	}
	obj.size++
	obj.data[obj.pushIndex] = v
	obj.pushIndex = (obj.pushIndex + 1) % obj.maxSize
}

func (obj *ArrQueue) Pop() int {
	if obj.IsEmpty() {
		panic("queue is empty")
	}
	obj.size--
	v := obj.data[obj.popIndex]
	obj.popIndex = (obj.popIndex + 1) % obj.maxSize
	return v
}
