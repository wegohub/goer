package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}
	N := len(s)
	dp := make([]int, N+1)
	dp[N] = 1
	for i := N - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
		} else if s[i] == '1' {
			dp[i] = dp[i+1]
			if i+1 < N {
				dp[i] += dp[i+2]
			}
		} else if s[i] == '2' {
			dp[i] = dp[i+1]
			if i+1 < N && s[i+1] >= '0' && s[i+1] <= '6' {
				dp[i] += dp[i+2]
			}
		} else {
			dp[i] = dp[i+1]
		}
	}
	return dp[0]
}
