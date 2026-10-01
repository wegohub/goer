package leetcode

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func sortedArrayToBST(nums []int) *TreeNode {
	return process(nums, 0, len(nums)-1)
}

func process(nums []int, L, R int) *TreeNode {
	if L > R {
		return nil
	}
	if L == R {
		return &TreeNode{Val: nums[L]}
	}

	M := (L + R) / 2
	head := &TreeNode{Val: nums[M]}
	head.Left = process(nums, L, M-1)
	head.Right = process(nums, M+1, R)
	return head
}

func sortedArrayToBST1(nums []int) *TreeNode {
	return processSortedArrayToBST(nums, 0, len(nums)-1)
}

func processSortedArrayToBST(nums []int, L, R int) *TreeNode {
	if L > R {
		return nil
	}
	M := (L + R) / 2
	root := &TreeNode{Val: nums[M]}
	root.Left = processSortedArrayToBST(nums, L, M-1)
	root.Right = processSortedArrayToBST(nums, M+1, R)
	return root
}
