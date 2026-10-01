package class02

import (
	"fmt"
	"math"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// rand.Float64() 等概率返回[0.0,1.0)上的数
	testTimes := 100000
	count := 0
	for i := 0; i < testTimes; i++ {
		if rand.Float64() < 0.75 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(testTimes))
	fmt.Println("=====================")

	// [0,1) -> [0,8)
	count = 0
	for i := 0; i < testTimes; i++ {
		if rand.Float64()*8 < 5 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(testTimes))
	fmt.Println(float64(5) / float64(8))

	fmt.Println("=======================")
	// K == 9, int(rand.Float64() * float64(K)) => [0,8]
	K := 9
	counts := make([]int, K+1)
	for i := 0; i < testTimes; i++ {
		ans := int(rand.Float64() * float64(K))
		counts[ans]++
	}
	fmt.Println(counts)

	fmt.Println("======================")

	// 概率由x调整为x^2
	count = 0
	x := 0.17
	for i := 0; i < testTimes; i++ {
		if xToXPower() < x {
			count++
		}
	}
	fmt.Println(float64(count) / float64(testTimes))
	fmt.Println(math.Pow(float64(x), 2))

	fmt.Println("======================")

	return "Hello World!", nil
}

// 返回[0,1)上的一个数
// 任意的x,x属于[0,1)，[0,x)范围上的数出现的概率由原来的x调整成x的平方
func xToXPower() float64 {
	// 两次撸出来的数都在0-x上概率才是x^2，所以是math.Max, 大的那个都在0-x，小的那个必定在0-x
	return math.Max(rand.Float64(), rand.Float64())
}
