package class11

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	permutation("aaa")
	return "Hello World!", nil
}

func permutation(str string) {
	ans := make([]string, 0)
	process([]byte(str), 0, &ans)
	fmt.Println(ans)
}

func process(arr []byte, index int, ans *[]string) {
	if index == len(arr) {
		*ans = append(*ans, string(arr))
		return
	}
	for j := index; j < len(arr); j++ {
		// 交换index 和 j位置的字符
		arr[j], arr[index] = arr[index], arr[j]
		process(arr, index+1, ans)
		// 恢复现场
		arr[j], arr[index] = arr[index], arr[j]
	}
}
