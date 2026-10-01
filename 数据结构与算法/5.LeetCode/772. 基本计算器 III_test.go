package leetcode

import (
	"container/list"
	"strconv"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func calculate(s string) int {
	ans, _ := f(s, 0)
	return ans
}

// 请从str[i...]往下算，遇到字符串终止位置或者右括号，就停止
// 返回两个值，长度为2的数组
// 0) 负责的这一段的结果是多少
// 1) 负责的这一段计算到了哪个位置
func f(str string, i int) (int, int) {
	stack := list.New()
	cur := 0
	// 遍历字符串，遇到字符串终止位置或者右括号，就停止
	for i < len(str) && str[i] != ')' {
		// 跳过空格
		if str[i] == ' ' {
			i++
			continue
		}
		// 遇到数字
		if str[i] >= '0' && str[i] <= '9' {
			cur = cur*10 + int(str[i]-'0')
			i++
			// 遇到了运算符
		} else if str[i] != '(' {
			addNum(stack, cur)
			stack.PushBack(string(str[i]))
			i++
			cur = 0
			// 遇到左括号了
		} else {
			res, index := f(str, i+1)
			cur = res
			i = index + 1
		}
	}
	// 最后一个数字还没放呢
	addNum(stack, cur)

	return getNum(stack), i
}

// 如果栈顶是乘除，把num和栈顶元素结算后压入栈中
func addNum(stack *list.List, num int) {
	if stack.Len() > 0 {
		top := stack.Back().Value.(string)
		stack.Remove(stack.Back())
		if top == "+" || top == "-" {
			stack.PushBack(top)
		} else {
			cur, _ := strconv.Atoi(stack.Back().Value.(string))
			stack.Remove(stack.Back())
			if top == "*" {
				num = cur * num
			} else {
				num = cur / num
			}
		}
	}
	stack.PushBack(strconv.Itoa(num))
}

// 结算栈中的元素, 从栈底到栈顶
func getNum(stack *list.List) int {
	res := 0
	isAdd := true
	cur := ""
	for stack.Len() > 0 {
		cur = stack.Front().Value.(string)
		stack.Remove(stack.Front())
		if cur == "+" {
			isAdd = true
		} else if cur == "-" {
			isAdd = false
		} else {
			num, _ := strconv.Atoi(cur)
			if isAdd {
				res += num
			} else {
				res -= num
			}
		}
	}
	return res
}
