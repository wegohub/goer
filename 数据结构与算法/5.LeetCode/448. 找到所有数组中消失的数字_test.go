package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 下标循环怼
func findDisappearedNumbers(nums []int) []int {
	ans := make([]int, 0)
	if len(nums) == 0 {
		return ans
	}
	N := len(nums)
	for i := 0; i < N; i++ {
		walk(nums, i)
	}

	// 遍历收集答案
	for i := 0; i < N; i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}

func walk(nums []int, i int) {
	// i 上已经躺在i+1了
	for nums[i] != i+1 {
		nexti := nums[i] - 1 // 下一个位置
		if nums[nexti] == nexti+1 {
			break
		}
		nums[i], nums[nexti] = nums[nexti], nums[i]
	}
}
