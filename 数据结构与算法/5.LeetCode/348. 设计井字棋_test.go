package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// TicTacToe represents a Tic-Tac-Toe game.
type TicTacToe struct {
	rows    [][]int
	cols    [][]int
	leftUp  []int
	rightUp []int
	matrix  [][]bool
	N       int
}

// NewTicTacToe initializes a new TicTacToe game with the given size.
func Constructor(n int) TicTacToe {
	rows := make([][]int, n) // 记录每一行玩家1和玩家2下了几回
	cols := make([][]int, n) // 记录每一列玩家1和玩家2下了几回
	for i := 0; i < n; i++ {
		rows[i] = make([]int, 3)
		cols[i] = make([]int, 3)
	}
	leftUp := make([]int, 3)    // 记录左对角线玩家1和玩家2下了几回
	rightUp := make([]int, 3)   // 记录右对角线玩家1和玩家2下了几回
	matrix := make([][]bool, n) // 该位置有没有下过
	for i := 0; i < n; i++ {
		matrix[i] = make([]bool, n)
	}
	return TicTacToe{
		rows:    rows,
		cols:    cols,
		leftUp:  leftUp,
		rightUp: rightUp,
		matrix:  matrix,
		N:       n,
	}
}

// Move makes a move for the given player at the specified row and column.
func (t *TicTacToe) Move(row, col, player int) int {
	if t.matrix[row][col] {
		return 0
	}
	t.matrix[row][col] = true
	t.rows[row][player]++
	t.cols[col][player]++
	if row == col {
		t.leftUp[player]++
	}
	if row+col == t.N-1 {
		t.rightUp[player]++
	}
	if t.rows[row][player] == t.N || t.cols[col][player] == t.N || t.leftUp[player] == t.N || t.rightUp[player] == t.N {
		return player
	}
	return 0
}
