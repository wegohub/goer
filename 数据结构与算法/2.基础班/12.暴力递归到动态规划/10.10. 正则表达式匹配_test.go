package class12

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func isMatch(s string, p string) bool {
	// 将多个位置相同的*合并为一个
	p1 := make([]byte, 0)
	for i := 0; i < len(p); i++ {
		if p[i] == '*' && p[i-1] == '*' {
			continue
		} else {
			p1 = append(p1, p[i])
		}
	}
	p = string(p1)
	if isValidate(s, p) {
		//dp := make(map[string]bool)
		return process1(s, p, 0, 0)
	}
	return false
}

func isValidate(s, p string) bool {
	// s 里面不能含有 . || *
	for i := 0; i < len(s); i++ {
		if s[i] == '.' || s[i] == '*' {
			return false
		}
	}
	// p里面不能有两个**挨着 && p 第一个字符不能为*
	for i := 0; i < len(p); i++ {
		if p[i] == '*' && (i == 0 || p[i-1] == '*') {
			return false
		}
	}
	return true
}

// 潜台词：p[pIndex] != '*'
func process1(s string, p string, sIndex int, pIndex int) bool {
	// 原串来到末尾
	if sIndex == len(s) {
		// 在pIndex位置什么样的字符能配出空字符串
		// 1. 模式串也来到末尾
		if pIndex == len(p) {
			return true
		}
		// 2. a*b*c* isValidate中保证pIndex不会压中单个字符*
		// 模式串必须>=2个字符，且pIndex+1字符必须是*
		if pIndex+1 < len(p) && p[pIndex+1] == '*' {
			return process1(s, p, sIndex, pIndex+2)
		}
		// 否则配不出空字符
		return false
	}

	// 模式串来到终点位置
	if pIndex == len(p) {
		// 原串必须也来到终点位置
		return sIndex == len(s)
	}

	// sIndex 和 pIndex都没有来到终点位置

	// 1. pIndex + 1 不是 *，pIndex 和 sIndex的字符要对上
	// pIndex + 1 已经越界了或者pIndex+1位置不是*
	if pIndex+1 >= len(p) || p[pIndex+1] != '*' {
		return (s[sIndex] == p[pIndex] || p[pIndex] == '.') && process1(s, p, sIndex+1, pIndex+1)
	}

	// 2. pIndex + 1 是*
	// 2.1 sIndex 跟 pIndex不能配上，pIndex+1位置又是*，
	// 那么必须将pIndex和pIndex+1位置变成一个空字符，接着尝试pIndex+2的字符能不能与sIndex字符配上
	// s = abbc p = b*abbc
	if p[pIndex] != '.' && s[sIndex] != p[pIndex] {
		return process1(s, p, sIndex, pIndex+2)
	}

	// 2.2 sIndex跟pIndex能配上，且pIndex=='*'
	// s = aaabcd p = a*aaabcd => a*变0个a
	// s = aaabcd p = a*aabcd => a*变1个a
	// s = aaabcd p = a*abcd => a*变2个a
	// s = aaabcd p = a*bcd => a*变3个a
	// s = aaabcd p = a*bcd => a*变4个a => 配不上
	if process1(s, p, sIndex, pIndex+2) { // 0个a
		return true
	}

	for sIndex < len(s) && (s[sIndex] == p[pIndex] || p[pIndex] == '.') {
		if process1(s, p, sIndex+1, pIndex+2) { // 1个a 2个a 3个a
			return true
		}
		sIndex++
	}

	return false
}
