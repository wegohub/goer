package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 只能被1或者自身整除的数叫质数
func countPrimes(n int) int {
	if n < 3 {
		return 0
	}
	// f[i] == true 表示i不是质数
	// f[i] == false 表示i是质数
	f := make([]bool, n)
	// 偶数不是质数，所以砍一半
	conunt := n / 2
	// 枚举所有的奇数因子
	for i := 3; i*i < n; i += 2 { // 3 5 7 9
		if f[i] {
			continue
		}
		// 枚举i的因子不是质数的数( i * i + 2i ), 这个公式能跳过已经验过的数
		for j := i * i; j < n; j += 2 * i { // 9 15 21
			if !f[j] {
				conunt--
				f[j] = true
			}
		}
	}
	return conunt
}
