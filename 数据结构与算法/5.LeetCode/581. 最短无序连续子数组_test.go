package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 从左往右，找到比左边最大值还小的最右下标，从右往左，找到比右边最小值还大的最左下标
func findUnsortedSubarray(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	N := len(nums)
	right := -1
	max := math.MinInt
	for i := 0; i < N; i++ {
		if max > nums[i] {
			right = i
		}
		max = int(math.Max(float64(max), float64(nums[i])))
	}
	min := math.MaxInt
	left := N
	for i := N - 1; i >= 0; i-- {
		if min < nums[i] {
			left = i
		}
		min = int(math.Min(float64(min), float64(nums[i])))
	}

	return int(math.Max(float64(0), float64(right-left+1)))
}
