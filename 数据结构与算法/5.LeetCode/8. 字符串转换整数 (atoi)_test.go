package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(myAtoi("\" \""))
	return "Hello World!", nil
}

func myAtoi(s string) int {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		return 0
	}
	// 过滤
	str := filter(s)
	if len(str) == 0 {
		return 0
	}
	// 验证
	if !validate(str) {
		return 0
	}

	// 是否负数
	isNeg := false
	if str[0] == '-' {
		isNeg = true
	}
	// 跳过开头的符号
	strart := 0
	if str[0] == '-' || str[0] == '+' {
		strart = 1
	}

	var (
		m   int32 = math.MinInt32 / 10
		o   int32 = math.MinInt32 % 10
		ans int32 = 0
	)
	for i := strart; i < len(str); i++ {
		// 处理成负数
		cur := int32('0') - int32(str[i])
		// 溢出判断
		if ans < m || (ans == m && cur < o) {
			if isNeg {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		// 更新结果
		ans = ans*10 + cur
	}

	// 处理正数
	if !isNeg {
		if ans == math.MinInt32 {
			return math.MaxInt32
		}
		ans = -ans
	}
	return int(ans)
}

// 过滤
func filter(str string) string {
	// 字符串前后的加减号
	flag := ""
	// 从左往右遍历的开始位置
	offset := 0
	if str[0] == '-' || str[0] == '+' {
		flag = string(str[0])
		offset = 1
	}

	// start到了第一个不是'0'字符的位置
	start := offset
	for ; start < len(str); start++ {
		if str[start] != '0' {
			break
		}
	}

	// end到了最左的, 不是数字字符的位置
	end := -1
	for i := len(str) - 1; i >= offset; i-- {
		if str[i] < '0' || str[i] > '9' {
			end = i
		}
	}
	if end == -1 {
		end = len(str)
	}
	return flag + str[start:end]
}

// 验证
func validate(s string) bool {
	if s[0] != '-' && s[0] != '+' && (s[0] < '0' || s[0] > '9') {
		return false
	}

	if (s[0] == '-' || s[0] == '+') && len(s) == 1 {
		return false
	}

	for i := 1; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}

	return true
}
