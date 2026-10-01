package class11

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	Hanoi(3)
	return "Hello World!", nil
}

// Hanoi 汉诺塔问题
func Hanoi(n int) {
	process(n, "Left", "Right", "Mid")
}

func process(n int, from, to, other string) {
	if n == 1 {
		fmt.Println(fmt.Sprintf("Move %d from %s to %s", n, from, to))
		return
	}

	process(n-1, from, other, to)
	fmt.Println(fmt.Sprintf("Move %d from %s to %s", n, from, to))
	process(n-1, other, to, from)
}
