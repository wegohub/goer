package leetcode

import "container/list"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func isValid(s string) bool {
	// 配对串定义
	mp := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := list.New()
	for _, v := range s {
		// 左扩后压栈
		if v == '(' || v == '[' || v == '{' {
			stack.PushFront(v)
		} else { // 否则弹栈
			if stack.Len() == 0 {
				return false
			}
			// 需要配对的左括号
			match := mp[rune(v)]
			// 栈中弹出左括号
			cur := stack.Front().Value.(rune)
			stack.Remove(stack.Front())
			// 是否匹配
			if match != cur {
				return false
			}
		}
	}
	// 栈是否为空，数量要匹配
	return stack.Len() == 0
}
