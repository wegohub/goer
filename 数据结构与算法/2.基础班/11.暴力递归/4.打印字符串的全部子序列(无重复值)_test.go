package class11

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	SubsNoRepet("sss")
	return "Hello World!", nil
}

func SubsNoRepet(str string) {
	set := make(map[string]struct{})
	process([]byte(str), 0, "", set)

	for v := range set {
		fmt.Println(v)
	}

}

func process(arr []byte, index int, path string, ans map[string]struct{}) {
	if index == len(arr) {
		ans[path] = struct{}{}
		return
	}

	// 不要index字符
	process(arr, index+1, path, ans)

	// 要index字符
	process(arr, index+1, path+string(arr[index]), ans)
}
