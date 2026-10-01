package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	pre2 := nums[0]
	pre1 := Max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cur := Max(nums[i]+pre2, pre1)
		pre2 = pre1
		pre1 = cur
	}
	return pre1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
