package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxSubArray(nums []int) int {
	cur := 0
	max := math.MinInt
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		max = int(math.Max(float64(max), float64(cur)))
		if cur < 0 {
			cur = 0
		}
	}

	return max
}

func maxSubArray2(nums []int) int {
	cur := nums[0]
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//                 只包含i             包含i-1和i
		cur = int(math.Max(float64(nums[i]), float64(nums[i]+cur)))
		max = int(math.Max(float64(max), float64(cur)))
	}
	return max
}
