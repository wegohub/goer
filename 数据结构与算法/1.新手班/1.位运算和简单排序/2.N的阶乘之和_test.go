package class01

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return factorial(3), nil
}

func factorial(n int) int {
	if n < 0 {
		return 0
	}
	ans := 0
	pre := 1
	for i := 1; i <= n; i++ {
		pre = pre * i
		ans += pre
	}
	return ans
}
