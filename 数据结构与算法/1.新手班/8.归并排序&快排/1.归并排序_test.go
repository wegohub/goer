package class08

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	arr := []int{6, 5, 4, 3, 2, 1}
	MergeSortIter(arr)
	fmt.Println(arr)

	return "Hello World!", nil
}

func MergeSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L >= R {
		return
	}
	mid := (L + R) / 2
	process(arr, L, mid)
	process(arr, mid+1, R)
	merge(arr, L, mid, R)
}

func merge(arr []int, L, M, R int) {
	ans := make([]int, R-L+1)
	index := 0
	i := L
	j := M + 1
	for i <= M && j <= R {
		if arr[i] <= arr[j] {
			ans[index] = arr[i]
			i++
			index++
		}
		if arr[j] < arr[i] {
			ans[index] = arr[j]
			j++
			index++
		}
	}

	for i <= M {
		ans[index] = arr[i]
		i++
		index++
	}

	for j <= R {
		ans[index] = arr[j]
		j++
		index++
	}

	// copy(arr[L:R+1], ans)
	for i, item := range ans {
		arr[L+i] = item
	}
}

func MergeSortIter(arr []int) {
	if len(arr) < 2 {
		return
	}

	N := len(arr)
	mergeSize := 1

	for mergeSize < N {
		L := 0
		for L < N {
			// 左组都凑不齐或刚好凑齐
			if mergeSize >= N-L {
				break
			}
			M := L + mergeSize - 1
			R := M + Min(mergeSize, N-M-1)
			Merge(arr, L, M, R)
			L = R + 1
		}

		// 防止溢出
		if mergeSize > N/2 {
			break
		}

		mergeSize <<= 1
	}
}

func Merge(arr []int, L, M, R int) {
	help := make([]int, R-L+1)
	i := L
	j := M + 1
	index := 0
	for i <= M && j <= R {
		if arr[i] <= arr[j] {
			help[index] = arr[i]
			i++
		} else {
			help[index] = arr[j]
			j++
		}
		index++
		//if arr[j] < arr[i] {
		//    help[index] = arr[j]
		//    j++
		//    index++
		//}
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

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
