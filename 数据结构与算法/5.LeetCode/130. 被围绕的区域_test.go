package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 从缺口处反向感染
// 1. 从缺口处开始标记该片区域不能被感染
// 2. 遍历所有的区域开始感染
func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	N := len(board)
	M := len(board[0])
	// 左右缺口标记
	for j := 0; j < M; j++ {
		if board[0][j] == 'O' {
			free(board, 0, j)
		}
		if board[N-1][j] == 'O' {
			free(board, N-1, j)
		}
	}
	// 上下缺口标记
	for i := 0; i < N; i++ {
		if board[i][0] == 'O' {
			free(board, i, 0)
		}
		if board[i][M-1] == 'O' {
			free(board, i, M-1)
		}
	}

	// 感染
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			}
		}
	}

}

// 防感染区域标记
func free(board [][]byte, i, j int) {
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) || board[i][j] != 'O' {
		return
	}
	board[i][j] = 'F'
	free(board, i-1, j)
	free(board, i+1, j)
	free(board, i, j-1)
	free(board, i, j+1)
}
