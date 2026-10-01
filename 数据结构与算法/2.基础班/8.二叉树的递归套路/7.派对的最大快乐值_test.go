package class08

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Employee struct {
	Happy int
	Nexts []*Employee
}

type Info struct {
	Yes int // 来
	No  int // 不来
}

// 最大快乐值
func maxHappy(root *Employee) int {
	info := process(root)
	return Max(info.Yes, info.No)
}

func process(root *Employee) *Info {
	if root == nil {
		return &Info{Yes: 0, No: 0}
	}

	no := 0
	yes := root.Happy
	for _, next := range root.Nexts {
		nextInfo := process(next)
		no += Max(nextInfo.Yes, nextInfo.No)
		yes += nextInfo.No
	}

	return &Info{Yes: yes, No: no}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
