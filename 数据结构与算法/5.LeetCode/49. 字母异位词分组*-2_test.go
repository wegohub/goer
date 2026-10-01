package leetcode

import "sort"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func groupAnagrams(strs []string) [][]string {
	// 词频和: 词
	wordMap := make(map[string][]string)
	for _, str := range strs {
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)
		wordMap[sortedStr] = append(wordMap[sortedStr], str)
	}

	ans := make([][]string, 0, len(wordMap))
	for _, v := range wordMap {
		ans = append(ans, v)
	}

	return ans
}
