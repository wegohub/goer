package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	process(nums, 0, &ans)
	return ans
}

func process(nums []int, index int, ans *[][]int) {
	if index == len(nums) {
		cur := make([]int, len(nums))
		for index, item := range nums {
			cur[index] = item
		}
		*ans = append(*ans, cur)
		return
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		process(nums, index+1, ans)
		nums[index], nums[i] = nums[i], nums[index]
	}
}
