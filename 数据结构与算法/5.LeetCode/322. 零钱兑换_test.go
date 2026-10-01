package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount < 0 {
		return -1
	}

	N := len(coins)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, amount+1)
	}

	// dp[i][0] = 0 0列不需要填
	// dp[0][1...] = arr[0]的整数倍，有张数，倍数，其他的格子-1（表示无方案）
	for j := 1; j <= amount; j++ {
		if j%coins[0] != 0 {
			dp[0][j] = -1
		} else {
			dp[0][j] = j / coins[0]
		}
	}

	for i := 1; i < N; i++ {
		for j := 1; j <= amount; j++ {
			dp[i][j] = math.MaxInt
			if dp[i-1][j] != -1 {
				dp[i][j] = dp[i-1][j]
			}

			if j-coins[i] >= 0 && dp[i][j-coins[i]] != -1 {
				dp[i][j] = int(math.Min(float64(dp[i][j]), float64(dp[i][j-coins[i]]+1)))
			}

			if dp[i][j] == math.MaxInt {
				dp[i][j] = -1
			}
		}
	}

	return dp[N-1][amount]
}
