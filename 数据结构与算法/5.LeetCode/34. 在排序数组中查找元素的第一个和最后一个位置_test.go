package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(searchRange([]int{}, 0))
	return "Hello World!", nil
}

func searchRange(nums []int, target int) []int {
	return []int{largeNumMostLeft(nums, target), lessNumMostRight(nums, target)}
}

// >= num 最左的位置
func largeNumMostLeft(arr []int, num int) int {
	L := 0
	R := len(arr) - 1

	ans := -1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] > num {
			R = mid - 1
		} else if arr[mid] < num {
			L = mid + 1
		} else {
			ans = mid
			R = mid - 1
		}

	}
	return ans
}

// <= num的最右位置
func lessNumMostRight(arr []int, num int) int {
	L := 0
	R := len(arr) - 1

	ans := -1
	for L <= R {
		mid := (L + R) / 2
		if arr[mid] < num {
			L = mid + 1
		} else if arr[mid] > num {
			R = mid - 1
		} else {
			ans = mid
			L = mid + 1
		}

	}
	return ans
}
