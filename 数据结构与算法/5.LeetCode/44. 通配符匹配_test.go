package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(isMatch2("cb", "?a"))
	return "Hello World!", nil
}

func isMatch(s string, p string) bool {
	return process(s, p, 0, 0)
}

func process(s string, p string, si, pi int) bool {
	// base case
	// si来到末尾
	// 1. pi也来到末尾
	// 2. pi没有到末尾，但是pi位置的字符是*，并且pi+1及以后的字符都能变成空字符串
	if si == len(s) {
		if pi == len(p) {
			return true
		}
		return p[pi] == '*' && process(s, p, si, pi+1)
	}

	// pi来到末尾，只有si也来到末尾才能匹配
	if pi == len(p) {
		return si == len(s)
	}

	// pi位置是字符，不是？也不是*，那么只能是s[si] == p[pi] && si+1 和 pi+1也能匹配
	if p[pi] != '?' && p[pi] != '*' {
		return s[si] == p[pi] && process(s, p, si+1, pi+1)
	}

	// 如果pi是？，可以匹配单个字符，si+1 和 pi+1也能匹配
	if p[pi] == '?' {
		return process(s, p, si+1, pi+1)
	}

	// pi位置是*，尝试每一种前缀
	for start := 0; start <= len(s)-si; start++ {
		// abbbcd
		// a*cd
		// 初始si==1
		// 尝试 * == b
		// * == bb
		// * == bbb
		if process(s, p, si+start, pi+1) {
			return true
		}
	}

	return false
}

func isMatch2(s string, p string) bool {
	N := len(s)
	M := len(p)
	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, M+1)
	}
	dp[N][M] = true
	// 填最后一行
	for pi := M - 1; pi >= 0; pi-- {
		dp[N][pi] = p[pi] == '*' && dp[N][pi+1]
	}

	for si := N - 1; si >= 0; si-- {
		for pi := M - 1; pi >= 0; pi-- {
			if p[pi] != '?' && p[pi] != '*' {
				dp[si][pi] = s[si] == p[pi] && dp[si+1][pi+1]
				continue
			}

			if p[pi] == '?' {
				dp[si][pi] = dp[si+1][pi+1]
				continue
			}

			// 尝试每一种前缀
			for start := 0; start <= len(s)-si; start++ {
				if dp[si+start][pi+1] {
					dp[si][pi] = true
					break
				}
			}

		}
	}

	return dp[0][0]
}

// 最优解
func isMatch3(s string, p string) bool {
	N := len(s)
	M := len(p)
	dp := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]bool, M+1)
	}
	dp[N][M] = true
	// 填最后一行
	for pi := M - 1; pi >= 0; pi-- {
		dp[N][pi] = p[pi] == '*' && dp[N][pi+1]
	}

	for si := N - 1; si >= 0; si-- {
		for pi := M - 1; pi >= 0; pi-- {
			if p[pi] != '?' && p[pi] != '*' {
				dp[si][pi] = s[si] == p[pi] && dp[si+1][pi+1]
				continue
			}

			if p[pi] == '?' {
				dp[si][pi] = dp[si+1][pi+1]
				continue
			}

			dp[si][pi] = dp[si][pi+1] || dp[si+1][pi]
		}
	}

	return dp[0][0]
}
