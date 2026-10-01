package class11

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	weight := []int{1, 2, 3, 4, 5}
	value := []int{5, 4, 3, 2, 1}
	bag := 5
	ans1 := backpack(weight, value, bag)
	fmt.Println(ans1)
	return "Hello World!", nil
}

func backpack(weights []int, values []int, bag int) int {
	return process(weights, values, 0, bag)
}

func process(weights []int, values []int, index int, rest int) int {
	if rest < 0 {
		return -1
	}
	if index == len(weights) {
		return 0
	}

	// 不要
	p1 := process(weights, values, index+1, rest)
	// 要
	p2 := -1
	p2Next := process(weights, values, index+1, rest-weights[index])
	if p2Next != -1 {
		p2 = p2Next + values[index]
	}
	return int(math.Max(float64(p1), float64(p2)))
}
