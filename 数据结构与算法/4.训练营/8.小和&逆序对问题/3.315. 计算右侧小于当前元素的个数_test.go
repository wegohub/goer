package pmerge

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{5, 2, 6, 1}
	ans := countSmaller(arr)
	fmt.Println(arr)
	fmt.Println(ans)
	return "Hello World!", nil
}

type Node struct {
	Value int
	Index int
}

func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	if len(nums) < 2 {
		return ans
	}
	var arr []*Node
	for index, num := range nums {
		arr = append(arr, &Node{Value: num, Index: index})
	}

	process(arr, 0, len(arr)-1, ans)
	return ans
}

func process(arr []*Node, L, R int, ans []int) {
	if L == R {
		return
	}
	mid := L + ((R - L) >> 1)
	process(arr, L, mid, ans)
	process(arr, mid+1, R, ans)
	merge1(arr, L, mid, R, ans)
}

// 右侧小于当前元素的个数
func merge1(arr []*Node, L, M, R int, ans []int) {
	help := make([]*Node, R-L+1)
	i := L
	j := M + 1
	index := 0
	for i <= M && j <= R {
		if arr[i].Value <= arr[j].Value {
			ans[arr[i].Index] += (j - 1) - (M + 1) + 1
			help[index] = arr[i]
			i++
			index++
		} else {
			help[index] = arr[j]
			j++
			index++
		}
	}

	for i <= M {
		ans[arr[i].Index] += R - (M + 1) + 1
		help[index] = arr[i]
		i++
		index++
	}

	for j <= R {
		help[index] = arr[j]
		j++
		index++
	}
	copy(arr[L:R+1], help)
}

// 右侧大于当前元素的个数
func merge2(arr []*Node, L, M, R int, ans []int) {
	help := make([]*Node, R-L+1)
	i := L
	j := M + 1
	index := 0
	for i <= M && j <= R {
		if arr[i].Value < arr[j].Value {
			ans[arr[i].Index] += R - j + 1
			help[index] = arr[i]
			i++
			index++
		} else {
			help[index] = arr[j]
			j++
			index++
		}
	}

	for i <= M {
		help[index] = arr[i]
		i++
		index++
	}

	for j <= R {
		help[index] = arr[j]
		j++
		index++
	}
	copy(arr[L:R+1], help)
}

func merge(arr []*Node, L, M, R int, ans []int) {
	help := make([]*Node, R-L+1)
	index := len(help) - 1
	left := M
	right := R
	for left >= L && right >= M+1 {
		if arr[left].Value > arr[right].Value {
			help[index] = arr[left]
			ans[arr[left].Index] = ans[arr[left].Index] + (right - M)
			left--
		} else {
			help[index] = arr[right]
			right--
		}
		index--
	}
	for left >= L {
		help[index] = arr[left]
		left--
		index--
	}
	for right >= M+1 {
		help[index] = arr[right]
		right--
		index--
	}
	for i, item := range help {
		arr[L+i] = item
	}
}
