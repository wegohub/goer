package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ans := longestPalindromeSubseq("cbbd")
	fmt.Println(ans)
	return "Hello World!", nil
}

func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	str := getManacherStr(s)
	pArr := make([]int, len(str))
	R := -1
	C := -1
	ans := math.MinInt
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

		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}

		// 再尝试往左右两边扩
		curMax := process(str, i-pArr[i], i+pArr[i], pArr[i])
		ans = int(math.Max(float64(ans), float64(curMax)))

	}
	return ans - 1
}

func getManacherStr(s string) string {
	sb := strings.Builder{}
	sb.WriteString("#")
	for _, item := range s {
		sb.WriteString(string(item) + "#")
	}
	return sb.String()
}

// 在str中从L、R位置往左右两边扩，返回左右匹配的最大长度
func process(str string, L, R, cur int) int {
	if L < 0 || R >= len(str) {
		return cur
	}

	if str[L] == str[R] {
		return process(str, L-1, R+1, cur+1)
	}

	p1 := process(str, L-1, R, cur)

	p2 := process(str, L, R+1, cur)

	return int(math.Max(float64(p1), float64(p2)))
}
