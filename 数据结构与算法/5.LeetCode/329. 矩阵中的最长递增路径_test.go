package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 时间复杂度: O(N*M)
func longestIncreasingPath(matrix [][]int) int {
	longest := 0
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			longest = Max(longest, process(matrix, i, j, dp))
		}
	}

	return longest
}

func process(matix [][]int, i, j int, dp [][]int) int {
	if dp[i][j] != 0 {
		return dp[i][j]
	}

	next := 0
	// 上
	if i > 0 && matix[i-1][j] > matix[i][j] {
		next = process(matix, i-1, j, dp)
	}
	// 下
	if i+1 < len(matix) && matix[i+1][j] > matix[i][j] {
		next = Max(next, process(matix, i+1, j, dp))
	}
	// 左
	if j > 0 && matix[i][i-1] > matix[i][j] {
		next = Max(next, process(matix, i, j-1, dp))
	}
	// 右
	if j+1 < len(matix[0]) && matix[i][j+1] > matix[i][j] {
		next = Max(next, process(matix, i, j+1, dp))
	}
	dp[i][j] = 1 + next
	return 1 + next
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
