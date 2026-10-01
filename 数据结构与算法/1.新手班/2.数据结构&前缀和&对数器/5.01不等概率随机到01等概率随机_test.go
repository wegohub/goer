package class02

import (
	"fmt"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	testTimes := 100000
	count := 0
	for i := 0; i < testTimes; i++ {
		if g() == 0 {
			count++
		}
	}
	fmt.Println(float64(count) / float64(testTimes))

	return "Hello World!", nil
}

func f() int {
	if rand.Float64() < 0.84 {
		return 0
	}
	return 1
}

// f() 产生的结果
// 1 1 -> 0.16 * 0.16
// 0 0 -> 0.84 * 0.84
// 1 0 -> 0.16 * 0.84
// 0 1 -> 0.84 * 0.16
// 显然产生 1 0 或 0 1 的概率是相等的
func g() int {
do:
	ans1 := f()
	ans2 := f()
	if ans1 == ans2 {
		goto do
	} else if ans1 == 1 {
		return 1
	} else {
		return 0
	}
}
