package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 在i点的时候一定要让要让两次交易做完，而且最后一次交易的卖出时机在i位置
func maxProfit(prices []int) int {
	ans := 0
	// 交易一次并减去买入价格的最优
	doneOneMinusBuyMax := -prices[0] // 第一个位置没法卖，只能买
	// 做完一次交易的最大值
	doneOneMax := 0
	// 价格最小值
	min := prices[0]
	for i := 1; i < len(prices); i++ {
		// 单前做完两次交易的总收益
		ans = int(math.Max(float64(ans), float64(doneOneMinusBuyMax+prices[i])))
		// 更新价格最小的实际
		min = int(math.Min(float64(min), float64(prices[i])))
		// 做完一次交易的最大值(股票1的解)
		doneOneMax = int(math.Max(float64(doneOneMax), float64(prices[i]-min)))
		// 当前在i位置做决定
		// 不在当前的i买入: 结果是i-1位置的答案 doneOneMinusBuyMax
		// 在当前i买入: doneOneMax-prices[i]
		doneOneMinusBuyMax = int(math.Max(float64(doneOneMinusBuyMax), float64(doneOneMax-prices[i])))
	}

	return ans
}
