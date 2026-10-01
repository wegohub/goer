package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	var row [9][10]int
	fmt.Println(row)
	return "Hello World!", nil
}

func isValidSudoku(board [][]byte) bool {
	var (
		row    [9][10]int // 在第几行上有没有数字
		col    [9][10]int // 在第几列上有没有数字
		bucket [9][10]int // 在第几个方块内有没有数字
	)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				gridNum := 3*(i/3) + j/3
				num := int(board[i][j] - '0')
				if row[i][num] == 1 || col[j][num] == 1 || bucket[gridNum][num] == 1 {
					return false
				}
				row[i][num] = 1
				col[j][num] = 1
				bucket[gridNum][num] = 1
			}
		}
	}

	return true
}
