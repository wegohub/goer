package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(isHappy(886))
	return "Hello World!", nil
}

func isHappy(n int) bool {
	// 记录循环
	set := make(map[int]struct{})
	for n != 1 {
		sum := 0
		for n != 0 {
			r := n % 10
			sum += r * r
			n /= 10
		}
		n = sum
		if _, ok := set[n]; ok {
			break
		}
		set[n] = struct{}{}
	}
	return n == 1
}

// 数学上的证明, 最后都会落到1或4上
func isHappy2(n int) bool {
	for n != 1 && n != 4 {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n /= 10
		}
		n = sum
	}
	return n == 1
}
