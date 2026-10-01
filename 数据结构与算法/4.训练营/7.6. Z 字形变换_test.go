package train

import "strings"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func convert(s string, numRows int) string {
	N := len(s)
	if numRows < 2 || numRows > N {
		return s
	}
	sb := strings.Builder{}
	X := 0
	Y := 0
	index := 0

	matrix := make([][]string, numRows)
	for i := range matrix {
		matrix[i] = make([]string, N/2+1)
	}

	for index < N {
		for index < N && X != numRows-1 {
			matrix[X][Y] = string(s[index])
			X++
			index++
		}
		for index < N && X != 0 {
			matrix[X][Y] = string(s[index])
			X--
			Y++
			index++
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			sb.WriteString(matrix[i][j])
		}
	}
	return sb.String()
}

// 最优解
func convert2(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2 // 6
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}
