package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ans := minCost("horse", "ros", 1, 1, 1)
	fmt.Println(ans)
	return "Hello World!", nil
}

func minDistance(word1, word2 string) int {
	return minCost(word1, word2, 1, 1, 1)
}

// str1转换为str2最小的编辑代价
func minCost(str1, str2 string, ic, dc, rc int) int {
	if str1 == "" && str2 == "" {
		return 0
	}
	N := len(str1) + 1
	M := len(str2) + 1
	dp := make([][]int, N)
	for index := range dp {
		dp[index] = make([]int, M)
	}

	// 第一列 怎么把str1变成空串
	for i := 1; i < N; i++ {
		dp[i][0] = dc * i
	}

	// 第一行 怎么把空串变成str2
	for j := 1; j < M; j++ {
		dp[0][j] = ic * j
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			// 这里是下标，所以表示为：[0,i-1]
			// 最后一个字符一样的时候
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				// 最后一个字符不一样，str1->str2 承担一个rc代价
			} else {
				dp[i][j] = dp[i-1][j-1] + rc
			}
			// 前面的字符加一个ic代价， str1整体变成str2的前缀 + 插入str2最后一个字符
			// ab2ca => [abc2a]x
			// 让str1[i]变成str2[j-1]的前缀，添加一个x, +一个插入代价
			dp[i][j] = Min(dp[i][j], dp[i][j-1]+ic)
			// 后面的字符加一个dc代价，str1-1变成str2整体 + 删除str1最后一个字符
			// [ab2ca]f => abc2a  => 让str1[i-1]变成str2[j]的整体, 多出来一个f, +一个删除代价
			dp[i][j] = Min(dp[i][j], dp[i-1][j]+dc)
		}
	}
	fmt.Println(dp)
	return dp[N-1][M-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
