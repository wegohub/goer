package class04

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	dueue := &Dueue{}

	// 头进尾出(队列)
	dueue.PushHead(1)
	dueue.PushHead(2)
	dueue.PushHead(3)
	dueue.PushHead(4)
	dueue.PushHead(5)
	for dueue.Size() > 0 {
		fmt.Println(dueue.PopTail())
	}
	fmt.Println("queue size: ", dueue.Size())
	fmt.Println("=========================")
	// 尾进头出(队列)
	dueue.PushTail(1)
	dueue.PushTail(2)
	dueue.PushTail(3)
	dueue.PushTail(4)
	dueue.PushTail(5)
	for dueue.Size() > 0 {
		fmt.Println(dueue.PopHead())
	}
	fmt.Println("queue size: ", dueue.Size())
	fmt.Println("=========================")
	// 头进头出(栈)
	dueue.PushHead(1)
	dueue.PushHead(2)
	dueue.PushHead(3)
	dueue.PushHead(4)
	dueue.PushHead(5)
	for dueue.Size() > 0 {
		fmt.Println(dueue.PopHead())
	}
	fmt.Println("stack size: ", dueue.Size())
	fmt.Println("=========================")
	// 尾进尾出(栈)
	dueue.PushTail(1)
	dueue.PushTail(2)
	dueue.PushTail(3)
	dueue.PushTail(4)
	dueue.PushTail(5)
	for dueue.Size() > 0 {
		fmt.Println(dueue.PopTail())
	}
	fmt.Println("stack size: ", dueue.Size())
	fmt.Println("=========================")

	return "Hello World!", nil
}

// 双端链表
type Dueue struct {
	head *DoubleLinkNode
	tail *DoubleLinkNode
	size int
}

// PushHead 头进
func (obj *Dueue) PushHead(val int) {
	obj.size++
	node := &DoubleLinkNode{Val: val}
	if obj.head == nil {
		obj.head = node
		obj.tail = node
	} else {
		node.Next = obj.head
		obj.head.Last = node
		obj.head = node
	}
}

// PopTail 尾出
func (obj *Dueue) PopTail() int {
	if obj.Size() == 0 {
		panic("Dueue is empty")
	}
	obj.size--
	val := obj.tail.Val
	obj.tail = obj.tail.Last
	if obj.tail == nil {
		obj.head = nil
	} else {
		obj.tail.Next = nil
	}
	return val
}

// PushTail 尾进
func (obj *Dueue) PushTail(val int) {
	obj.size++
	node := &DoubleLinkNode{Val: val}
	if obj.tail == nil {
		obj.head = node
		obj.tail = node
	} else {
		obj.tail.Next = node
		node.Last = obj.tail
		obj.tail = node
	}
}

// PopHead 头出
func (obj *Dueue) PopHead() int {
	if obj.Size() == 0 {
		panic("Dueue is empty")
	}
	obj.size--
	val := obj.head.Val
	obj.head = obj.head.Next
	if obj.head == nil {
		obj.tail = nil
	} else {
		obj.head.Last = nil
	}
	return val
}

// Size 队列大小
func (obj *Dueue) Size() int {
	return obj.size
}
