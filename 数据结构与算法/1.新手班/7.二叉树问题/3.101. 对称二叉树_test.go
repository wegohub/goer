package class06

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	return "Hello World!", nil
}

func isSymmetric(root *TreeNode) bool {
	return checkIsSymmetric(root, root)
}

func checkIsSymmetric(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Val == q.Val && checkIsSymmetric(p.Left, q.Right) && checkIsSymmetric(p.Right, q.Left)
}
