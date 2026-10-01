package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func partition(s string) [][]string {
	// 动态规划辅助结构
	dp := getDP(s)
	path := make([]string, 0)
	ans := make([][]string, 0)
	process(s, 0, path, &ans, dp)
	return ans
}

// s[L] == s[R] && dp[L+1][R-1]是回文
func getDP(s string) [][]bool {
	N := len(s)
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, N)
	}
	for i := 0; i < N-1; i++ {
		dp[i][i] = true
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}
	// 右下角那一格
	dp[N-1][N-1] = true
	// 从第2列开始填对角线
	for j := 2; j < N; j++ {
		row := 0
		col := j
		for row < N && col < N {
			if s[row] == s[col] && dp[row+1][col-1] {
				dp[row][col] = true
			}
			row++
			col++
		}
	}
	return dp
}

// 深度优先遍历(DFS)
func process(s string, index int, path []string, ans *[][]string, dp [][]bool) {
	if index == len(s) {
		tmp := make([]string, 0, len(path))
		for _, item := range path {
			tmp = append(tmp, item)
		}
		*ans = append(*ans, tmp)
	} else {
		for end := index; end < len(s); end++ {
			if dp[index][end] { // [index, end]位置是回文，剩下的字符去后续的流程做决定
				path = append(path, s[index:end+1])
				process(s, end+1, path, ans, dp)
				path = path[0 : len(path)-1]
			}
		}
	}
}
