package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func subarraySum(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	// key: 前缀和 value: 出现了几次
	mp := make(map[int]int)
	// 重要：0位置默认出现了一次
	mp[0] = 1
	// 0-i的累加和
	all := 0
	// 累加和为k的个数
	ans := 0
	for i := 0; i < len(nums); i++ {
		all += nums[i]
		if v, ok := mp[all-k]; ok {
			ans += v
		}
		if v, ok := mp[all]; ok {
			mp[all] = v + 1
		} else {
			mp[all] = 1
		}
	}
	return ans
}
