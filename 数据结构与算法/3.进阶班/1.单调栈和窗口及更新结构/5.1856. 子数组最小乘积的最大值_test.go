package class01

import (
	"container/list"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{1, 2, 3, 2}
	ans := maxSumMinProduct(arr)
	fmt.Println(ans)

	return "Hello World!", nil
}

func maxSumMinProduct(arr []int) int {
	// 前缀和数组
	sum := make([]int, len(arr))
	sum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		sum[i] = arr[i] + sum[i-1]
	}
	// 单调栈更新结构
	ans := math.MinInt
	n := len(arr)

	stack := list.New()

	for i := 0; i < len(arr); i++ {
		for stack.Len() > 0 && arr[stack.Back().Value.(int)] >= arr[i] {
			cur := stack.Back().Value.(int)
			stack.Remove(stack.Back())

			leftIndex := 0
			if stack.Len() > 0 {
				leftIndex = stack.Back().Value.(int) + 1
			}

			curSum := sum[i-1]
			if leftIndex > 0 {
				curSum = sum[i-1] - sum[leftIndex-1]
			}

			ans = Max(ans, curSum*arr[cur])
		}
		stack.PushBack(i)
	}

	for stack.Len() > 0 {
		cur := stack.Back().Value.(int)
		stack.Remove(stack.Back())

		leftIndex := 0
		if stack.Len() > 0 {
			leftIndex = stack.Back().Value.(int) + 1
		}

		curSum := sum[n-1]
		if leftIndex > 0 {
			curSum = sum[n-1] - sum[leftIndex-1]
		}

		ans = Max(ans, curSum*arr[cur])
	}
	m := int(math.Pow10(9)) + 7
	return ans % m
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
