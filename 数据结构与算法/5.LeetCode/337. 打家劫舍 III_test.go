package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Info struct {
	No  int // 整个子树，在不抢头节点的情况下，获得的最好收益
	Yes int // 整个子树，在抢头节点的情况下，获得的最好收益
}

func rob(root *TreeNode) int {
	info := process(root)
	return Max(info.No, info.Yes)
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{No: 0, Yes: 0}
	}
	leftInfo := process(root.Left)
	rightInfo := process(root.Right)

	// 不抢x的情况下，最好收益
	no := Max(leftInfo.No, leftInfo.Yes) + Max(rightInfo.No, rightInfo.Yes)

	// 抢x的情况下，最好收益
	yes := root.Val + leftInfo.No + rightInfo.No

	return &Info{No: no, Yes: yes}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
