package leetcode

import (
	"container/list"
	"strconv"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func evalRPN(tokens []string) int {
	stack := list.New()
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			compute(stack, token)
		} else {
			num, _ := strconv.Atoi(token)
			stack.PushBack(num)
		}
	}
	return stack.Back().Value.(int)
}

func compute(stack *list.List, op string) {
	num2 := stack.Back().Value.(int)
	stack.Remove(stack.Back())
	num1 := stack.Back().Value.(int)
	stack.Remove(stack.Back())
	ans := 0
	switch op {
	case "+":
		ans = num1 + num2
	case "-":
		ans = num1 - num2
	case "*":
		ans = num1 * num2
	case "/":
		ans = num1 / num2
	}
	stack.PushBack(ans)
}
