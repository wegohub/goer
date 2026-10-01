package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 荷兰国旗问题，等于1的放中间
func sortColors(nums []int) {
	less := -1
	more := len(nums)
	index := 0
	for index < more {
		if nums[index] < 1 {
			nums[index], nums[less+1] = nums[less+1], nums[index]
			index++
			less++
		} else if nums[index] > 1 {
			nums[index], nums[more-1] = nums[more-1], nums[index]
			more--
		} else {
			index++
		}
	}
}
