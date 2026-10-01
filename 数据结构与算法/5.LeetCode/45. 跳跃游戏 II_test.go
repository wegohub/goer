package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 当前最少跳几步能到i
	step := 0
	// 跳的步数不超过step，能到的最右位置
	cur := 0
	// 跳的步数不超过step+1，能到的最右位置
	next := nums[0]
	for i := 1; i < len(nums); i++ {
		// 小加速
		if next >= len(nums)-1 {
			return step + 1
		}
		if i > cur {
			step++
			cur = next
		}
		next = int(math.Max(float64(next), float64(nums[i]+i)))
	}
	return step
}
