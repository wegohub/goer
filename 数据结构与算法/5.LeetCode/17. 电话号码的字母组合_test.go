package leetcode

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(letterCombinations(""))
	return "Hello World!", nil
}

var phone = [][]string{
	{"a", "b", "c"},      // 2    0
	{"d", "e", "f"},      // 3    1
	{"g", "h", "i"},      // 4    2
	{"j", "k", "l"},      // 5    3
	{"m", "n", "o"},      // 6
	{"p", "q", "r", "s"}, // 7
	{"t", "u", "v"},      // 8
	{"w", "x", "y", "z"}, // 9
}

func letterCombinations(digits string) []string {
	ans := list.New()
	if len(digits) > 0 {
		process(digits, 0, "", ans)
	}

	res := make([]string, 0, ans.Len())
	for ans.Len() > 0 {
		v := ans.Front().Value.(string)
		ans.Remove(ans.Front())
		res = append(res, v)
	}

	return res
}

func process(digits string, index int, path string, ans *list.List) {
	if index == len(digits) {
		ans.PushBack(path)
	} else {
		cur := phone[digits[index]-'2']
		for _, item := range cur {
			process(digits, index+1, path+string(item), ans)
		}
	}

}
