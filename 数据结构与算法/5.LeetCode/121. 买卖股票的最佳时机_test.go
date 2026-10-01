package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 0 - i上的最小值
// 每一步抓一个答案
func maxProfit(prices []int) int {
	min := prices[0]
	ans := 0
	for i := 0; i < len(prices); i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		ans = int(math.Max(float64(ans), float64(prices[i]-min)))
	}
	return ans
}
