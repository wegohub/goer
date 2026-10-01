package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buy1 := Max(-prices[0], -prices[1])
	sell1 := Max(0, prices[1]-prices[0])
	sell2 := 0
	for i := 2; i < len(prices); i++ {
		tmp := sell1
		sell1 = Max(sell1, buy1+prices[i])
		buy1 = Max(buy1, sell2-prices[i])
		sell2 = tmp
	}
	return sell1
}

func maxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	N := len(prices)
	buy := make([]int, N)
	sell := make([]int, N)
	buy[1] = Max(-prices[0], -prices[1])
	sell[1] = Max(0, prices[1]-prices[0])
	for i := 2; i < N; i++ {
		buy[i] = Max(buy[i-1], sell[i-2]-prices[i])
		sell[i] = Max(sell[i-1], buy[i-1]+prices[i])
	}
	return sell[N-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
