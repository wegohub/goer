package class06

import (
	. "backend/utils/algo"
	"container/list"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// Qsize方式收集
func levelOrderBottom(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	// 准备一个队列
	queue := list.New()
	queue.PushFront(root)

	for queue.Len() > 0 {
		size := queue.Len()
		cur := make([]int, size)
		for i := 0; i < size; i++ {
			node := queue.Back().Value.(*TreeNode)
			queue.Remove(queue.Back())
			cur[i] = node.Val

			if node.Left != nil {
				queue.PushFront(node.Left)
			}

			if node.Right != nil {
				queue.PushFront(node.Right)
			}
		}
		ans = append(ans, cur)
	}

	L := 0
	R := len(ans) - 1

	for L < R {
		ans[L], ans[R] = ans[R], ans[L]
		L++
		R--
	}
	return ans
}

// Flag方式收集
func levelOrderBottomFlag(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}

	// 准备一个队列
	queue := list.New()
	queue.PushBack(root)
	curEnd := root
	var nextEnd *TreeNode
	var tmp []int

	for queue.Len() > 0 {
		cur := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		if cur.Left != nil {
			queue.PushBack(cur.Left)
			nextEnd = cur.Left
		}

		if cur.Right != nil {
			queue.PushBack(cur.Right)
			nextEnd = cur.Right
		}

		tmp = append(tmp, cur.Val)
		// 收集答案
		if cur == curEnd {
			ans = append(ans, tmp)
			tmp = nil
			curEnd = nextEnd
			nextEnd = nil
		}
	}

	L := 0
	R := len(ans) - 1

	for L < R {
		ans[L], ans[R] = ans[R], ans[L]
		L++
		R--
	}
	return ans
}
