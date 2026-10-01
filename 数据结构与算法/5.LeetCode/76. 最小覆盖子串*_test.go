package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	// 准备一张欠款表
	mp := make([]int, 256)
	for _, item := range t {
		mp[item]++
	}
	// 总欠款长度
	all := len(t)
	// 窗口的左右边界
	L := 0
	R := 0
	// 最小长度(-1重来没有找到过)
	minLen := -1
	// 最小长度对应的左右边界
	ansL := -1
	ansR := -1

	// [L: R) R表示即将把R位置的数入窗口
	for R != len(s) {
		mp[s[R]]--
		if mp[s[R]] >= 0 {
			all--
		}
		// 欠款还完了
		if all == 0 {
			// L滑倒第一个不为0的数, 最多就到R(排除掉其余字符)
			for mp[s[L]] < 0 {
				mp[s[L]]++
				L++
			}
			// 更新答案
			if minLen == -1 || minLen > R-L+1 {
				minLen = R - L + 1
				ansL = L
				ansR = R
			}
			// 更新此时L得信息
			all++
			mp[s[L]]++
			// 当前位置的L已经结算了
			L++
		}
		R++
	}

	if minLen == -1 {
		return ""
	}
	return s[ansL : ansR+1]
}
