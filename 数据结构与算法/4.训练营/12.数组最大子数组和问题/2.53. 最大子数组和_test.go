package sub_arr

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxSubArray(nums []int) int {
	cur := nums[0]
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		cur = Max(nums[i], cur+nums[i])
		maxSum = Max(maxSum, cur)
	}
	return maxSum
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
