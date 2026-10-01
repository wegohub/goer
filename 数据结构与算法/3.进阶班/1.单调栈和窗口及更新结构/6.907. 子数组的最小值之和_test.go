package class01

import (
	"container/list"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{3, 1, 2, 2, 2, 2, 2, 2, 4}
	ans := sumSubarrayMins(arr)
	ansRight := right(arr)
	fmt.Println(ans, ansRight)
	return "Hello World!", nil
}

// 滑动窗口
func right(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := 0
	N := len(arr)
	for L := 0; L < N; L++ {
		qMin := list.New()
		for R := L; R < N; R++ {
			for qMin.Len() > 0 && arr[qMin.Back().Value.(int)] >= arr[R] {
				qMin.Remove(qMin.Back())
			}
			qMin.PushBack(R)
			ans += arr[qMin.Front().Value.(int)]
		}
	}
	m := int(math.Pow10(9)) + 7
	return ans % m
}

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	if n == 0 {
		return 0
	}
	monoArr := getMono(arr)
	ans := 0
	for index, item := range arr {
		leftIndex := monoArr[index][0]
		rightIndex := n
		if monoArr[index][1] != -1 {
			rightIndex = monoArr[index][1]
		}
		ans += (index - leftIndex) * (rightIndex - index) * item // 有多少个以cur为最小值的子数组
	}
	m := int(math.Pow10(9)) + 7
	return ans % m
}

func getMono(arr []int) [][2]int {
	stack := list.New()
	ans := make([][2]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for stack.Len() > 0 && arr[stack.Back().Value.(int)] >= arr[i] {
			cur := stack.Back().Value.(int)
			stack.Remove(stack.Back())

			leftIndex := -1
			if stack.Len() > 0 {
				leftIndex = stack.Back().Value.(int)
			}
			ans[cur] = [2]int{leftIndex, i}
		}
		stack.PushBack(i)
	}

	for stack.Len() > 0 {
		cur := stack.Back().Value.(int)
		stack.Remove(stack.Back())

		leftIndex := -1
		if stack.Len() > 0 {
			leftIndex = stack.Back().Value.(int)
		}

		ans[cur] = [2]int{leftIndex, -1}
	}

	return ans
}
