package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func missingNumber(nums []int) int {
	L := 0
	R := len(nums)
	for L < R {
		if nums[L] == L {
			L++
		} else if nums[L] < L || nums[L] >= R || nums[nums[L]] == nums[L] {
			// 发货到垃圾区
			// nums[L], nums[R-1] = nums[R-1], nums[L]
			R--
			nums[L] = nums[R]
		} else { // nums[nums[L]-1] != nums[L]
			nums[L], nums[nums[L]] = nums[nums[L]], nums[L]
		}
	}
	return L
}

func missingNumber2(nums []int) (xor int) {
	for i, num := range nums {
		xor ^= i ^ num
	}
	return xor ^ len(nums)
}
