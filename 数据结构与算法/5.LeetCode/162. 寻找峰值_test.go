package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func findPeakElement(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	length := len(arr)
	if length == 1 {
		return 0
	}
	if arr[0] > arr[1] {
		return 0
	}
	if arr[length-1] > arr[length-2] {
		return length - 1
	}
	// 二分
	left := 1
	right := length - 2
	for left < right {
		mid := (left + right) / 2
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		} else if arr[mid] < arr[mid-1] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
