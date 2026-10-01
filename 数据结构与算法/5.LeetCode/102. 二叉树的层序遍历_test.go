package leetcode

import (
	. "backend/utils/algo"
	"container/list"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 这个是flag方式的解法，还有qsize和map的解法
func levelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	// 准备一个队列
	queue := list.New()
	queue.PushBack(root)
	// 当前层结束
	curEnd := root
	// 下一层结束
	var nextEnd *TreeNode

	// 每一层的答案
	var tmp []int

	for queue.Len() > 0 {
		cur := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())
		tmp = append(tmp, cur.Val)

		if cur.Left != nil {
			queue.PushBack(cur.Left)
			nextEnd = cur.Left
		}

		if cur.Right != nil {
			queue.PushBack(cur.Right)
			nextEnd = cur.Right
		}

		// 结算当前层答案
		if cur == curEnd {
			ans = append(ans, tmp)
			tmp = nil
			curEnd = nextEnd
			nextEnd = nil
		}
	}

	return ans
}
