package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 进位只可能是1
func plusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	ans := make([]int, n+1)
	ans[0] = 1
	// [1, n] 默认全是0
	// ans = append(ans, digits...)
	return ans
}
