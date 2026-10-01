package class03

import (
	"backend/utils/algo"
	"fmt"
	"math/rand"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := algo.GenArray(100, 1000)
	sort.Ints(arr)
	target := -1
	if len(arr) > 0 {
		index := int(rand.Float64() * float64(len(arr)))
		target = arr[index]
	}
	fmt.Println("arr: ", arr, ", target: ", target)
	fmt.Println(findNum(arr, target))
	return "Hello World!", nil
}

func findNum(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	L := 0
	R := n - 1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return -1
}
