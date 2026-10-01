package leetcode

import "strings"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func wordBreak(s string, wordDict []string) []string {
	set := make(map[string]struct{})
	for _, item := range wordDict {
		set[item] = struct{}{}
	}
	path := make([]string, 0)
	ans := make([]string, 0)
	process(s, set, 0, path, &ans)
	return ans
}

func process(s string, wordMap map[string]struct{}, index int, path []string, ans *[]string) {
	if index == len(s) {
		*ans = append(*ans, strings.Join(path, " "))
	}
	for end := index; end < len(s); end++ {
		cur := s[index : end+1]
		if _, ok := wordMap[cur]; ok {
			path = append(path, cur)
			process(s, wordMap, end+1, path, ans)
			path = path[0 : len(path)-1]
		}
	}
}
