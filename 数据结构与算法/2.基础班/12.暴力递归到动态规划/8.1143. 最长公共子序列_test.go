package class12

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	text1 := "abcde"
	text2 := "ace"
	fmt.Println(longestCommonSubsequence(text1, text2))
	return "Hello World!", nil
}

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1))
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
	}

	// 填第一行
	for j := 0; j < len(text2); j++ {
		if text1[0] == text2[j] || (j > 0 && dp[0][j-1] == 1) {
			dp[0][j] = 1
		}
	}

	// 填第一列
	for i := 0; i < len(text1); i++ {
		if text2[0] == text1[i] || (i > 0 && dp[i-1][0] == 1) {
			dp[i][0] = 1
		}
	}

	// 普遍位置
	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			// 1. 即不以i位置的字符结尾，也不以j位置的字符结尾
			p1 := dp[i-1][j-1]

			// 2. 以i字符的结尾，不以j位置字符结果
			p2 := dp[i][j-1]

			// 3. 不以i位置的字符结尾，以j位置的字符结尾
			p3 := dp[i-1][j]

			// 4. 以i和j位置的字符结尾
			p4 := -1
			if text1[i] == text2[j] {
				p4 = dp[i-1][j-1] + 1
			}

			p12Max := math.Max(float64(p1), float64(p2))
			p34Max := math.Max(float64(p3), float64(p4))
			dp[i][j] = int(math.Max(p12Max, p34Max))

			// 化简后的结果
			//dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))

			//if text1[i] == text2[j] {
			//    dp[i][j] = int(math.Max(
			//        float64(dp[i][j]),
			//        float64(dp[i-1][j-1]+1),
			//    ))
			//}
		}
	}

	return dp[len(text1)-1][len(text2)-1]

}
