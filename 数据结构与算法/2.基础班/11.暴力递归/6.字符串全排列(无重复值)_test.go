package class11

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	permutationNoRepet("aaa")
	return "Hello World!", nil
}

func permutationNoRepet(str string) {
	ans := make([]string, 0)
	process([]byte(str), 0, &ans)
	fmt.Println(ans)
}

func process(arr []byte, index int, ans *[]string) {
	if index == len(arr) {
		*ans = append(*ans, string(arr))
		return
	}
	// 开始是0位置index
	visit := make(map[byte]struct{})
	// 0位置的a来到index, 1位置的a来到index,这是已经尝试过0位置的a的，跳过
	for j := index; j < len(arr); j++ {
		if _, ok := visit[arr[j]]; !ok {
			visit[arr[j]] = struct{}{}
			arr[j], arr[index] = arr[index], arr[j]
			process(arr, index+1, ans)
			arr[j], arr[index] = arr[index], arr[j]
		}
	}
}
