package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 贪心算法
// a.b >= b.a a放前
// a.b < b.a b放前
func largestNumber(nums []int) string {
	strs := make([]string, 0, len(nums))
	for _, num := range nums {
		strs = append(strs, strconv.Itoa(num))
	}
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] >= strs[j]+strs[i]
	})
	ans := strings.Join(strs, "")
	// 过滤掉0
	index := -1
	for i := 0; i < len(ans); i++ {
		if ans[i] != '0' {
			index = i
			break
		}
	}
	// 全是0
	if index == -1 {
		return "0"
	}
	return ans[index:]
}
