package class09

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func spiralOrder(matrix [][]int) []int {
	// 左上角的点
	x1 := 0
	y1 := 0
	// 右下角的点
	x2 := len(matrix) - 1
	y2 := len(matrix[0]) - 1

	ans := make([]int, 0, len(matrix)*len(matrix[0]))
	for x1 <= x2 && y1 <= y2 {
		spiralOrderEdge(matrix, x1, y1, x2, y2, &ans)
		x1++
		y1++
		x2--
		y2--
	}

	return ans
}

func spiralOrderEdge(matrix [][]int, x1, y1, x2, y2 int, ans *[]int) {
	// 共行
	if x1 == x2 {
		for i := y1; i <= y2; i++ {
			*ans = append(*ans, matrix[x1][i])
		}
		// 共列
	} else if y1 == y2 {
		for i := x1; i <= x2; i++ {
			*ans = append(*ans, matrix[i][y1])
		}
	} else {
		curX := x1
		curY := y1
		for curY != y2 {
			*ans = append(*ans, matrix[x1][curY])
			curY++
		}

		for curX != x2 {
			*ans = append(*ans, matrix[curX][y2])
			curX++
		}

		for curY != y1 {
			*ans = append(*ans, matrix[x2][curY])
			curY--
		}

		for curX != x1 {
			*ans = append(*ans, matrix[curX][y1])
			curX--
		}
	}
}
