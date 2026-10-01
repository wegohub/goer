package leetcode

import (
	"fmt"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	strs := []string{"ac", "d"}
	ans := groupAnagrams(strs)
	fmt.Println(ans)

	var a []int
	a = append(a, 123)
	fmt.Println(a)

	// map就必须要分配内存，因为append操作是反射实现的
	var mp map[string]string
	mp["a"] = "1"
	fmt.Println(mp)

	return "Hello World!", nil
}

// 排序
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

// 计数， [26]int{} 数组可以作为map的键
func groupAnagrams2(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
