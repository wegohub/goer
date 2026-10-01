package class07

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := GenTreeWidthMarshal(3, 100, 0)
	root := GenTreeWithWidthMarshal(arr)
	PrintBinaryTree(root)
	in(root)
	return "Hello World!", nil
}

func in(head *TreeNode) {
	fmt.Println("二叉树中序遍历非递归版本")
	if head != nil {
		stack := list.New()

		for stack.Len() > 0 || head != nil {
			if head != nil {
				stack.PushFront(head)
				head = head.Left
			} else {
				head = stack.Front().Value.(*TreeNode)
				stack.Remove(stack.Front())
				fmt.Println(head.Val)
				head = head.Right
			}
		}
	}
}
