package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(reverse2(123))

	return "Hello World!", nil
}

// 最优解
func reverse2(x int) int {
	num := int32(x)
	// 是否是负数
	isNeg := false
	if (num >> 31 & 1) == 1 {
		isNeg = true
	}
	// 正数转为负数处理，因为负数的绝对值域 > 1
	if !isNeg {
		num = ^num + 1
	}

	var m int32 = math.MinInt32 / 10
	var o int32 = math.MinInt32 % 10

	var ans int32
	for num != 0 {
		// ans * 10      num % 10系数
		if ans < m || ans == m && num%10 < o {
			return 0
		}

		ans = ans*10 + num%10
		num = num / 10
	}

	// 如果是正数要转成绝对值
	if !isNeg {
		ans = int32(math.Abs(float64(ans)))
	}
	return int(ans)
}

func reverse(x int) int {
	num := int32(x)
	// num 位数
	bit := dig(num)

	var ans int32
	for i := bit; i >= 0; i-- {
		curNum := num / int32(math.Pow(10, float64(i)))
		num = num % int32(math.Pow(10, float64(i)))

		// 处理溢出，此时的ans能不能加后面的一坨
		var curVal int32
		if math.Abs(float64(curNum)) > 2 && bit-i >= 9 {
			return 0
		} else {
			curVal = curNum * int32(math.Pow(10, float64(bit-i)))
		}
		// curVal 可能溢出
		if x < 0 && curVal < math.MinInt32-ans {
			return 0
		}
		if x > 0 && curVal > math.MaxInt32-ans {
			return 0
		}
		ans += curVal
	}

	return int(ans)
}

// 计算num有几位-1
func dig(num int32) int {
	ans := 0
	for num/10 != 0 {
		num = num / 10
		ans++
	}
	return ans
}
