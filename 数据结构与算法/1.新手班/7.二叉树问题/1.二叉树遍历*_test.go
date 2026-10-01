package class06

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	rootMarshal := GenTreeWidthMarshal(4, 100, 0.5)
	fmt.Println(rootMarshal)
	root := GenTreeWithWidthMarshal([]int{1, 2, 3, 4, 5, 6, 7})
	PrintBinaryTree(root)
	fmt.Println("========先序遍历-递归方式========")
	pre(root)
	fmt.Println("==============================")

	fmt.Println("========先序遍历-迭代方式========")
	preIteration(root)
	fmt.Println("==============================")

	fmt.Println("========中序遍历-递归方式========")
	in(root)
	fmt.Println("==============================")

	fmt.Println("========中序遍历-迭代方式========")
	inIteration(root)
	fmt.Println("==============================")

	fmt.Println("========中序遍历-迭代方式(易于理解)========")
	inIteration2(root)
	fmt.Println("==============================")

	fmt.Println("========后序遍历-递归方式========")
	post(root)
	fmt.Println("==============================")

	fmt.Println("========后序遍历-迭代方式========")
	postIteration(root)
	fmt.Println("==============================")
	return "Hello World!", nil
}

// 先序遍历-递归方式
func pre(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	pre(root.Left)
	pre(root.Right)
}

// 先序遍历-迭代
func preIteration(root *TreeNode) {
	if root == nil {
		return
	}
	// 准备一个栈
	stack := list.New()
	stack.PushFront(root)

	// 迭代栈
	for stack.Len() > 0 {
		node := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		fmt.Println(node.Val)

		if node.Right != nil {
			stack.PushFront(node.Right)
		}

		if node.Left != nil {
			stack.PushFront(node.Left)
		}
	}

}

// 中序遍历-递归方式
func in(root *TreeNode) {
	if root == nil {
		return
	}
	in(root.Left)
	fmt.Println(root.Val)
	in(root.Right)
}

// 中序遍历-迭代方式(用左边界将整个树分解掉)
func inIteration(root *TreeNode) {
	if root == nil {
		return
	}
	stack := list.New()
	for stack.Len() > 0 || root != nil {
		if root != nil {
			stack.PushFront(root)
			root = root.Left
		} else {
			root = stack.Front().Value.(*TreeNode)
			stack.Remove(stack.Front())
			fmt.Println(root.Val)
			root = root.Right
		}
	}
}

// 中序遍历-易于理解(用左边界将整个树分解掉)
func inIteration2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := list.New()

	decompositionLeftBound(root, stack)

	for stack.Len() > 0 {
		node := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		fmt.Println(node.Val)
		if node.Right != nil {
			decompositionLeftBound(node.Right, stack)
		}
	}
}

func decompositionLeftBound(root *TreeNode, stack *list.List) {
	for root != nil {
		stack.PushFront(root)
		root = root.Left
	}
}

// 后序遍历-递归方式
func post(root *TreeNode) {
	if root == nil {
		return
	}
	post(root.Left)
	post(root.Right)
	fmt.Println(root.Val)
}

// 后序遍历-迭代方式
func postIteration(root *TreeNode) {
	if root == nil {
		return
	}
	// 结果栈
	ans := list.New()
	// 迭代栈
	stack := list.New()
	stack.PushFront(root)

	for stack.Len() > 0 {
		node := stack.Front().Value.(*TreeNode)
		stack.Remove(stack.Front())
		ans.PushFront(node)

		if node.Left != nil {
			stack.PushFront(node.Left)
		}

		if node.Right != nil {
			stack.PushFront(node.Right)
		}
	}

	for ans.Len() > 0 {
		node := ans.Front().Value.(*TreeNode)
		ans.Remove(ans.Front())
		fmt.Println(node.Val)
	}
}
