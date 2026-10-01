package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 滑动窗口+记账表
func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if len(s) == 0 || k < 1 {
		return 0
	}

	N := len(s)
	// 词频统计表
	count := make([]int, 256)
	// 字符的种类
	diff := 0
	R := 0
	ans := 0

	for L := 0; L < len(s); L++ {
		// R 窗口的右边界           当前位置的字符收集过
		for R < N && (diff < k || count[s[R]] > 0) {
			if count[s[R]] == 0 {
				diff += 1
			}
			count[s[R]]++
			R++
		}
		// R 来到违规的第一个位置
		ans = Max(ans, R-L)
		// i位置的字符只有一次了
		if count[s[L]] == 1 {
			diff -= 1
		}
		count[s[L]]--
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLongestSubstringKDistinct2(s string, k int) int {
	N := len(s)
	// 词频表
	mp := make(map[rune]int)
	all := 0
	R := 0
	ans := 0
	for L := 0; L < N; L++ {
		// R来到第一个不达标的位置
		for R < N && (all < k || mp[rune(s[R])] > 0) {
			if mp[rune(s[R])] == 0 {
				all++
			}
			mp[rune(s[R])]++
			R++
		}
		ans = Max(ans, R-L)

		mp[rune(s[L])]--
		if mp[rune(s[L])] == 0 {
			all--
		}
	}
	return ans
}
