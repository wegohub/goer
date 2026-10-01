package class03

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	arr := []int{3, 2, 3, 2, 3, 2, 3}
	fmt.Println(localMinimum(arr))

	return "Hello World!", nil
}

func localMinimum(arr []int) int {
	n := len(arr)
	// 没有数
	if n == 0 {
		return -1
	}
	// 有一个数
	if n == 1 {
		return arr[0]
	}
	// 有两个数
	if arr[0] < arr[1] {
		return 0
	}
	if arr[n-1] < arr[n-2] {
		return n - 1
	}
	// >2个数
	L := 0
	R := n - 1
	for L < R-1 { // L ... R-1 R 保证三个数及三个数以上，不出现越界
		mid := L + (R-L)/2
		if arr[mid] < arr[mid-1] && arr[mid] < arr[mid+1] {
			return mid
		} else if arr[mid] > arr[mid-1] {
			R = mid - 1
		} else {
			L = mid + 1
		}
	}

	// 迭代完成后要么剩一个数，要么剩两个数，谁小返回谁
	ans := R
	if arr[L] < arr[R] {
		ans = L
	}
	return ans
}
