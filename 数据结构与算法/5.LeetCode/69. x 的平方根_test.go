package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 二分， 小于等于num的最右
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	if x < 3 {
		return 1
	}
	ans := 1
	L := 1
	R := x
	for L <= R {
		M := (L + R) / 2
		if M*M <= x {
			ans = M
			L = M + 1
		} else {
			R = M - 1
		}
	}
	return ans
}
