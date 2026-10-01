package class03

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	nums := []int{1, 2, 3, 4, 5}
	nums = InsertElem(nums, 2, 6)
	nums = InsertElem(nums, 2, 6)
	nums = InsertElem(nums, 2, 6)
	nums = InsertElem(nums, 2, 6)
	fmt.Println(nums)
	return "Hello World!", nil
}

// 实现插入功能
func InsertElem(arr []int, index int, value int) []int {
	return append(arr[:index], append([]int{value}, arr[index:]...)...)
}

// 实现数组copy
func CopyElem(nums []int) []int {
	help := make([]int, len(nums))
	copy(nums[0:], help)
	return help
}
