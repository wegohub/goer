package class01

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{5, 21, 4, 6, 1, 3, 0}
	bubbleSort(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

func bubbleSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// 0 ~ n-1推一个最大的
	// 0 ~ n-2推一个最大的
	// ...
	// 0 ～ 1推一个最大的
	for i := n - 1; i >= 0; i-- {
		for j := 1; j <= i; j++ {
			if nums[j-1] > nums[j] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}
