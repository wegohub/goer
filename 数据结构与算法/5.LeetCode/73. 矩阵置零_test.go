package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func setZeroes(matrix [][]int) {
	// 第一行是否是0
	row0Zero := false
	// 第一列是否是0
	col0Zero := false

	// 看第一行变0否
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			row0Zero = true
			break
		}
	}
	// 看第一列变0否
	for j := 0; j < len(matrix); j++ {
		if matrix[j][0] == 0 {
			col0Zero = true
			break
		}
	}

	// 看普遍位置的行列是否变0
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	// 更改普遍位置的值
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 改第一行的值
	if row0Zero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}

	// 改第一列的值
	if col0Zero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}
