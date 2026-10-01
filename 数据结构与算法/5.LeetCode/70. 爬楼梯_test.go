package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	ans := 0
	n2 := 1
	n1 := 2
	for i := 3; i <= n; i++ {
		ans = n1 + n2
		n2 = n1
		n1 = ans
	}
	return ans
}
