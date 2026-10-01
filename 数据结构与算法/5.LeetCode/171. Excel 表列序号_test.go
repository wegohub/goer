package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(numberToTitle(128))
	return "Hello World!", nil
}

// 伪进制
// 从列名称到列序号
func titleToNumber(columnTitle string) int {
	ans := 0
	for i := 0; i < len(columnTitle); i++ {
		ans = ans*26 + int(columnTitle[i]-'A'+1)
	}
	return ans
}

// 从列序号到列名称
func numberToTitle(columnNumber int) string {
	result := ""
	for columnNumber > 0 {
		columnNumber-- // 调整以适应 'A' 代表 1 的情况
		result = string(rune('A'+columnNumber%26)) + result
		columnNumber /= 26
	}
	return result
}
