package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func longestValidParentheses(str string) int {
	if str == "" {
		return 0
	}
	chas := []rune(str)
	dp := make([]int, len(chas))
	pre := 0
	res := 0
	for i := 1; i < len(chas); i++ {
		if chas[i] == ')' {
			// pre是，当前i位置的), 应该找哪个位置的左括号
			pre = i - dp[i-1] - 1
			if pre >= 0 && chas[pre] == '(' {
				dp[i] = dp[i-1] + 2 + func() int {
					if pre > 0 {
						return dp[pre-1]
					}
					return 0
				}()
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses1(s string) int {
	N := len(s)
	if N <= 1 {
		return 0
	}
	dp := make([]int, N)
	pre := 0
	maxLen := 0
	for i := 1; i < N; i++ {
		if s[i] == ')' {
			// 当前谁和i位置的)去配
			pre = i - dp[i-1] - 1
			// 越界了        能匹配
			if pre >= 0 && s[pre] == '(' {
				// [4]   (     [2]      )
				//  pre  pre    dp[i-1]  i
				dp[i] = dp[i-1] + 2
				if pre > 0 {
					dp[i] += dp[pre-1]
				}
			}
		}

		maxLen = Max(maxLen, dp[i])
	}

	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
