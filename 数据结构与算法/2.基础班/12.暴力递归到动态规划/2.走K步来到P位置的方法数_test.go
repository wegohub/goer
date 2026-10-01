package class12

import (
	"fmt"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	var (
		N int
		M int
		K int
		P int
	)
	defer func(N, M, K, P int) {
		if err := recover(); err != nil {
			fmt.Println("err:", err, N, M, K, P)
		}
	}(N, M, K, P)

	testTimes := 100000
	maxVal := 10

	for i := 0; i < testTimes; i++ {
		// [2, 100]
		N = int(rand.Float64()*float64(maxVal)) + 2
		// [1, N]
		M = int(rand.Float64()*float64(N)) + 1
		// [1, N]
		K = int(rand.Float64()*float64(N)) + 1
		// [1, N]
		P = int(rand.Float64()*float64(N)) + 1

		ans1 := Num(N, M, K, P)
		ans2 := NumCache(N, M, K, P)
		ans3 := NumDP(N, M, K, P)

		if ans1 != ans2 || ans2 != ans3 {
			fmt.Println("Err: ", N, M, K, P)
			break
		}
	}

	fmt.Println("success")

	return "Hello World!", nil
}

// 暴力递归
func Num(N, M, K, P int) int {
	return process(N, M, K, P)
}

func process(N, cur, rest, P int) int {
	if rest == 0 {
		if cur == P {
			return 1
		} else {
			return 0
		}
	}

	if cur == 1 {
		return process(N, 2, rest-1, P)
	}

	if cur == N {
		return process(N, N-1, rest-1, P)
	}

	return process(N, cur+1, rest-1, P) + process(N, cur-1, rest-1, P)
}

// 记忆化搜索(动态规划)
func NumCache(N, M, K, P int) int {
	dp := make([][]int, N+1) // 总长度
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j < K+1; j++ {
			dp[i][j] = -1
		}
	}
	return processCache(N, M, K, P, dp)
}

func processCache(N, cur, rest, P int, dp [][]int) int {
	if dp[cur][rest] != -1 {
		return dp[cur][rest]
	}

	if rest == 0 {
		if cur == P {
			dp[cur][rest] = 1
			return 1
		} else {
			dp[cur][rest] = 0
			return 0
		}
	}

	if cur == 1 {
		dp[cur][rest] = process(N, 2, rest-1, P)
		return dp[cur][rest]
	}

	if cur == N {
		dp[cur][rest] = process(N, N-1, rest-1, P)
		return dp[cur][rest]
	}

	dp[cur][rest] = process(N, cur+1, rest-1, P) + process(N, cur-1, rest-1, P)
	return dp[cur][rest]
}

// 经典动态规划
func NumDP(N, M, K, P int) int {
	dp := make([][]int, N+1) // 总长度
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, K+1)
	}
	dp[P][0] = 1

	for rest := 1; rest < K+1; rest++ {
		for cur := 1; cur < N+1; cur++ {
			if cur == 1 {
				dp[cur][rest] = dp[2][rest-1]
			} else if cur == N {
				dp[cur][rest] = dp[N-1][rest-1]
			} else {
				dp[cur][rest] = dp[cur+1][rest-1] + dp[cur-1][rest-1]
			}
		}
	}

	return dp[M][K]
}
