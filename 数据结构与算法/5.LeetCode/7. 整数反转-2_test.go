package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(reverse(-123))
	return "Hello World!", nil
}

func reverse(x int) int {
	var (
		num int32 = int32(x)
		// 常数边界
		m int32 = math.MinInt32 / 10
		// 余数边界
		o   int32 = math.MinInt32 % 10
		ans int32 = 0
	)
	// num是否为负数
	isNeg := (num >> 31 & 1) == 1
	// 如果num是正数，转为负数处理
	if !isNeg {
		num = ^num + 1
	}
	for num != 0 {
		if ans < m || (ans == m && ans%10 < o) {
			return 0
		}
		ans = ans*10 + num%10
		num = num / 10
	}
	if !isNeg {
		ans = ^ans + 1
	}
	return int(ans)
}
