package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	N := len(matrix)
	M := len(matrix[0])

	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, M)
	}
	// 最大边长
	maxEdge := 0
	// 第一列
	for i := 0; i < N; i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
			maxEdge = 1
		}
	}
	// 第一行
	for j := 1; j < M; j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
			maxEdge = 1
		}
	}
	// 普遍位置
	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = Min(Min(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]) + 1
				maxEdge = Max(maxEdge, dp[i][j])
			}
		}
	}

	return maxEdge * maxEdge
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
