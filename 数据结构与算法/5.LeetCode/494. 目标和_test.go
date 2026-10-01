package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 暴力递归
func findTargetSumWays1(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	return process(nums, 0, target)
}

func process(nums []int, index int, rest int) int {
	if index == len(nums) {
		if rest == 0 {
			return 1
		}
		return 0
	}
	// -
	p1 := process(nums, index+1, rest+nums[index])
	// +
	p2 := process(nums, index+1, rest-nums[index])
	return p1 + p2
}

func findTargetSumWays(nums []int, target int) int {
	if len(nums) == 0 {
		if target == 0 {
			return 1
		}
		return 0
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || target < -sum {
		return 0
	}
	N := len(nums)
	dp := make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, 2*sum+1)
	}
	// 0          100
	// -100       0
	// f(index, rest) => dp[index][sum+rest]
	dp[N][0+sum] = 1

	for index := N - 1; index >= 0; index-- {
		for rest := -sum; rest <= sum; rest++ {
			// -
			dp[index][sum+rest] = 0
			if rest+nums[index]+sum <= 2*sum {
				dp[index][sum+rest] += dp[index+1][rest+nums[index]+sum]
			}
			// +
			if rest-nums[index]+sum >= 0 {
				dp[index][sum+rest] += dp[index+1][rest-nums[index]+sum]
			}
		}
	}

	return dp[0][target+sum]
}
