package class04

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ans := rotateString("aa", "a")
	fmt.Println(ans)
	return "Hello World!", nil
}

func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	ss := s + s
	ans := IndexOf(ss, goal)
	return ans != -1
}

func IndexOf(str string, match string) int {
	if len(str) == 0 || len(match) == 0 || len(str) < len(match) {
		return -1
	}

	x := 0
	y := 0
	next := getNextArr(match)

	for x < len(str) && y < len(match) {
		if str[x] == match[y] {
			x++
			y++
		} else if next[y] == -1 {
			x++
		} else {
			y = next[y]
		}
	}

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
	i := 2
	cn := 0 // i - 1
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
