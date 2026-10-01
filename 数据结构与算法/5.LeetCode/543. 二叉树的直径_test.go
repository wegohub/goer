package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func diameterOfBinaryTree(root *TreeNode) int {
	return process(root).MaxDistance
}

type Info struct {
	MaxDistance int
	Height      int
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{MaxDistance: 0, Height: 0}
	}

	left := process(root.Left)
	right := process(root.Right)

	height := Max(left.Height, right.Height) + 1
	maxDistance := Max(Max(left.MaxDistance, right.MaxDistance), left.Height+right.Height)

	return &Info{
		Height:      height,
		MaxDistance: maxDistance,
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
