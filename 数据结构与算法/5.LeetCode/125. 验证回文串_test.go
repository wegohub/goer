package leetcode

import (
	"fmt"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(filter("0p"))
	fmt.Println(string(byte(32)))
	return "Hello World!", nil
}

func isPalindrome(s string) bool {
	str := filter(s)
	mid := len(str) / 2
	L := 0
	R := len(str) - 1
	for L < mid && R > mid {
		if str[L] != str[R] {
			return false
		}
		L++
		R--
	}
	return true
}

func filter(s string) string {
	sb := strings.Builder{}
	sb.WriteString("#")
	for _, item := range s {
		if (item >= 'a' && item <= 'z') || (item >= '0' && item <= '9') {
			sb.WriteString(string(item) + "#")
		}
		if item >= 'A' && item <= 'Z' {
			sb.WriteString(string(item+32) + "#")
		}
	}
	return sb.String()
}
