package leetcode

import (
	"fmt"
	"strconv"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	solveSudoku(board)
	fmt.Println(board)
	return "Hello World!", nil
}

func solveSudoku(board [][]byte) {
	var (
		row    [9][10]int // 在第几行上有没有数字
		col    [9][10]int // 在第几列上有没有数字
		bucket [9][10]int // 在第几个方块内有没有数字
	)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				//         确定哪一行      确定哪一列
				gridNum := 3*(i/3) + j/3
				num := int(board[i][j] - '0')
				row[i][num] = 1
				col[j][num] = 1
				bucket[gridNum][num] = 1
			}
		}
	}

	process(board, 0, 0, row, col, bucket)
}

func process(board [][]byte, i int, j int, row, col, bucket [9][10]int) bool {
	if i == len(board) {
		return true
	}

	// 下次递归处理行
	nextI := i
	if j == 8 {
		nextI = i + 1
	}
	nextJ := j + 1
	if j == 8 {
		nextJ = 0
	}
	if board[i][j] != '.' {
		return process(board, nextI, nextJ, row, col, bucket)
	} else {
		gridNum := 3*(i/3) + j/3
		for num := 1; num <= 9; num++ {
			if row[i][num] == 0 && col[j][num] == 0 && bucket[gridNum][num] == 0 {
				row[i][num] = 1
				col[j][num] = 1
				bucket[gridNum][num] = 1
				board[i][j] = strconv.Itoa(num)[0]
				if process(board, nextI, nextJ, row, col, bucket) {
					return true
				}
				row[i][num] = 0
				col[j][num] = 0
				bucket[gridNum][num] = 0
				board[i][j] = '.'
			}
		}
		return false
	}
}
