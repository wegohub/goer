package class03

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr1 := []int{1, 3, 4, 2, 5}
	arr2 := []int{1, 3, 4, 2, 5}
	fmt.Println(smallSum(arr1))
	fmt.Println(smallSum2(arr2))
	return "Hello World!", nil
}

func smallSum(arr []int) int {
	ans := 0
	N := len(arr)
	mergeSize := 1
	for mergeSize < N {
		L := 0
		for L < N {
			M := L + mergeSize - 1
			if M >= N-1 { // == N 没有右组，无需合并
				break
			}
			R := int(math.Min(float64(N-1), float64(M+mergeSize)))
			ans += merge(arr, L, M, R)
			L = R + 1
		}

		if mergeSize > N/2 {
			break
		}

		mergeSize <<= 1
	}

	return ans
}

func smallSum2(arr []int) int {
	return process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) int {
	if L >= R {
		return 0
	}
	M := (L + R) / 2
	return process(arr, L, M) + process(arr, M+1, R) + merge(arr, L, M, R)
}

func merge(arr []int, L, M, R int) int {
	ans := 0
	help := make([]int, R-L+1)
	index := 0
	i := L
	j := M + 1

	for i <= M && j <= R {
		if arr[i] < arr[j] {
			help[index] = arr[i]
			// 产生小和
			ans += arr[i] * (R - j + 1)
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
	return ans
}
