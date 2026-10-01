package class12

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
	ans2 := backpackCache(weight, value, bag)
	ans3 := backpackDP(weight, value, bag)
	fmt.Println(ans1, ans2, ans3)
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
	if p2Next != -1 { // 容量要装的下才能获取到价值
		p2 = p2Next + values[index]
	}
	return int(math.Max(float64(p1), float64(p2)))
}

// 记忆化搜索方式
func backpackCache(weights []int, values []int, bag int) int {
	n := len(weights)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		tmp := make([]int, bag+1)
		for j := 0; j < bag+1; j++ {
			tmp[j] = -1
		}
		dp[i] = tmp
	}
	return processCache(weights, values, 0, bag, dp)
}

func processCache(weights []int, values []int, index int, rest int, dp [][]int) int {
	if rest < 0 {
		return -1
	}
	if dp[index][rest] != -1 {
		return dp[index][rest]
	}

	if index == len(weights) {
		dp[index][rest] = 0
		return dp[index][rest]
	}

	// 不要
	p1 := process(weights, values, index+1, rest)
	// 要
	p2 := -1
	p2Next := process(weights, values, index+1, rest-weights[index])
	if p2Next != -1 { // 容量要装的下才能获取到价值
		p2 = p2Next + values[index]
	}
	dp[index][rest] = int(math.Max(float64(p1), float64(p2)))
	return dp[index][rest]
}

// 动态规划
func backpackDP(weights []int, values []int, bag int) int {
	N := len(weights)
	if N < 0 || len(values) < 0 || len(weights) != len(values) || bag < 0 {
		return 0
	}
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, bag+1)
	}

	// 潜台词：index == N 为 0
	for index := N - 1; index >= 0; index-- {
		for rest := 0; rest < bag+1; rest++ {
			// 不要
			p1 := dp[index+1][rest]
			// 要
			p2 := -1
			if rest-weights[index] >= 0 {
				p2 = dp[index+1][rest-weights[index]] + values[index]
			}
			dp[index][rest] = int(math.Max(float64(p1), float64(p2)))
		}
	}
	return dp[0][bag]
}
