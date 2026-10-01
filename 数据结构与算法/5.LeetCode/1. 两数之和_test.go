package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func twoSum(nums []int, target int) []int {
	// key: 某一个数字 value: 在那个位置
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if v, ok := mp[target-nums[i]]; ok {
			return []int{v, i}
		}
		mp[nums[i]] = i
	}
	// 答案不存在返回[-1, -1]
	return []int{-1, -1}
}
