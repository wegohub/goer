package class03

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	arr := []int{4, 2, 1, 5, 6, 3}
	MergeSortIteration(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

func MergeSortIteration(arr []int) {
	if len(arr) < 2 {
		return
	}
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		L := 0
		for L < N {
			M := L + mergeSize - 1 // 上中点
			if M >= N-1 {          // 左组都凑不齐
				break
			}
			R := int(math.Min(float64(N-1), float64(M+mergeSize))) // 右组有可能凑不齐 (M + 1 + mergeSize - 1)
			//R := M + int(math.Min(float64(mergeSize), float64(N-1-M)))
			merge(arr, L, M, R) // 合并两组
			L = R + 1
		}
		if mergeSize > N/2 { // 防止溢出
			break
		}
		mergeSize = mergeSize << 1
	}
}

func merge(arr []int, L, M, R int) {
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
