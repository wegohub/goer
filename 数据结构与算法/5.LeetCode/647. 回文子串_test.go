package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	ans := countSubstrings("aaa")
	fmt.Println(ans)

	return "Hello World!", nil
}

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}

	str := getManacherStr(s)
	pArr := make([]int, len(str))
	R := -1
	C := -1
	ans := 0

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

		// 收集答案

		ans += pArr[i] / 2
	}

	return ans
}

func getManacherStr(s string) string {
	sb := strings.Builder{}
	sb.WriteString("#")
	for _, item := range s {
		sb.WriteString(string(item) + "#")
	}
	return sb.String()
}
