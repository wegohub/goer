package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(countAndSay(4))
	return "Hello World!", nil
}

func countAndSay(n int) string {
	if n < 1 {
		return ""
	}
	if n == 1 {
		return "1"
	}
	last := countAndSay(n - 1)
	ans := strings.Builder{}
	times := 1
	for i := 1; i < len(last); i++ {
		if last[i-1] == last[i] {
			times++
		} else {
			ans.WriteString(strconv.Itoa(times))
			ans.WriteByte(last[i-1])
			times = 1
		}
	}
	// 最后一个字符 返回：11
	ans.WriteString(strconv.Itoa(times))
	ans.WriteByte(last[len(last)-1])

	return ans.String()
}
