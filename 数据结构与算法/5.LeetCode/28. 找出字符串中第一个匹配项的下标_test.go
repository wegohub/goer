package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(strStr("mississippi", "issip"))
	return "Hello World!", nil
}

func strStr(haystack string, needle string) int {
	str := haystack
	match := needle

	x := 0
	y := 0
	next := getNextArr(match)
	fmt.Println(next)
	for x < len(str) && y < len(match) {
		if str[x] == match[y] {
			x++
			y++
		} else if y > 0 {
			y = next[y]
		} else {
			x++
		}
	}

	// match串越界了
	if y == len(match) {
		return x - y
	}

	return -1
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
	for i < len(match) {
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
