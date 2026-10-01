package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func lengthOfLIS(arr []int) int {
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	right := 0
	maxLen := 1
	for i := 1; i < len(arr); i++ {
		L := 0
		R := right
		for L <= R {
			M := (L + R) / 2
			if ends[M] >= arr[i] {
				R = M - 1
			} else {
				L = M + 1
			}
		}

		right = Max(right, L)
		ends[L] = arr[i]
		maxLen = Max(maxLen, L+1)
	}
	return maxLen
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
