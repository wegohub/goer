package class06

import (
	. "backend/utils/algo"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	root := GenTreeWithWidthMarshal([]int{5, 4, 8, 11, 13, 4, 7, 2, Nil, Nil, Nil, 1})
	PrintBinaryTree(root)
	fmt.Println(hasPathSum(root, 22))
	return "Hello World!", nil
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return process(root, 0, targetSum)
}

func process(root *TreeNode, currentSum, targetSum int) bool {
	if root.Left == nil && root.Right == nil {
		if currentSum+root.Val == targetSum {
			return true
		}
		return false
	}

	ans := false

	if root.Left != nil {
		p1 := process(root.Left, currentSum+root.Val, targetSum)
		if p1 {
			ans = true
		}
	}

	if root.Right != nil {
		p2 := process(root.Right, currentSum+root.Val, targetSum)
		if p2 {
			ans = true
		}
	}

	return ans
}
