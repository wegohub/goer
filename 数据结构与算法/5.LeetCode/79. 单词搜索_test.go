package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if process(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 从board[i][j] 位置出发能不能搞定 word从k出发的所有
// 三个可变参数无法改动态规划
func process(board [][]byte, word string, i int, j int, k int) bool {
	// word走完了
	if k == len(word) {
		return true
	}
	// 越界位置
	if i < 0 || i == len(board) || j < 0 || j == len(board[0]) {
		return false
	}
	// 字符不同
	if board[i][j] != word[k] {
		return false
	}

	tmp := board[i][j]
	// 标记已经走过了
	board[i][j] = 0
	ans := process(board, word, i+1, j, k+1) ||
		process(board, word, i-1, j, k+1) ||
		process(board, word, i, j-1, k+1) ||
		process(board, word, i, j+1, k+1)
	// 恢复现场
	board[i][j] = tmp

	return ans
}
