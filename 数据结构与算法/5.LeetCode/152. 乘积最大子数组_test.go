package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(maxProduct([]int{0, 10, 10, 10, 10, 10, 10, 10, 10, 10, -10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 0}))
	return "Hello World!", nil
}

// 动态规划解
func maxProduct(nums []int) int {
	// 答案
	ans := nums[0]
	// 最小累乘积
	min := nums[0]
	// 最大累乘积
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		// 防止溢出
		if min*nums[i] < math.MinInt32 || min*nums[i] > math.MaxInt32 || max*nums[i] < math.MinInt32 || max*nums[i] > math.MaxInt32 {
			return ans
		}
		// 以i位置结尾讨论：
		// 1. 只含i自己
		// 2. 不只含i自己：
		// 2.1 i-1的最大累乘积 * i
		// 2.2 i-1的最小累乘积 * i
		curMin := Min(nums[i], Min(min*nums[i], max*nums[i]))
		curMax := Max(nums[i], Max(min*nums[i], max*nums[i]))
		min = curMin
		max = curMax
		ans = Max(ans, max)
	}
	return ans
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
