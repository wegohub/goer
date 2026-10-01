package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// kthSmallest 找到二叉搜索树中第 k 小的元素
func kthSmallest(head *TreeNode, k int) int {
	if head == nil {
		return -1
	}

	cur := head
	var mostRight *TreeNode
	index := 1

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
		if index == k {
			return cur.Val
		}
		index++
		cur = cur.Right
	}
	return -1
}
