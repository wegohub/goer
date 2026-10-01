package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func firstUniqChar(s string) int {
	N := len(s)
	count := make([]int, 26)
	for i := 0; i < N; i++ {
		count[s[i]-'a']++
	}
	for i := 0; i < N; i++ {
		if count[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
