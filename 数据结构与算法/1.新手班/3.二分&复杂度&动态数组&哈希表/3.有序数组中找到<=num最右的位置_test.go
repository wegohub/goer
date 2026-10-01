package class03

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{1, 2, 4, 4, 4, 4, 4, 4, 4, 5, 6}
	fmt.Println(findMostRigntNum(arr, 4))
	return "Hello World!", nil
}

func findMostRigntNum(arr []int, num int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	L := 0
	R := n - 1
	ans := -1
	for L <= R {
		mid := L + (R-L)/2
		if arr[mid] <= num {
			ans = mid
			L = mid + 1
		} else {
			R = mid - 1
		}
	}

	return ans
}
