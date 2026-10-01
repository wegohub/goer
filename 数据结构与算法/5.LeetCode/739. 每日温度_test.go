package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func dailyTemperatures(arr []int) []int {
	if arr == nil || len(arr) == 0 {
		return []int{}
	}
	N := len(arr)
	ans := make([]int, N)
	stack := [][]int{}

	for i := 0; i < N; i++ {
		for len(stack) > 0 && arr[stack[len(stack)-1][0]] < arr[i] {
			popIs := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, popi := range popIs {
				ans[popi] = i - popi
			}
		}
		if len(stack) > 0 && arr[stack[len(stack)-1][0]] == arr[i] {
			stack[len(stack)-1] = append(stack[len(stack)-1], i)
		} else {
			list := []int{i}
			stack = append(stack, list)
		}
	}
	return ans
}
