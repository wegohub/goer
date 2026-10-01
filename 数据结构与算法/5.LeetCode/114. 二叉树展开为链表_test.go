package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// flatten flattens the binary tree to a linked list in-place.
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	var pre *TreeNode = nil
	cur := root
	var mostRight *TreeNode = nil
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				if pre != nil {
					pre.Left = cur
				}
				pre = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			if pre != nil {
				pre.Left = cur
			}
			pre = cur
		}
		cur = cur.Right
	}
	cur = root
	var next *TreeNode = nil
	for cur != nil {
		next = cur.Left
		cur.Left = nil
		cur.Right = next
		cur = next
	}
}
