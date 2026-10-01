package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 5因子的数量
// 2因子 * 5因子 = 10就会配出一个0来
func trailingZeroes(n int) int {
	// 每5个数就会出现一个5的因子，第一轮 n/5
	// 没25个数又会多一个5因子，n/125
	ans := 0
	for n != 0 {
		n /= 5 // 5 25 125....
		ans += n
	}
	return ans
}
