package class05

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ans := shortestPalindrome("aabababababaababaa")
	ans1 := shortestPalindromeKMP("aabababababaababaa")
	fmt.Println(ans, "==", ans1)

	return "Hello World!", nil
}

func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	// 先将s逆序, 并处理成manacher串
	sb := strings.Builder{}
	sb.WriteString("#")
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteString(string(s[i]) + "#")
	}
	str := sb.String()

	// 回文半径数组
	pArr := make([]int, len(str))
	// 回文半径右边界
	R := -1
	// 回文半径右边界中心点
	C := -1

	// 左边需要加到右边的字符范围[0, index]
	index := -1

	for i := 0; i < len(str); i++ {

		pArr[i] = 1
		if i < R {
			pArr[i] = int(math.Min(float64(pArr[2*C-i]), float64(R-i)))
		}

		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}

		// 更新回文半径
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}

		// 收集答案
		if R == len(str) {
			index = i - pArr[i]
			break
		}
	}

	// 添加字符
	for ; index >= 0; index-- {
		str = str + string(str[index])
	}

	// 去掉辅助字符
	ans := strings.Builder{}
	for _, item := range str {
		if item != '#' {
			ans.WriteByte(byte(item))
		}
	}

	return ans.String()

}

// 思路
// 1. s 做match串
// 2. s串的逆序串 做str原串
// 3. 执行KMP算法原型流程
// 4. 抓原串来到len(str)的时刻，match串匹配到那个位置index
// 5. 将[index, len(match)-1]的字符逆序加到s开头
func shortestPalindromeKMP(s string) string {
	if len(s) < 2 {
		return s
	}
	match := s
	sb := strings.Builder{}
	for index := len(s) - 1; index >= 0; index-- {
		sb.WriteString(string(s[index]))
	}
	str := sb.String()

	next := getNextArr(match)
	x := 0
	y := 0

	matchI := 0
	for x < len(str) && y < len(match) {
		if str[x] == match[y] {
			x++
			y++
		} else if y > 0 {
			y = next[y]
		} else {
			x++
		}

		// 匹配玩str中的最后一个字符, 得到此时match串中的位置
		if x == len(str) {
			matchI = y
		}
	}

	// 将[matchI, len(match)-1]的字符逆序加到s开头
	ans := strings.Builder{}
	for i := len(match) - 1; i >= matchI; i-- {
		ans.WriteString(string(match[i]))
	}
	return ans.String() + s
}

func getNextArr(match string) []int {
	if len(match) == 1 {
		return []int{-1}
	}
	next := make([]int, len(match))
	next[0] = -1
	next[1] = 0
	cn := 0
	i := 2
	for i < len(next) {
		if match[i-1] == match[cn] {
			next[i] = cn + 1
			i++
			cn++
		} else if cn > 0 {
			cn = next[cn]
		} else {
			next[i] = 0
			i++
		}
	}

	return next
}
