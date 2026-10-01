package class12

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	return "Hello World!", nil
}

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n == 0 || n == 1 {
		return n
	}

	// 准备一个map,记录当前字符上一次出现的位置
	lastIndexMap := make(map[byte]int)
	pre := -1 // i-1 位置往左推不动的位置
	ans := 0  // 收集结果

	for i := 0; i < n; i++ {
		if lastIndex, ok := lastIndexMap[s[i]]; ok {
			// i 位置和 i-1 位置能推到的最左记录
			// 经过这一步将pre从表示i-1位置能推的最远距离更新成了 i 位置能推的最远距离
			pre = int(math.Max(float64(pre), float64(lastIndex)))
		}
		cur := i - pre // 当前能推到最左边的长度
		ans = int(math.Max(float64(cur), float64(ans)))
		lastIndexMap[s[i]] = i
	}

	return ans
}
