package leetcode

import (
	. "backend/utils/algo"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	//if root.Left == nil && root.Right == nil {
	//    return 1
	//}
	return Max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
