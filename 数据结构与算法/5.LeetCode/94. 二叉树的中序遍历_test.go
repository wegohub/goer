package leetcode

import (
	. "backend/utils/algo"
	"container/list"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func inorderTraversal(head *TreeNode) []int {
	ans := make([]int, 0)
	if head != nil {
		stack := list.New()

		for stack.Len() > 0 || head != nil {
			if head != nil {
				stack.PushFront(head)
				head = head.Left
			} else {
				head = stack.Front().Value.(*TreeNode)
				stack.Remove(stack.Front())
				ans = append(ans, head.Val)
				head = head.Right
			}
		}
	}

	return ans
}

// 递归
func inorderTraversal1(head *TreeNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	processInorderTraversal1(head, &ans)
	return ans
}

func processInorderTraversal1(head *TreeNode, ans *[]int) {
	if head == nil {
		return
	}
	processInorderTraversal1(head.Left, ans)
	*ans = append(*ans, head.Val)
	processInorderTraversal1(head.Right, ans)
}

// 迭代
func inorderTraversal2(head *TreeNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	stack := list.New()
	for stack.Len() > 0 || head != nil {
		if head != nil {
			stack.PushBack(head)
			head = head.Left
		} else {
			head = stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
			ans = append(ans, head.Val)
			head = head.Right
		}
	}
	return ans
}

// morris
func inorderTraversal3(head *TreeNode) []int {
	ans := make([]int, 0)
	if head == nil {
		return ans
	}
	var mostRight *TreeNode
	cur := head
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

		ans = append(ans, cur.Val)

		cur = cur.Right
	}

	return ans
}
