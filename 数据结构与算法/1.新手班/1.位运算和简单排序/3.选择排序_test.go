package class01

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{5, 21, 4, 6, 1, 3, 0}
	selectSort(arr)
	fmt.Println(arr)
	return "", nil
}

// 0 - n上选一个最小的
// 1 - n上选一个最小的
// 2 - n上选一个最小的
// ...
// n-1 - n上选一个最小的
func selectSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
