package class09

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 求x y 的最大公约数(质数)
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
