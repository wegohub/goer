package class06

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func isBalanced(root *TreeNode) bool {
	return process(root).isBalanced
}

type Info struct {
	isBalanced bool
	height     int
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			isBalanced: true,
			height:     0,
		}
	}

	left := process(root.Left)
	right := process(root.Right)

	isBalanced := left.isBalanced && right.isBalanced && math.Abs(float64(left.height)-float64(right.height)) < 2
	height := Max(left.height, right.height) + 1
	return &Info{
		isBalanced: isBalanced,
		height:     height,
	}

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
