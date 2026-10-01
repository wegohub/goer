package class12

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	drinks := []int{1, 1, 5, 5, 7, 10, 12, 12, 12, 12, 12, 12, 15}
	a := 3
	b := 10
	fmt.Println(process(drinks, a, b, 0, 0))
	fmt.Println(processDP(drinks, a, b))
	return "Hello World!", nil
}

// a 洗一杯的时间
// b 自己挥发干净的时间
// drinks 每个员工喝完的时间
// index 当前来到第几杯
// washLine 表示机器何时可用
func process(drinks []int, a int, b int, index int, washLine int) int {
	if index == len(drinks)-1 {
		return int(math.Min(
			// 决定洗
			math.Max(float64(washLine), float64(drinks[index]))+float64(a),
			// 决定挥发
			float64(drinks[index]+b),
		))
	}

	// 剩下不止一杯咖啡

	// 决定洗
	// 洗完index这杯的时间
	wash := int(math.Max(float64(washLine), float64(drinks[index]))) + a
	// index+1变干净的最早时间
	next1 := process(drinks, a, b, index+1, wash)
	p1 := int(math.Max(float64(wash), float64(next1)))

	// 决定挥发
	dry := drinks[index] + b
	next2 := process(drinks, a, b, index+1, washLine)
	p2 := int(math.Max(float64(dry), float64(next2)))

	return int(math.Min(float64(p1), float64(p2)))
}

func processDP(drinks []int, a int, b int) int {
	// 挥发时间大于洗的时间，最后一杯喝完的时间+挥发时间
	if a > b {
		return drinks[len(drinks)-1] + b
	}
	N := len(drinks)
	limit := 0
	for _, drink := range drinks {
		limit = int(math.Max(float64(limit), float64(drink))) + a
	}

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, limit+1)
	}

	for washLine := 0; washLine <= limit; washLine++ {
		dp[N-1][washLine] = int(math.Min(
			// 决定洗
			math.Max(float64(washLine), float64(drinks[N-1]))+float64(a),
			// 决定挥发
			float64(drinks[N-1]+b),
		))

		for index := N - 2; index >= 0; index-- {
			for washLine := 0; washLine <= limit; washLine++ {
				// 决定洗
				// 洗完index这杯的时间
				p1 := math.MaxInt
				wash := int(math.Max(float64(washLine), float64(drinks[index]))) + a
				if wash <= limit {
					p1 = int(math.Max(float64(wash), float64(dp[index+1][wash])))
				}
				// 决定挥发
				p2 := int(math.Max(float64(drinks[index]+b), float64(dp[index+1][washLine])))

				dp[index][washLine] = int(math.Min(float64(p1), float64(p2)))

			}
		}

	}

	return dp[0][0]
}
