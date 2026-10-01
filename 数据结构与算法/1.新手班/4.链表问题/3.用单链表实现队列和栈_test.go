package class04

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 测试队列
	queue := &Queue{}
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	queue.Push(5)
	for !queue.IsEmpty() {
		fmt.Println(queue.Pop())
	}
	fmt.Println(queue.IsEmpty())
	// 测试栈
	fmt.Println("=============================")
	stack := &Stack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	for !stack.IsEmpty() {
		fmt.Println(stack.Pop())
	}
	fmt.Println(stack.IsEmpty())

	return "Hello World!", nil
}

// 单链表实现队列(从尾部加入，从头部弹出，两个指针，O(1))
type Queue struct {
	head *ListNode
	tail *ListNode
	size int
}

// Push 加入队列
func (obj *Queue) Push(val int) {
	obj.size++
	node := &ListNode{Val: val}
	if obj.head == nil {
		obj.head = node
		obj.tail = node
	} else {
		obj.tail.Next = node
		obj.tail = node
	}
}

// Pop 出队列
func (obj *Queue) Pop() int {
	if obj.IsEmpty() {
		panic("queue is empty")
	}
	obj.size--
	val := obj.head.Val
	obj.head = obj.head.Next
	if obj.head == nil {
		obj.tail = nil
	}
	return val
}

// IsEmpty 队列是否为空
func (obj *Queue) IsEmpty() bool {
	return obj.size == 0
}

// 单链表实现栈
type Stack struct {
	head *ListNode
	size int
}

// Push 入栈
func (obj *Stack) Push(val int) {
	obj.size++
	node := &ListNode{Val: val}
	if obj.head == nil {
		obj.head = node
	} else {
		node.Next = obj.head
		obj.head = node
	}
}

// Pop 出栈
func (obj *Stack) Pop() int {
	if obj.IsEmpty() {
		panic("stack is empty")
	}
	obj.size--
	val := obj.head.Val
	obj.head = obj.head.Next
	return val
}

// Peek 栈顶元素
func (obj *Stack) Peek() int {
	if obj.IsEmpty() {
		panic("stack is empty")
	}
	return obj.head.Val
}

// IsEmpty 栈是否为空
func (obj *Stack) IsEmpty() bool {
	return obj.size == 0
}
