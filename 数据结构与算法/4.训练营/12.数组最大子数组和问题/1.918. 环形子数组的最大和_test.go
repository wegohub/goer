package sub_arr

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxSubarraySumCircular(nums []int) int {
	N := len(nums)
	// 环形数组最大子数组的和
	ans := nums[0]
	// 当前位置最大子数组和
	pre := nums[0]
	// 从左往右前缀和
	leftSum := nums[0]
	// 从左往右最大前缀和子数组
	leftMax := make([]int, N)
	leftMax[0] = nums[0]
	for i := 1; i < N; i++ {
		pre = Max(nums[i], nums[i]+pre)
		ans = Max(ans, pre)
		leftSum += nums[i]
		leftMax[i] = Max(leftSum, leftMax[i-1])
	}

	// 枚举每个后缀
	// 从右往左累加和
	rightSum := 0
	for i := N - 1; i > 0; i-- {
		rightSum += nums[i]
		ans = Max(ans, rightSum+leftMax[i-1])
	}

	return ans
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
