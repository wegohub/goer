package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func moveZeroes(nums []int) {
	if len(nums) < 2 {
		return
	}
	to := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[to] = nums[to], nums[i]
			to++
		}
	}
}
