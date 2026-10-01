package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func rotate(matrix [][]int) {
	// 左上角点
	tR := 0
	tC := 0
	// 右上角点
	dR := len(matrix) - 1
	dC := len(matrix) - 1

	// 往中间移动
	for tR < dR {
		rotateEdge(matrix, tR, tC, dR, dC)
		tR++
		tC++
		dR--
		dC--
	}
}

func rotateEdge(matrix [][]int, tR, tC, dR, dC int) {
	// i是组号
	for i := 0; i < dC-tC; i++ {
		//tmp := matrix[tR][tC + i]
		//matrix[tR][tC + i] = matrix[dR - i][tC]
		//matrix[dR - i][tC] = matrix[dR][dC-i]
		//matrix[dR][dC-i] = matrix[tR+i][dC]
		//matrix[tR+i][dC] = tmp
		tmp := matrix[tR][tC+i]
		matrix[tR][tC+i] = matrix[dR-i][tC]
		matrix[dR-i][tC] = matrix[dR][dC-i]
		matrix[dR][dC-i] = matrix[tR+i][dC]
		matrix[tR+i][dC] = tmp
	}
}
