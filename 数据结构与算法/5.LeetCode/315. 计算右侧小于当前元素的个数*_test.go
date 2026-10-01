package leetcode

import "fmt"

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
