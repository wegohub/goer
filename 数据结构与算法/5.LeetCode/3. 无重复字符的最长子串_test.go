package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	return "Hello World!", nil
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	// key: 字符 value: 上次出现的位置
	mp := make(map[byte]int)
	// i-1位置能往左推的最远距离
	pre := -1
	ans := math.MinInt
	for i := 0; i < len(s); i++ {
		// 当前字符能往左推的最长距离，默认-1
		if v, ok := mp[byte(s[i])]; ok {
			pre = int(math.Max(float64(pre), float64(v)))
		}
		// 记录位置
		mp[byte(s[i])] = i
		// 更新答案，当前能推的距离为i-(pre+1)+1
		ans = int(math.Max(float64(ans), float64(i-pre)))
	}
	return ans
}
