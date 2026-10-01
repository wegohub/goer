package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// minPathSum 返回从左上角到右下角的最小路径和
func minPathSum(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	N := len(grid)
	M := len(grid[0])
	// 其实应该是一张二维表，但是用了空间压缩技巧
	dp := make([]int, M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if i == 0 && j == 0 {
				dp[j] = grid[i][j]
			} else {
				dp[j] = min((func() int {
					if i > 0 {
						return dp[j]
					} else {
						return math.MaxInt32
					}
				})(),
					(func() int {
						if j > 0 {
							return dp[j-1]
						} else {
							return math.MaxInt32
						}
					})()) + grid[i][j]
			}
		}
	}
	return dp[M-1]
}

// minPathSum2 返回从左上角到右下角的最小路径和，使用空间压缩技巧
func minPathSum2(grid [][]int) int {
	if grid == nil || len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	N := len(grid)
	M := len(grid[0])
	// 其实应该是一张二维表，但是用了空间压缩技巧
	dp := make([]int, M)
	// dp数组变成，想象中二维表的第0行数据
	// m : 3 2 1 6           3 5 6 12
	dp[0] = grid[0][0]
	for i := 1; i < M; i++ {
		dp[i] = dp[i-1] + grid[0][i]
	}
	for i := 1; i < N; i++ {
		// dp此时是想象中二维表的第i-1行数据
		// 想更新成，想象中二维表的第i行数据
		// dp[0]
		dp[0] = dp[0] + grid[i][0]
		for j := 1; j < M; j++ {
			dp[j] = min(dp[j-1], dp[j]) + grid[i][j]
		}
	}
	return dp[M-1]
}

// min 返回两个整数中的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
