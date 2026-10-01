package train

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 找左边比右边小的数的位置(第一个降序的位置)
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	if i >= 0 {
		j := n - 1
		// 找到刚刚比i位置大的数，[i+1, n-1] 是降序的, 找到上坡的坡底
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// 反转i+1, n-1
	L := i + 1
	R := n - 1
	for L < R {
		nums[L], nums[R] = nums[R], nums[L]
		R--
		L++
	}
}
