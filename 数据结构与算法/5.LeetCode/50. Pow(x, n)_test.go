package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(myPow(2.5, -2))

	return "Hello World!", nil
}

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	// 防止溢出，系统最小转不成正数
	pow := n
	if n == math.MinInt32 {
		pow = n + 1
	}

	// n转换为绝对值
	pow = int(math.Abs(float64(pow)))

	// n是正数
	var ans float64 = 1
	t := x
	for pow != 0 {
		if pow&1 != 0 {
			ans = ans * t
		}

		t = t * t
		pow = pow >> 1
	}

	// 如果是负数在乘x, 因为前面n+1
	if n == math.MinInt32 {
		ans = ans * x
	}

	// 如果n是负数
	if n < 0 {
		return float64(1) / ans
	}
	return ans
}
