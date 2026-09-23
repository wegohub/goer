package d22

import (
	"math/rand/v2"
	"testing"
)

var benchmarkSink int64

// TODO: 编译成汇编
func BenchmarkBitwise(b *testing.B) {

	const (
		maxNum       int64 = 1_000_000
		elementCount       = 100_000
	)

	values := make([]int64, elementCount)
	for i := range values {
		values[i] = rand.Int64N(maxNum)
	}

	// 取模运算
	// num % 64 == num & 63
	// 12345 % 100 = 45
	b.Run("modulo", func(b *testing.B) {
		b.ResetTimer()
		var sum int64
		for i := 0; i < b.N; i++ {
			for _, value := range values {
				sum += value % 64
			}
		}
		benchmarkSink = sum
	})

	// 位运算
	b.Run("bitwise", func(b *testing.B) {
		b.ResetTimer()
		var sum int64
		for i := 0; i < b.N; i++ {
			for _, value := range values {
				sum += value & 63
			}
		}
		benchmarkSink = sum
	})
}
