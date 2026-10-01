package class03

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	nums := []int{5, 2, 6, 1}
	fmt.Println(countSmaller2(nums))
	return "Hello World!", nil
}

type Node struct {
	Value int
	Index int
}

func countSmaller2(nums []int) []int {
	ans := make([]int, len(nums))
	arr := make([]*Node, 0, len(nums))
	for index, num := range nums {
		arr = append(arr, &Node{
			Index: index,
			Value: num,
		})
	}
	MergeSort(arr, ans)
	return ans
}

func MergeSort(arr []*Node, ans []int) {
	if len(arr) == 0 {
		return
	}
	if len(arr) == 1 {
		ans[0] = 0
	}
	process(arr, 0, len(arr)-1, ans)
}

func process(arr []*Node, L, R int, ans []int) {
	if L >= R {
		return
	}
	M := (L + R) / 2
	process(arr, L, M, ans)
	process(arr, M+1, R, ans)
	merge(arr, L, M, R, ans)
}

func merge(arr []*Node, L, M, R int, ans []int) {
	help := make([]*Node, R-L+1)
	i := L
	j := M + 1
	index := 0
	for i <= M && j <= R {
		if arr[i].Value <= arr[j].Value {
			ans[arr[i].Index] += (j - 1) - (M + 1) + 1
			help[index] = arr[i]
			index++
			i++
		} else {
			help[index] = arr[j]
			j++
			index++
		}
	}

	for i <= M {
		ans[arr[i].Index] += R - (M + 1) + 1
		help[index] = arr[i]
		index++
		i++
	}

	for j <= R {
		help[index] = arr[j]
		j++
		index++
	}

	copy(arr[L:R+1], help)
}
