package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(longestCommonPrefix([]string{""}))
	return "Hello World!", nil
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 拿出第一个字符串
	str := strs[0]
	min := math.MaxInt

	// 从第二个字符串开始比较
	for i := 0; i < len(strs); i++ {
		cur := strs[i]
		index := 0
		for index < len(str) && index < len(cur) {
			if str[index] != cur[index] {
				break
			}
			index++
		}

		min = int(math.Min(float64(min), float64(index)))
		if min == 0 {
			return ""
		}
	}

	return str[0:min]
}
