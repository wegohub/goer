package leetcode

import (
	"container/list"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	return "Hello World!", nil
}

// 单调栈
func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	maxArea := 0

	// 准备一个单调栈，存下标
	stack := list.New()
	for i := 0; i < len(heights); i++ {
		// 结算答案
		for stack.Len() > 0 && heights[stack.Back().Value.(int)] >= heights[i] {
			cur := stack.Back().Value.(int)
			stack.Remove(stack.Back())
			// 右边离它最近的最小的是i
			right := i
			// 左边离它最近的最小是它盖着的数，没有就是-1
			left := -1
			if stack.Len() > 0 {
				left = stack.Back().Value.(int)
			}
			// ((right - 1) - (left + 1)) + 1  当前弹出的数高度是最小的，其他位置都比它高，否则cur都被结算了
			curArea := (right - left - 1) * heights[cur]
			maxArea = int(math.Max(float64(curArea), float64(maxArea)))
		}
		stack.PushBack(i)
	}

	// 栈未空， 右边离它小的就没有
	for stack.Len() > 0 {
		cur := stack.Back().Value.(int)
		stack.Remove(stack.Back())
		// 右边离它最近的最小的是数组的长度
		right := len(heights)
		// 左边离它最近的最小是它盖着的数，没有就是-1
		left := -1
		if stack.Len() > 0 {
			left = stack.Back().Value.(int)
		}
		curArea := (right - left - 1) * heights[cur]
		maxArea = int(math.Max(float64(curArea), float64(maxArea)))
	}

	return maxArea
}

// 补单调栈经典实现
