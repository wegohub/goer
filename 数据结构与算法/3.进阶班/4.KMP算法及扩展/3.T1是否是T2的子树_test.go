package class04

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// big做头节点的树，其中是否有某棵子树的结构，是和small为头的树，完全一样的
func containsTree(big *TreeNode, small *TreeNode) bool {
	if small == nil {
		return true
	}

	if big == nil {
		return false
	}

	if isSameValueStructure(big, small) {
		return true
	}

	return containsTree(big.Left, small) || containsTree(big.Right, small)
}

func isSameValueStructure(head1 *TreeNode, head2 *TreeNode) bool {
	if head1 == nil && head2 != nil {
		return false
	}

	if head1 != nil && head2 == nil {
		return false
	}

	if head1 == nil && head2 == nil {
		return true
	}

	if head1.Val != head2.Val {
		return false
	}

	return isSameValueStructure(head1.Left, head2.Left) && isSameValueStructure(head1.Right, head2.Right)
}

// 优化，先序方式序列化big和small， kmp判断 small是big的子串
