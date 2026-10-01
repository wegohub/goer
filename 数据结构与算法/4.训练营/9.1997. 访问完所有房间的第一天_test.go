package train

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func firstDayBeenInAllRooms(nextVisit []int) int {
	mod := int(1e9 + 7)
	n := len(nextVisit)
	dp := make([]int, n)
	// [X X j X X X] i (i+1)
	//           i穿越回j花费1天   从j回到i花费天数   来到i+1话费天数
	// dp[i+1] = (dp[i] + 1)  +  (dp[i]-dp[j])  + 1
	for i := 0; i < n-1; i++ {
		j := nextVisit[i]
		//                         减法有可能出现负数
		dp[i+1] = (dp[i]+1+dp[i]-dp[j]+mod)%mod + 1
	}
	return dp[n-1] % mod
}
