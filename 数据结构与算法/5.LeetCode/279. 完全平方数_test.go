package leetcode

import (
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 输入参数简单的一个数，尝试打表技巧
func numSquares1(n int) int {
	res := n
	num := 2
	for num*num <= n {
		a := n / (num * num)
		b := n % (num * num)
		res = Min(res, a+numSquares1(b))
		num++
	}
	return res
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 1 : 1, 4, 9, 16, 25, 36, ...
// 4 : 7, 15, 23, 28, 31, 39, 47, 55, 60, 63, 71, ...
// 规律解
// 规律一：个数不超过4
// 规律二：出现1个的时候，显而易见
// 规律三：任何数 % 8 == 7，一定是4个
// 规律四：任何数消去4的因子之后，剩下rest，rest % 8 == 7，一定是4个
func numSquares(n int) int {
	rest := n
	// 消掉4的因子
	for rest%4 == 0 {
		rest /= 4
	}

	if rest%8 == 7 {
		return 4
	}

	f := int(math.Sqrt(float64(n)))
	if f*f == n {
		return 1
	}

	for first := 1; first*first <= n; first++ {
		second := int(math.Sqrt(float64(n - first*first)))
		if first*first+second*second == n {
			return 2
		}
	}

	return 3
}

func numSquares2(n int) int {

	f := int(math.Sqrt(float64(n)))
	if f*f == n {
		return 1
	}

	rest := n
	for rest%4 == 0 {
		rest /= 4
	}

	for first := 1; first*first <= n; first++ {
		second := int(math.Sqrt(float64(n - first*first)))
		if first*first+second*second == n {
			return 2
		}
	}

	if rest%8 == 7 {
		return 4
	}

	return 3
}
