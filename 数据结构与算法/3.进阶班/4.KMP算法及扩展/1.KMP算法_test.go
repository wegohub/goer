package class04

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	str := "abcdeabcde"
	match := "cdeab"
	ans := IndexOf(str, match)
	fmt.Println(ans)
	return "Hello World!", nil
}

func IndexOf(str string, match string) int {
	if len(str) == 0 || len(match) == 0 || len(match) > len(str) {
		return -1
	}

	x := 0                      // str中当前比对的位置
	y := 0                      // mathch中当前比对的位置
	next := getNextArray(match) // match的next数组,

	// O(N)
	for x < len(str) && y < len(match) {
		if str[x] == match[y] { // x和y能匹配，一起往下走
			x++
			y++
		} else if next[y] == -1 { // y == 0
			x++
		} else {
			y = next[y] // 来到最长前缀的下一个位置，继续比对
		}
	}

	// 1. x越界，y没有越界， -1
	// 2. x没越界，y越界了，返回每一个开头

	if y == len(match) { // y越界了
		// 0 1 2 3 4 5 6 7
		// a a b a a b c d(x)
		// 0 1 2 3
		// a b c y
		// y 此时的位置表示，mathch字符的长度
		// x的位置-长度就是字符开始的位置
		return x - y
	}
	// y没有越界
	return -1
}

// 求next数组 O(M)
func getNextArray(match string) []int {
	// 0 => -1
	// 1 => 0
	// 2 => 0 == 1 ? 1 : 0
	// a b a s a b a t a b a  s  a  b  a  s  z
	// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16
	// 当前来到z(i), 看s(i-1)位置nex数组中的位置，t(7)
	// 入个t(7)位置的字符跟s(i-1)位置的字符相同，z(i)位置的字符next数组长度就是8
	// 如果不同，再看t(7)的前一个字符a(6)在next数组中的长度s(3), 3位置的字符跟s(i-1)的字符相同，i位置next数组中的长度是4
	if len(match) == 1 {
		return []int{-1}
	}
	next := make([]int, len(match))
	next[0] = -1 // 规定
	next[1] = 0  // 规定
	i := 2       // 从2位置出发求每个位置的值
	cn := 0      // 1. i-1位置的最长前后缀 2. i-1位置的字符跟cn位置的字符比较
	for i < len(next) {
		if match[i-1] == match[cn] { // 跳出来了
			next[i] = cn + 1
			i++
			cn++ // 被i位置使用，跳出来本来就是cn+1
		} else if cn > 0 { // 比对不成功，cn往前跳
			cn = next[cn]
		} else { // 跳到头了，i位置的最长公共前后缀就是0
			next[i] = 0
			i++
		}
	}

	return next
}
