package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// k >=  n/2 就是问题2(无限次交易)
// dp定义：在[0,i]，不超过j次交易，最大收益是多少？
func maxProfit(k int, prices []int) int {
	N := len(prices)
	if k >= N/2 {
		return allTrans(prices)
	}
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, k+1)
	}
	for i := 1; i < N; i++ {
		for j := 1; j <= k; j++ {
			// 1. [i]不参与交易
			dp[i][j] = dp[i-1][j]
			// 2. 贪心策略，i只参与最后一次卖出，枚举所有的买入时机，能否推高dp[i][j]
			// 当前位置是p, 那么在[0,p-1]是做不超过k-1次交易的最好收益 + (当前i卖出的价格 - p位置买入的价格)
			// 在p位置买入，在i位置卖出
			for p := 0; p <= i; p++ {
				dp[i][j] = int(math.Max(float64(dp[p][j-1]+prices[i]-prices[p]), float64(dp[i][j])))
			}
		}
	}
	return dp[N-1][k]
}

func allTrans(prices []int) int {
	ans := 0
	for i := 1; i < len(prices); i++ {
		ans += int(math.Max(float64(prices[i]-prices[i-1]), 0))
	}
	return ans
}

// 最优解
func maxProfit2(k int, prices []int) int {
	N := len(prices)
	if k >= N/2 {
		return allTrans(prices)
	}
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, N)
	}
	ans := 0
	for j := 1; j <= k; j++ {
		pre := dp[j][0]
		best := pre - prices[0]
		for i := 1; i < N; i++ {
			pre = dp[j-1][i]
			dp[j][i] = int(math.Max(float64(dp[j][i-1]), float64(prices[i]+best)))
			best = int(math.Max(float64(best), float64(pre-prices[i])))
			ans = int(math.Max(float64(dp[j][i]), float64(ans)))
		}
	}
	return ans
}
