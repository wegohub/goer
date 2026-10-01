package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func firstMissingPositive(nums []int) int {
	L := 0
	R := len(nums)
	for L < R {
		if nums[L] == L+1 {
			L++
		} else if nums[L] <= L || nums[L] > R || nums[nums[L]-1] == nums[L] {
			// 发货到垃圾区
			// nums[L], nums[R-1] = nums[R-1], nums[L]
			nums[L] = nums[R-1]
			R--
		} else { // nums[nums[L]-1] != nums[L]
			nums[L], nums[nums[L]-1] = nums[nums[L]-1], nums[L]
		}
	}
	return L + 1
}
