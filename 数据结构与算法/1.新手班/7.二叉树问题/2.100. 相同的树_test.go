package class06

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	p := GenTreeWithWidthMarshal([]int{1, 2, 3})
	q := GenTreeWithWidthMarshal([]int{1, 2, 3})
	fmt.Println(isSameTree(p, q))
	return "Hello World!", nil
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (p != nil && q == nil) {
		return false
	}
	if p == nil && q == nil {
		return true
	}

	pArr := getInArr(p)
	qArr := getInArr(q)

	if len(pArr) != len(qArr) {
		return false
	}

	for i := 0; i < len(pArr); i++ {
		if pArr[i] != qArr[i] {
			return false
		}
	}

	return true
}

func getInArr(root *TreeNode) []int {
	arr := make([]int, 0)
	stack := list.New()
	stack.PushFront(root)
	for stack.Len() > 0 {
		node, ok := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		if node == nil || !ok {
			arr = append(arr, math.MaxInt)
			continue
		}
		arr = append(arr, node.Val)

		stack.PushFront(node.Right)
		stack.PushFront(node.Left)
	}
	return arr
}
