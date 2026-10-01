package class06

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	root := GenTreeWithWidthMarshal([]int{5, 4, 8, 11, Nil, 13, 4, 7, 2, Nil, Nil, 5, 1})
	PrintBinaryTree(root)
	fmt.Println(pathSum(root, 22))
	return "Hello World!", nil
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	l := list.New()
	process(root, targetSum, 0, make([]int, 0), l)
	for l.Len() > 0 {
		item := l.Front().Value.([]int)
		l.Remove(l.Front())
		ans = append(ans, item)
	}
	return ans
}

func process(root *TreeNode, targetSum int, currentSum int, path []int, ans *list.List) {
	if root.Left == nil && root.Right == nil {
		if currentSum+root.Val == targetSum {
			path = append(path, root.Val)
			ans.PushBack(copy1(path))
			path = path[0 : len(path)-1]
		}
		return
	}

	path = append(path, root.Val)
	if root.Left != nil {
		process(root.Left, targetSum, currentSum+root.Val, path, ans)
	}
	if root.Right != nil {
		process(root.Right, targetSum, currentSum+root.Val, path, ans)
	}
	path = path[0 : len(path)-1]
}

func copy1(arr []int) []int {
	ans := make([]int, len(arr))
	for index, item := range arr {
		ans[index] = item
	}
	return ans
}

// coding重构
func pathSum2(root *TreeNode, targetSum int) [][]int {
	ans := make([][]int, 0)
	if root == nil {
		return ans
	}
	path := make([]int, 0)
	process2(root, targetSum, 0, path, &ans)
	return ans
}

func process2(root *TreeNode, targetSum int, curSum int, path []int, ans *[][]int) {
	// 去重
	if root.Left == nil && root.Right == nil {
		if targetSum == curSum+root.Val {
			tmp := make([]int, len(path))
			copy(tmp, path)
			tmp = append(tmp, root.Val)
			*ans = append(*ans, tmp)
		}
		return
	}
	path = append(path, root.Val)
	if root.Left != nil {
		process2(root.Left, targetSum, root.Val+curSum, path, ans)
	}
	if root.Right != nil {
		process2(root.Right, targetSum, root.Val+curSum, path, ans)
	}
	path = path[0 : len(path)-1]
}
