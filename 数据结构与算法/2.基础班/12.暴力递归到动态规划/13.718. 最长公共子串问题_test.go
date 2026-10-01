package class12

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func findLength(nums1 []int, nums2 []int) int {
	N := len(nums1)
	M := len(nums2)

	dp := make([][]int, N)
	for index := range dp {
		dp[index] = make([]int, M)
	}

	ans := 0

	// 第一行
	for j := 0; j < M; j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
			ans = 1
		}
	}

	// 第一列
	for i := 0; i < N; i++ {
		if nums2[0] == nums1[i] {
			dp[i][0] = 1
			ans = 1
		}
	}

	for i := 1; i < N; i++ {
		for j := 1; j < M; j++ {
			// nums1 必须以i位置的数作为公共子数组
			// nums2 必须以j位置的数作为公共子数组
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
				ans = Max(ans, dp[i][j])
			}
		}
	}

	return ans
}
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
