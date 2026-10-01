package class11

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 岛屿数量 O(N*M)
func islandNum(m [][]int) int {
	if len(m) == 0 || len(m[0]) == 0 {
		return 0
	}
	N := len(m)
	M := len(m[0])
	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if m[i][j] == 1 {
				ans++
				infect(m, i, j, N, M)
			}
		}
	}

	return ans
}

// 感染函数
// i j 当前行和列
// N M 总共多少行和列
func infect(m [][]int, i, j, N, M int) {
	if i < 0 || i >= N || j < 0 || j >= M || m[i][j] != 1 {
		return
	}

	m[i][j] = 2
	infect(m, i+1, j, N, M)
	infect(m, i-1, j, N, M)
	infect(m, i, j-1, N, M)
	infect(m, i, j+1, N, M)
}
