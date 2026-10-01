package class12

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	s := "111"
	ans1 := ToStringNum(s)

	fmt.Println(ans1)
	return "Hello World!", nil
}

func ToStringNum(str string) int {
	return process([]byte(str), 0)
}

func process(arr []byte, i int) int {
	// 能来到最后位置，找到一种答案
	if i == len(arr) {
		return 1
	}
	// 单独来到0字符是无效解
	if arr[i] == '0' {
		return 0
	}

	if arr[i] == '1' {
		ans := process(arr, i+1)
		if i+1 < len(arr) {
			ans += process(arr, i+2)
		}
		return ans
	}

	if arr[i] == '2' {
		ans := process(arr, i+1)
		if i+1 < len(arr) && arr[i+1] >= '0' && arr[i+1] <= '6' {
			ans += process(arr, i+2)
		}
		return ans
	}

	return process(arr, i+1)
}

func numDecodingsDP(s string) int {
	N := len(s)
	dp := make([]int, N+1)
	dp[N] = 1
	for i := N - 1; i >= 0; i-- {
		if s[i] == '0' {
			dp[i] = 0
			continue
		}
		ans := dp[i+1]
		if s[i] == '1' {
			if i+1 < len(s) {
				ans += dp[i+2]
			}
		} else if s[i] == '2' {
			if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '6' {
				ans += dp[i+2]
			}
		}
		dp[i] = ans
	}
	return dp[0]
}

// 动态规划，空间优化
func numDecodingsDP1(s string) int {
	N := len(s)
	n1 := 1
	n2 := -1
	for i := N - 1; i >= 0; i-- {
		if s[i] == '0' {
			n2 = n1
			n1 = 0
			continue
		}
		ans := n1
		if s[i] == '1' {
			if i+1 < len(s) {
				ans += n2
			}
		} else if s[i] == '2' {
			if i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '6' {
				ans += n2
			}
		}
		n2 = n1
		n1 = ans
	}
	return n1
}
