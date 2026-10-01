package leetcode

import (
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func decodeString(s string) string {
	return process(s, 0).Ans
}

type Info struct {
	Ans  string
	Stop int
}

// s[i....]  何时停？遇到   ']'  或者遇到 s的终止位置，停止
// 返回Info
// 0) 串
// 1) 算到了哪
func process(s string, i int) *Info {
	ans := strings.Builder{}
	count := 0
	for i < len(s) && s[i] != ']' { // 字符
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') {
			ans.WriteByte(s[i])
			i++
		} else if s[i] >= '0' && s[i] <= '9' { // 数字
			count = count*10 + int(s[i]-'0')
			i++
		} else { // 遇到 [
			next := process(s, i+1)
			ans.WriteString(timeString(count, next.Ans))
			count = 0
			i = next.Stop + 1
		}
	}
	return &Info{Ans: ans.String(), Stop: i}
}

func timeString(times int, str string) string {
	ans := strings.Builder{}
	for i := 0; i < times; i++ {
		ans.WriteString(str)
	}
	return ans.String()
}
