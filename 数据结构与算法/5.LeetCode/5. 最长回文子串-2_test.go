package leetcode

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(longestPalindrome("babad"))
	return "Hello World!", nil
}

func longestPalindrome(s string) string {
	// manacher串
	str := getManacherStr(s)
	// 回文半径数组
	parr := make([]int, len(str))
	// 回文半径右边界
	R := -1
	// 回文半径右边界对应的中心点
	C := -1
	// 最大回文半径
	maxR := -1
	// 最大回文半径对于的下标
	maxI := -1
	for i := 0; i < len(str); i++ {
		parr[i] = 1
		if i < R {
			//                             这里老是出错，是求i得对称点的回文半径
			parr[i] = int(math.Min(float64(parr[2*C-i]), float64(R-i)))
		}

		for i+parr[i] < len(str) && i-parr[i] > -1 {
			if str[i+parr[i]] == str[i-parr[i]] {
				parr[i]++
			} else {
				break
			}
		}

		if i+parr[i] > R {
			R = i + parr[i]
			C = i
		}

		if parr[i] > maxR {
			maxR = parr[i]
			maxI = i
		}
	}
	// 回文半径访问 [i-parr[i]+1, i+parr[i]-1]
	ans := str[maxI-maxR+1 : maxI+maxR]

	// 去掉辅助字符#
	sb := strings.Builder{}
	for _, item := range ans {
		if string(item) != "#" {
			sb.WriteString(string(item))
		}
	}
	return sb.String()
}

// 获取manacher串
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
