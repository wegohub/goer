package GoBeginnerLevel

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"testing"
)

func Test_Binary_Print(t *testing.T) {
	fmt.Printf("32位二进制输出: %032b\n", 123456)
	// 位编号 位权 位宽

	// 一个数的相反数: 按位取反 + 1
	var x int32 = 256
	fmt.Println(^x + 1)

	// 原码 反码 补码 => 把减法运算转换为加法  => 负数进行加法运算
	var y int8 = -3
	fmt.Printf("负数二进制1: %08b\n", uint8(y))
	// 怎么得到的
	// 3 的原码: 00000011
	// 取反:     11111100
	// +1:      11111101

	// int8 的最大值: -2^7 ~ 2^7-1  => -128 ~ 127
	// 溢出怎么处理:
	var z int8 = math.MaxInt8 // 1 << 7 - 1
	fmt.Println("溢出: ", z+1)  // -128  为什么？ java里面就报错了
	// go的运算只看补码!!!
	// 127的补码(正数是他自己):   01111111
	// 1的补码(正式是他自己):     00000001
	// 相加:                    10000000
	// -128: 正数溢出时，结果按 2ⁿ 取模回绕到负数区间。只有 +1 时恰好跳到最小负数，其他值按实际偏移量对应。

	// uint8(y) 发生了什么?
	// -3 在内存中是以补码的形式存在: 11111101  [-3 的原码：10000011; -3 的反码：11111100（符号位不变，数值位取反）;-3 的补码：11111101（反码 + 1）]
	// uint(y) 不会重新计算，直接当前8位无符号数读取: 11111101 = 128 + 64 + 32 + 16 + 8 + 4 + 1 = 253, 所以 uint8(-3) == 253

	// 补码转原码： 1. 先减1，再取反  2. 再求一次补码
}

// Test_PrintB 打印一个数的二进制位
// num： 00000001110101010
//
//	1000000000000000
//
// 0 - 31
// num & (1 << i) == 0 ? "0" : "1"
func Test_PrintB(t *testing.T) {

	var num int32 = -5

	sb := strings.Builder{}
	for i := 31; i >= 0; i-- {
		bit := "1"
		if (num & (1 << i)) == 0 {
			bit = "0"
		}
		sb.WriteString(bit)
	}

	fmt.Println(sb.String())
}

// 给定一个正整数参数N, 返回 1! + 2! + 3! + ... + N! 的结果
func Test_Factorial(t *testing.T) {
	// 对数器
	maxStep := 10000
	for i := 0; i < maxStep; i++ {
		N := rand.Intn(100)
		low := FactorialLow(N)
		high := FactorialHigh(N)
		if low != high {
			t.Fatalf("对数出错, low: %d, high: %d", low, high)
		}
	}
}

func FactorialLow(N int) int {
	ans := 0
	for i := 1; i <= N; i++ {
		inter := 1
		for j := 1; j <= i; j++ {
			// fix： 溢出
			if inter > math.MaxInt/j {
				return 0
			}
			inter = inter * j
		}
		// fix: 溢出
		if ans > math.MaxInt-inter {
			return 0
		}
		ans += inter
	}

	return ans
}

func FactorialHigh(N int) int {
	ans := 0
	pre := 1
	for i := 1; i <= N; i++ {
		// fix: 乘法溢出, 返回0
		if pre > math.MaxInt/i {
			return 0
		}
		pre = pre * i

		// fix: 加法溢出, 返回0
		if ans > math.MaxInt-pre {
			return 0
		}
		ans += pre
	}

	return ans
}

func Benchmark_Factorial(b *testing.B) {
	// 固定 N=20
	b.Run("N=20", func(b *testing.B) {
		b.Run("Low", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialLow(20)
			}
		})
		b.Run("High", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialHigh(20)
			}
		})
	})

	// 固定 N=10
	b.Run("N=10", func(b *testing.B) {
		b.Run("Low", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialLow(10)
			}
		})
		b.Run("High", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialHigh(10)
			}
		})
	})

	// 随机 N (1-20)
	b.Run("Random", func(b *testing.B) {
		r := rand.New(rand.NewSource(42))
		b.ResetTimer()

		b.Run("Low", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialLow(r.Intn(20) + 1)
			}
		})
		b.Run("High", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				FactorialHigh(r.Intn(20) + 1)
			}
		})
	})
}

func Test_Swap_Num(t *testing.T) {
	fmt.Println(SwapNum(1, 2))
}

func SwapNum(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
