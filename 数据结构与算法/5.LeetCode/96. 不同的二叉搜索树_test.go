package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func numTrees(N int) int {
	if N < 0 {
		return 0
	}
	if N < 2 {
		return 1
	}
	var a, b int64 = 1, 1
	for i, j := 1, N+1; i <= N; i, j = i+1, j+1 {
		a *= int64(i)
		b *= int64(j)
		gcd := gcd(a, b)
		a /= gcd
		b /= gcd
	}
	return int(b / a / int64(N+1))
}

func gcd(m, n int64) int64 {
	for n != 0 {
		m, n = n, m%n
	}
	return m
}

func numTrees2(n int) int {
	C := 1
	for i := 0; i < n; i++ {
		C = C * 2 * (2*i + 1) / (i + 2)
	}
	return C
}
