package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func canJump(nums []int) bool {
	if len(nums) < 2 {
		return true
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 小加速
		if max > len(nums)-1 {
			return true
		}
		if i > max {
			return false
		}
		max = int(math.Max(float64(max), float64(i+nums[i])))
	}
	return true
}
