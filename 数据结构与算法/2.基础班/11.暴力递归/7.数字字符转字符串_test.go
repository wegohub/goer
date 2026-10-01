package class11

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	s := "111"
	fmt.Println(toStringNum(s))
	return "Hello World!", nil
}

func toStringNum(str string) int {
	return process([]byte(str), 0)
}

func process(arr []byte, i int) int {
	if i == len(arr) {
		return 1
	}

	if arr[i] == '0' {
		return 0
	}

	if arr[i] == '1' {
		ans := process(arr, i+1)
		if i+1 < len(arr) {
			ans += process(arr, i+2)
		}
		return ans
	}

	if arr[i] == '2' {
		ans := process(arr, i+1)
		if i+1 < len(arr) && arr[i+1] >= '0' && arr[i+1] <= '6' {
			ans += process(arr, i+2)
		}
		return ans
	}

	return process(arr, i+1)
}
