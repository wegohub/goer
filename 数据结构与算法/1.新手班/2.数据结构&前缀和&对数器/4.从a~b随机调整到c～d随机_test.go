package class02

import (
	"fmt"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	testTimes := 100000
	counts := make([]int, 10)
	for i := 0; i < testTimes; i++ {
		counts[g()]++
	}
	fmt.Println(counts)
	return "Hello World!", nil
}

// 等概率返回[1,5]
func f() int {
	// [0, 4]
	return int(rand.Float64()*5) + 1
}

// 等概率01发生器
func zeroOne() int {
do:
	ans := f()
	if ans == 3 {
		goto do
	} else if ans > 3 {
		return 1
	} else {
		return 0
	}
}

// 用三个二进制位产生[0,7]的等概率发生器
func binaryGen() int {
	return (zeroOne() << 2) + (zeroOne() << 1) + (zeroOne() << 0)
}

// 等概率返回[1,7]
func g() int {
do:
	ans := binaryGen()
	// 产生[0,6]随机
	if ans == 7 {
		goto do
	} else {
		// 产生[1,7]随机
		return ans + 1
	}
}
