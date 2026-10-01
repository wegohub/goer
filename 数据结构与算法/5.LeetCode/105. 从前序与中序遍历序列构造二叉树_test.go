package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	dp := make(map[int]int)
	for index, item := range inorder {
		dp[item] = index
	}
	return process(preorder, inorder, 0, len(preorder)-1, 0, len(inorder), dp)
}

func process(pre []int, in []int, L1, R1, L2, R2 int, dp map[int]int) *TreeNode {
	if L1 > R1 {
		return nil
	}
	root := &TreeNode{Val: pre[L1]}
	if L1 == R1 {
		return root
	}
	find := dp[pre[L1]]
	root.Left = process(pre, in, L1+1, L1+find-L2, L2, find-1, dp)
	root.Right = process(pre, in, L1+find-L2+1, R1, find+1, R2, dp)

	return root
}

// 从中序与后序遍历序列构造二叉树
func buildTree1(inorder []int, postorder []int) *TreeNode {
	mp := make(map[int]int)
	for index, val := range inorder {
		mp[val] = index
	}
	return process1(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1, mp)

}

func process1(in []int, L1, R1 int, post []int, L2, R2 int, mp map[int]int) *TreeNode {
	if L2 > R2 {
		return nil
	}
	root := &TreeNode{Val: post[R2]}
	if L2 == R2 {
		return root
	}
	find := mp[post[R2]]
	root.Left = process1(in, L1, find-1, post, L2, R2-(R1-find)-1, mp)
	root.Right = process1(in, find+1, R1, post, R2-(R1-find), R2-1, mp)
	return root
}
