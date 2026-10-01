package class03

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	arr := []int{1, 2, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	fmt.Println(findMostLeftNum(arr, 3))

	return "Hello World!", nil
}

func findMostLeftNum(arr []int, num int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}
	L := 0
	R := n - 1
	ans := -1
	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] >= num {
			ans = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return ans
}
