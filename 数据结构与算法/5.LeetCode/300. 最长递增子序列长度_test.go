package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 最优解：N*Log(N)  dp的解法是：以每个字符作为结尾的情况，看前面比它小且最长的值+1，时间复杂度O(N^2)
func lengthOfLIS(arr []int) int {
	// ends[i]表示 : 目前所有长度为i+1的递增子序列的最小结尾
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	right := 0
	maxLen := 1

	for i := 1; i < len(arr); i++ {
		l := 0
		r := right

		for l <= r {
			m := (l + r) / 2
			if ends[m] >= arr[i] {
				r = m - 1
			} else {
				l = m + 1
			}
		}
		// l就是ends数组中arr[i]放的位置，l有可能越界出来，有可能在范围内
		right = max(right, l)
		ends[l] = arr[i]
		maxLen = max(maxLen, l+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
