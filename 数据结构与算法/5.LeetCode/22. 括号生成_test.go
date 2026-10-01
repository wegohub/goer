package leetcode

import (
	"container/list"
	"fmt"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(generateParenthesis(3))
	return "Hello World!", nil
}

// 最优解
func generateParenthesis(n int) []string {
	// 单个字符的长度 2*n
	path := make([]string, 2*n)
	ans := make([]string, 0)
	process2(path, 0, 0, n, &ans)
	return ans
}

func process2(path []string, index int, leftMinusRight int, leftRest int, ans *[]string) {
	if index == len(path) {
		*ans = append(*ans, strings.Join(path, ""))
	} else {
		// 做剪枝
		if leftRest > 0 { // 左括号的数量
			path[index] = "("
			process2(path, index+1, leftMinusRight+1, leftRest-1, ans)
		}

		if leftMinusRight > 0 { // 右括号的数量
			path[index] = ")"
			process2(path, index+1, leftMinusRight-1, leftRest, ans)
		}
	}
}

func generateParenthesis2(n int) []string {
	// 单个字符的长度 2*n
	path := make([]string, 2*n)
	ans := list.New()

	process(path, 0, 0, n, ans)

	res := make([]string, 0, ans.Len())
	for ans.Len() > 0 {
		cur := ans.Front().Value.(string)
		ans.Remove(ans.Front())
		res = append(res, cur)
	}

	return res
}

func process(path []string, index int, leftMinusRight int, leftRest int, ans *list.List) {
	if index == len(path) {
		ans.PushBack(strings.Join(path, ""))
	} else {
		if leftRest > 0 { // 左括号的数量
			path[index] = "("
			process(path, index+1, leftMinusRight+1, leftRest-1, ans)
		}

		if leftMinusRight > 0 { // 右括号的数量
			path[index] = ")"
			process(path, index+1, leftMinusRight-1, leftRest, ans)
		}
	}
}

// 不剪枝版本(暴力解)
func generateParenthesis3(n int) []string {
	// 单个字符的长度 2*n
	path := make([]string, 2*n)
	ans := make([]string, 0)

	process3(path, 0, &ans)

	return ans
}

func process3(path []string, index int, ans *[]string) {
	if index == len(path) {
		if isValidate(path) {
			*ans = append(*ans, strings.Join(path, ""))
		}
	} else {
		path[index] = "("
		process3(path, index+1, ans)
		path[index] = ")"
		process3(path, index+1, ans)
	}
}

// 检查有效性
func isValidate(path []string) bool {
	count := 0
	for _, item := range path {
		if item == "(" {
			count++
		} else {
			count--
		}
		if count < 0 {
			return false
		}
	}

	return count == 0
}
