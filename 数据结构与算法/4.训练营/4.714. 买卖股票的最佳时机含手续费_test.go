package train

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxProfit(prices []int, fee int) int {
	if len(prices) < 1 {
		return 0
	}
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = Max(dp[i-1][0], dp[i-1][1]+prices[i]-fee)
		dp[i][1] = Max(dp[i-1][1], dp[i-1][0]-prices[i])
	}

	return dp[n-1][0]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
