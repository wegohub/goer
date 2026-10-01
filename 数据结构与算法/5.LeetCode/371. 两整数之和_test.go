package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func getSum(a int, b int) int {
	sum := a
	for b != 0 {
		sum = a ^ b
		b = (a & b) << 1
		a = sum
	}
	return sum
}
