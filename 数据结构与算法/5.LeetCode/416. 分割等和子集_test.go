package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func canPartition(nums []int) bool {
	N := len(nums)
	sum := 0
	for i := 0; i < N; i++ {
		sum += nums[i]
	}
	if sum&1 != 0 {
		return false
	}
	sum >>= 1
	dp := make([][]bool, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < N; i++ {
		dp[i][0] = true
	}
	if nums[0] <= sum {
		dp[0][nums[0]] = true
	}
	for i := 1; i < N; i++ {
		for j := 1; j <= sum; j++ {
			dp[i][j] = dp[i-1][j]
			if j-nums[i] >= 0 {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i]]
			}
		}
		if dp[i][sum] {
			return true
		}
	}
	return false
}
