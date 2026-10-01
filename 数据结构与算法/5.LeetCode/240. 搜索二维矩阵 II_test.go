package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 从右上角往左下角走
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	N := len(matrix)
	M := len(matrix[0])
	row := 0
	col := M - 1
	for row < N && col >= 0 {
		if matrix[row][col] > target { // 往左走
			col--
		} else if matrix[row][col] < target { // 往下走
			row++
		} else {
			return true
		}
	}

	return false
}
