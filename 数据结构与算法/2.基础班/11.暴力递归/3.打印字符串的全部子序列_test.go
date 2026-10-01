package class11

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	Subs("abcd")
	return "Hello World!", nil
}

func Subs(str string) {
	ans := list.New()
	process([]byte(str), 0, "", ans)
	for ans.Len() > 0 {
		v := ans.Front().Value.(string)
		ans.Remove(ans.Front())
		if v == "" {
			fmt.Println("空")
		} else {
			fmt.Println(v)
		}
	}
}

func process(arr []byte, index int, path string, ans *list.List) {
	if index == len(arr) {
		ans.PushBack(path)
		return
	}

	// 不要index字符
	process(arr, index+1, path, ans)

	// 要index字符
	process(arr, index+1, path+string(arr[index]), ans)
}
