package leetcode

import (
	. "backend/utils/algo"
	"container/list"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// ZigZag打印 尾进头出，先左再右
func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	// 是否从左到右
	isHead := true
	for queue.Len() > 0 {
		var tmp []int
		size := queue.Len()
		if isHead { // 从左道右的过程，头出，尾巴进，先加左再加右
			for i := 0; i < size; i++ {
				cur := queue.Front().Value.(*TreeNode)
				queue.Remove(queue.Front())
				if cur.Left != nil {
					queue.PushBack(cur.Left)
				}
				if cur.Right != nil {
					queue.PushBack(cur.Right)
				}
				tmp = append(tmp, cur.Val)
			}
		} else { // 从右往左的过程，尾巴出，头进，先加右再加左
			for i := 0; i < size; i++ {
				cur := queue.Back().Value.(*TreeNode)
				queue.Remove(queue.Back())
				if cur.Right != nil {
					queue.PushFront(cur.Right)
				}
				if cur.Left != nil {
					queue.PushFront(cur.Left)
				}
				tmp = append(tmp, cur.Val)
			}
		}
		ans = append(ans, tmp)
		tmp = nil
		isHead = !isHead
	}

	return ans
}
