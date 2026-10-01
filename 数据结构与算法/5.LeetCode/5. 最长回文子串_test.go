package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	ans := longestPalindrome("cbbd")
	fmt.Println(ans)
	return "Hello World!", nil
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	str := getManacherStr(s)
	pArr := make([]int, len(str))
	R := -1
	C := -1
	maxIR := math.MinInt
	maxI := -1

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

		// 收集最大的回文半径
		if pArr[i] > maxIR {
			maxIR = pArr[i]
			maxI = i
		}
	}

	return manacherStrToStr(str, maxI-maxIR+1, maxI+maxIR-1)
}

func getManacherStr(s string) string {
	sb := strings.Builder{}
	sb.WriteString("#")
	for _, item := range s {
		sb.WriteString(string(item) + "#")
	}
	return sb.String()
}

// 在manacher串上从[L, R] 上转为普通字符
func manacherStrToStr(mStr string, L, R int) string {
	if L > R || L < 0 || R < 0 || L >= len(mStr) || R >= len(mStr) {
		return ""
	}
	sb := strings.Builder{}
	for i := L; i <= R; i++ {
		if string(mStr[i]) != "#" {
			sb.WriteString(string(mStr[i]))
		}
	}
	return sb.String()
}
