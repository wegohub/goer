package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 抓所有上坡
func maxProfit(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += Max(prices[i]-prices[i-1], 0)
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
