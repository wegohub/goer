package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func partitionLabels(s string) []int {
	// 每个字符能冲到的最右位置
	mp := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		mp[s[i]] = i
	}

	// 结算字符
	ans := make([]int, 0)
	moreRight := -1
	start := 0
	for i := 0; i < len(s); i++ {
		moreRight = Max(moreRight, mp[s[i]])
		if i == moreRight {
			ans = append(ans, i-start+1)
			start = i + 1
		}
	}
	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
