package class08

import (
	"fmt"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{7, 3, 5, 8, 5, 2, 1, 5, 7, 787, 45}
	QuickSort(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

func QuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	process(arr, 0, len(arr)-1)
}

func process(arr []int, L, R int) {
	if L > R {
		return
	}
	// L - R 上随机选一个
	i := L + int(rand.Float64()*float64(R-L+1))
	arr[i], arr[R] = arr[R], arr[i]
	res := partition(arr, L, R)
	process(arr, L, res[0]-1)
	process(arr, res[1]+1, R)
}

func partition(arr []int, L, R int) [2]int {
	base := arr[R]
	less := L - 1 // 小于区
	more := R     // 大于区
	index := L
	for index < more {
		if arr[index] < base {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			less++
			index++
		} else if arr[index] == base {
			index++
		} else {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		}
	}
	arr[R], arr[more] = arr[more], arr[R]

	return [2]int{less + 1, more}

}
