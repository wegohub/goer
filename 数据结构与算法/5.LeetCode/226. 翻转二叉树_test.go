package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(left)
	return root
}
