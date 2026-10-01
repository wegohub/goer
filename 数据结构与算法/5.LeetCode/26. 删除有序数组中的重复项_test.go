package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 有序数组去重
func removeDuplicates(nums []int) int {
	done := 0
	cur := 1

	for cur < len(nums) {
		if nums[cur] != nums[done] {
			nums[cur], nums[done+1] = nums[done+1], nums[cur]
			done++
		}
		cur++
	}
	// 返回长度
	return done + 1
}
