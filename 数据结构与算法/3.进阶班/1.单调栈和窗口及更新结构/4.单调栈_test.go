package class01

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ans := MonotonousStack([]int{3, 2, 1, 4, 5})
	fmt.Println(ans)
	return "Hello World!", nil
}

// MonotonousStack 单调栈结构
// arr [3, 2, 1, 4, 5]
// ans : [
//
//	   0: [-1, 1]
//	   1: [-1, 2]
//	]
//
// 小: 栈底到栈顶从小到大
func MonotonousStack(arr []int) [][2]int {
	// 结果集
	ans := make([][2]int, len(arr))

	// 单调栈结构
	stack := list.New()

	for i := 0; i < len(arr); i++ {
		// 栈不为空且栈顶元素比当前值大
		for stack.Len() > 0 && arr[stack.Front().Value.(*list.List).Back().Value.(int)] > arr[i] {
			// 弹出list
			queue := stack.Front().Value.(*list.List)
			stack.Remove(stack.Front())

			// 左边离它近比它小的
			leftIndex := -1
			if stack.Len() > 0 {
				leftIndex = stack.Front().Value.(*list.List).Back().Value.(int)
			}

			// 结算list
			for queue.Len() > 0 {
				cur := queue.Front().Value.(int)
				queue.Remove(queue.Front())
				ans[cur][0] = leftIndex
				ans[cur][1] = i
			}
		}

		// 如果栈顶元素和当前元素相等
		if stack.Len() > 0 && arr[stack.Front().Value.(*list.List).Back().Value.(int)] == arr[i] {
			stack.Front().Value.(*list.List).PushBack(i)
		} else {
			// 生成一个新队列压入栈中
			queue := list.New()
			queue.PushBack(i)
			stack.PushFront(queue)
		}
	}

	// 结算单调栈中剩余元素
	for stack.Len() > 0 {
		queue := stack.Front().Value.(*list.List)
		stack.Remove(stack.Front())

		leftIndex := -1
		if stack.Len() > 0 {
			leftIndex = stack.Front().Value.(*list.List).Back().Value.(int)
		}

		for queue.Len() > 0 {
			cur := queue.Front().Value.(int)
			queue.Remove(queue.Front())

			ans[cur][0] = leftIndex
			ans[cur][1] = -1 // 没有右边离它近且比它小的
		}
	}

	return ans
}
