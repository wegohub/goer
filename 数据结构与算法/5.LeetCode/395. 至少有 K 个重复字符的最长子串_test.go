package leetcode

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 滑动窗口的方法 O(26*N)
func longestSubstring2(s string, k int) int {
	str := []rune(s)
	N := len(str)
	max := 0
	// 尝试每一个种类
	for require := 1; require <= 26; require++ {
		// a~z  a~z 出现次数
		// count[0  1  2]  a b c
		count := make([]int, 26)
		// 目前窗口内收集了几种字符了
		collect := 0
		// 目前窗口内出现次数>=k次的字符，满足了几种
		satisfy := 0
		// 窗口右边界
		R := -1
		for L := 0; L < N; L++ { // L要尝试每一个窗口的最左位置
			// [L..R]  R+1      字符种数不够 或者 后面的字符是出现过的，不会增加种数 停！
			//for R+1 < N && !(collect == require && count[str[R+1]-'a'] == 0) {
			for R+1 < N && (collect != require || count[str[R+1]-'a'] != 0) {
				R++
				// 之前没出现过，种类++
				if count[str[R]-'a'] == 0 {
					collect++
				}
				// 次数达标，达标种类++
				if count[str[R]-'a'] == k-1 {
					satisfy++
				}
				// 增加词频
				count[str[R]-'a']++
			}
			// [L...R] 全部达标
			if satisfy == require {
				max = int(math.Max(float64(max), float64(R-L+1)))
			}
			// L位置出窗口，清理词频，尝试下一个字符能不能让R往右扩
			if count[str[L]-'a'] == 1 {
				collect--
			}
			if count[str[L]-'a'] == k {
				satisfy--
			}
			count[str[L]-'a']--
		}
	}
	return max
}
