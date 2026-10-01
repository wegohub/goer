package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 滑动窗口+欠账本+总账目
func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	if len(s) == 0 || len(p) == 0 || len(s) < len(p) {
		return ans
	}
	N := len(s)
	M := len(p)

	// 账本
	mp := make(map[byte]int)
	for _, item := range p {
		if v, ok := mp[byte(item)]; ok {
			mp[byte(item)] = v + 1
		} else {
			mp[byte(item)] = 1
		}
	}

	// 总账目
	all := M

	// 初具窗口规模
	for end := 0; end < M-1; end++ {
		if v, ok := mp[s[end]]; ok {
			if v > 0 {
				all--
			}
			mp[s[end]] = v - 1
		}
	}

	// 滑动窗口
	end := M - 1
	start := 0
	for end < N {
		// 入窗口
		if v, ok := mp[s[end]]; ok {
			if v > 0 {
				all--
			}
			mp[s[end]] = v - 1
		}
		// 收集答案
		if all == 0 {
			ans = append(ans, start)
		}
		// 出窗口
		if v, ok := mp[s[start]]; ok {
			if v >= 0 {
				all++
			}
			mp[s[start]] = v + 1
		}
		start++
		end++
	}

	return ans
}
