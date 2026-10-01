package class11

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	stack := list.New()
	stack.PushFront(3)
	stack.PushFront(2)
	stack.PushFront(1)

	reverse(stack)

	for stack.Len() > 0 {
		cur := stack.Front().Value.(int)
		stack.Remove(stack.Front())
		fmt.Println(cur)
	}

	return "Hello World!", nil
}

func reverse(stack *list.List) {
	if stack.Len() == 0 {
		return
	}
	i := f(stack) // 收集当前值
	reverse(stack)
	stack.PushFront(i)
}

// 将栈底元素弹出
func f(stack *list.List) int {
	current := stack.Front().Value.(int)
	stack.Remove(stack.Front())

	if stack.Len() == 0 {
		return current
	} else {
		next := f(stack)
		stack.PushFront(current)
		return next
	}
}
