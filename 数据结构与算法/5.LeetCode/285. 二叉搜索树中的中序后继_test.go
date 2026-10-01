package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	cur := root
	var mostRight *TreeNode
	var pre *TreeNode

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}

		if pre == p {
			return cur
		} else {
			pre = cur
		}
		cur = cur.Right
	}
	return nil
}
