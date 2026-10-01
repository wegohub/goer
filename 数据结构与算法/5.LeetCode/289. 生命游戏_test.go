package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func gameOfLife(board [][]int) {
	N := len(board)
	M := len(board[0])
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			// 周围活细胞的数量
			neighborsNum := neighbors(board, i, j)
			if neighborsNum == 3 || (board[i][j] == 1 && neighborsNum == 2) {
				set(board, i, j)
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			board[i][j] = get(board, i, j)
		}
	}
}

// 获取领居的数量
func neighbors(board [][]int, i, j int) int {
	count := 0
	count += getCount(board, i-1, j-1)
	count += getCount(board, i-1, j)
	count += getCount(board, i-1, j+1)
	count += getCount(board, i, j-1)
	count += getCount(board, i, j+1)
	count += getCount(board, i+1, j-1)
	count += getCount(board, i+1, j)
	count += getCount(board, i+1, j+1)
	return count
}

func getCount(board [][]int, i, j int) int {
	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) && (board[i][j]&1) == 1 {
		return 1
	} else {
		return 0
	}
}

// 利用二进制位的第二位标记修改后的值
func set(board [][]int, i, j int) {
	board[i][j] |= (1 << 1)
}

func get(board [][]int, i, j int) int {
	return board[i][j] >> 1
}
