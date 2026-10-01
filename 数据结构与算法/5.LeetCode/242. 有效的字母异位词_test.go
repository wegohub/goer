package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 词频统计 时间复杂度O(N) 空间复杂度O(N)
// 排序对比：时间复杂度N*LogN 空间复杂度O(1)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := make(map[rune]int)
	for _, item := range s {
		mp[rune(item)] += 1
	}
	for _, item := range t {
		mp[rune(item)] -= 1
		if mp[rune(item)] < 0 {
			return false
		}
	}
	return true
}
