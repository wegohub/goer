package class06

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func Morris(head *TreeNode) {
	if head == nil {
		return
	}

	cur := head
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil { // 有左树
			// 左数上最右节点
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			// 第一次来到mostRight
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else { // mostRight.Right == cur, 第二次来到mostRight
				mostRight.Right = nil
			}
		}

		// 没有左树
		cur = cur.Right
	}
}

// morris加工出先序
func MorrisPre(head *TreeNode) {
	if head == nil {
		return
	}

	cur := head
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		// 有左树
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}

			if mostRight.Right == nil {
				mostRight.Right = cur
				// 有左树的节点第一次来到自己
				fmt.Println(cur.Val)
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			// 没有左树的节点，第一次来到自己
			fmt.Println(cur.Val)
		}

		// 没有左树
		cur = cur.Right
	}
}

// morris加工出中序, 第二次来到自己的时候打印
func MorrisIn(head *TreeNode) {
	if head == nil {
		return
	}

	cur := head
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		// 有左树
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

		// 没有左树
		// 第二次来到自己打印，中序遍历
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

// morris加工出后序
// 第二次来到自己的时候第二次来到自己打印，逆序打印左树的右边界
// 最后打印整棵树的右边界
func MorrisOut(head *TreeNode) {
	if head == nil {
		return
	}

	cur := head
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		// 有左树
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
				// 第二次来到自己打印，逆序打印左树的右边界
				printMostRight(cur.Left)
			}
		}

		// 没有左树
		cur = cur.Right
	}

	// 打印整棵树的右边界
	printMostRight(head)
}

// 打印右边界
func printMostRight(head *TreeNode) {
	tail := reverse(head)
	cur := tail
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Right
	}
	reverse(tail)
}

func reverse(head *TreeNode) *TreeNode {
	var pre *TreeNode
	cur := head
	for cur != nil {
		next := cur.Right
		cur.Right = pre
		pre = cur
		cur = next
	}
	return pre
}
