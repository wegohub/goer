package class02

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{6, 4, 0, 3, 2, 60}
	fmt.Println(getMaxVal(arr))
	return "Hello World!", nil
}

// 递归方法获取数组中的最大值
func getMaxVal(arr []int) int {
	return process(arr, 0)
}

func process(arr []int, index int) int {
	if index == len(arr)-1 {
		return arr[index]
	}
	return Max(arr[index], process(arr, index+1))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
