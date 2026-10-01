package leetcode

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(fractionToDecimal(71, 3))
	return "Hello World!", nil
}

// 记余数出现的位置, 有理数必定是有穷尽的
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	sb := strings.Builder{}

	// 判断正负
	if (numerator > 0 && denominator < 0) || (numerator < 0 && denominator > 0) {
		sb.WriteString("-")
	}

	// 转为绝对值
	num := int(math.Abs(float64(numerator)))
	den := int(math.Abs(float64(denominator)))

	// 整数部分
	sb.WriteString(strconv.Itoa(num / den))

	// 小数部分
	num = num % den

	// 没有小数
	if num == 0 {
		return sb.String()
	}

	// 有小数
	sb.WriteString(".")

	// 用map记录余数出现的位置
	mp := make(map[int]int)
	mp[num] = sb.Len()

	start := -1
	for num != 0 {
		num = num * 10
		sb.WriteString(strconv.Itoa(num / den))
		num = num % den
		// 如果余数出现重复
		if index, ok := mp[num]; ok {
			start = index
			break
		} else {
			mp[num] = sb.Len()
		}
	}
	ans := sb.String()
	if start == -1 {
		return ans
	}
	return ans[0:start] + "(" + ans[start:] + ")"
}
