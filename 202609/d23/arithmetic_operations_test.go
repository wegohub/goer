// Package d23 位运算实现加减乘除运算
// LetCode原题: https://leetcode.cn/problems/divide-two-integers/
package d23

import "math"

// 加法运算
func add(a, b int) int {
	sum := a
	// 进位不为0
	for b != 0 {
		// 无进位相加
		sum = a ^ b
		// a + b 的进位
		b = (a & b) << 1
		a = sum
	}
	return sum
}

// 减法
func minus(a, b int) int {
	// a - b = a + (-b)
	return add(a, negNum(b))
}

// 乘法运算
func multi(a, b int) int {
	res := 0
	for b != 0 {
		if b&1 != 0 {
			res = add(res, a)
		}
		a <<= 1
		b = int(uint(b) >> 1) // 不带符号右移
	}
	return res
}

// 除法(a, b不能是系统最小)
func div(a, b int) int {
	x := a
	if isNeg(a) {
		x = negNum(a)
	}
	y := b
	if isNeg(b) {
		y = negNum(b)
	}
	res := 0
	for i := 31; i > negNum(1); i = minus(i, 1) {
		if (x >> i) >= y {
			res |= 1 << i
			x = minus(x, y<<i)
		}
	}
	if isNeg(a) != isNeg(b) {
		return negNum(res)
	}
	return res
}

// dividend被除数和divisor除数
func divide(dividend, divisor int) int {
	// 如果除数是系统最小
	if divisor == math.MinInt32 {
		// 如果被除数是系统最小
		if dividend == math.MinInt32 {
			return 1
		}
		return 0
	}
	// 除数不是系统最小，被除数是系统最小
	if dividend == math.MinInt32 {
		if divisor == negNum(1) { // 系统最小/-1 = 系统最大
			return math.MaxInt32
		}
		// 因为系统最小无法转成正数(负数比正数多一个)
		// 假设 -10 系统最小， 9系统最大
		// -10 / 2
		// res = (-10 + 1) / 5 = -1
		// 补偿： -10 - (-1 * 5) = -5 / 5 = -1
		// res = -1 + -1 = -2
		res := div(add(dividend, 1), divisor)                               // 被除数+1 / 除数 => 防止溢出
		return add(res, div(minus(dividend, multi(res, divisor)), divisor)) // 补偿
	}
	// 被除数和除数都不是系统最小
	return div(dividend, divisor)
}

// 求n的相反数
func negNum(n int) int {
	return add(^n, 1)
}

// n是否是负数
func isNeg(n int) bool {
	return n < 0
}
