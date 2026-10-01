package leetcode

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	tree := GenTreeWithWidthMarshal([]int{1})
	PrintBinaryTree(tree)
	ans := widthOfBinaryTree(tree)
	fmt.Println(ans)

	ans2 := widthOfBinaryTree2(tree)
	fmt.Println(ans2)
	return "Hello World!", nil
}

type Payload struct {
	node  *TreeNode
	index int
}

// 最优解 时间复杂度O(N), 空间复杂度O(N), 优化常数时间
func widthOfBinaryTree2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	queue := list.New()
	queue.PushBack(&Payload{node: root, index: 0})

	for queue.Len() > 0 {
		// 结算当前层ans
		ans = int(math.Max(
			float64(queue.Back().Value.(*Payload).index-queue.Front().Value.(*Payload).index+1),
			float64(ans),
		))

		qsize := queue.Len()
		for i := 0; i < qsize; i++ {
			// 弹出节点
			cur := queue.Front().Value.(*Payload)
			queue.Remove(queue.Front())

			if cur.node.Left != nil {
				queue.PushBack(&Payload{node: cur.node.Left, index: 2*cur.index + 1})
			}

			if cur.node.Right != nil {
				queue.PushBack(&Payload{node: cur.node.Right, index: 2*cur.index + 2})
			}
		}
	}

	return ans
}

// 该方法超出时间限制
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	queue := list.New()
	hasNext := true
	queue.PushBack(root)
	for hasNext && queue.Len() > 0 {
		qsize := queue.Len()
		start := -1 // 当前层第一次节点不是空
		end := -1   // 当前层最后一次节点不是空
		hasNext = false
		for i := 0; i < qsize; i++ {
			cur, ok := queue.Front().Value.(*TreeNode)
			queue.Remove(queue.Front())

			if !ok { // 是空节点
				queue.PushBack(nil)
				queue.PushBack(nil)
				continue
			}

			// 当前层第一次节点不是空
			if start == -1 {
				start = i
			}
			// 当前层节点不为空就更新end
			end = i

			if cur.Left != nil {
				queue.PushBack(cur.Left)
				hasNext = true
			} else {
				queue.PushBack(nil)
			}

			if cur.Right != nil {
				queue.PushBack(cur.Right)
				hasNext = true
			} else {
				queue.PushBack(nil)
			}
		}
		// 一层收集完毕
		if start != -1 {
			ans = int(math.Max(float64(ans), float64(end-start+1)))
		}
	}

	return ans
}
