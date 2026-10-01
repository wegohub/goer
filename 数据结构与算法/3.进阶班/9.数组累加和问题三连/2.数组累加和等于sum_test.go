package class10

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxLength(arr []int, target int) int {
	if len(arr) == 0 {
		return 0
	}

	// key: 累加和 value: 最早出现的位置
	mp := make(map[int]int)
	// 重要: 0最早出现在-1位置
	mp[0] = -1

	maxLen := 0
	sum := 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if v, ok := mp[sum-target]; ok {
			// i - (v + 1) + 1 化简等于 i - v
			maxLen = int(math.Max(float64(maxLen), float64(i-v)))
		} else {
			mp[sum] = i
		}
	}

	return maxLen
}

func subarraySum(nums []int, k int) int {
	ans := 0
	sum := 0
	// key：和 value：次数
	mp := make(map[int]int)
	mp[0] = 1
	for _, num := range nums {
		sum += num
		if count, ok := mp[sum-k]; ok {
			ans += count
		}
		if v, ok := mp[sum]; ok {
			mp[sum] = v + 1
		} else {
			mp[sum] = 1
		}
	}
	return ans
}
