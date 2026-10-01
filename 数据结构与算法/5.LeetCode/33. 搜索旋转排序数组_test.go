package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 这个解是考虑了值相同情况的
func search(nums []int, target int) int {
	L := 0
	R := len(nums) - 1
	for L <= R {
		M := (L + R) / 2
		if nums[M] == target {
			return M
		}
		// nums[M] != target
		// [2 2 2 2 2 2 2 2 2 1 2 2 2 2] L M R 都相等，L一直走到不相等的位置，
		// 如果走到M都相等，那么 [M+1, R] 继续二分
		if nums[L] == nums[M] && nums[R] == nums[M] {
			for L != M && nums[L] == nums[M] {
				L++
			}
			if L == M {
				L = M + 1
				continue
			}
		}

		// L M R 不都一样的情况可以二分
		if nums[L] != nums[M] { // L != M
			if nums[M] > nums[L] {
				if target >= nums[L] && target < nums[M] {
					R = M - 1
				} else {
					L = M + 1
				}
			} else {
				if target > nums[M] && target <= nums[R] {
					L = M + 1
				} else {
					R = M - 1
				}
			}
		} else { // L == M => M != R
			if nums[M] < nums[R] {
				if target > nums[M] && target <= nums[R] {
					L = M + 1
				} else {
					R = M - 1
				}
			} else {
				if target >= nums[L] && target < nums[M] {
					R = M - 1
				} else {
					L = M + 1
				}
			}
		}
	}

	return -1
}
