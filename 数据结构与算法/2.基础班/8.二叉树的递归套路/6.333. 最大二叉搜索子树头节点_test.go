package class08

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Info struct {
	IsBST      bool
	MinVal     int
	MaxVal     int
	MaxBSTHead *TreeNode
	MaxBSTSize int
	Size       int // 整棵树的节点数
}

func largestBSTSubtree(root *TreeNode) int {
	return process(root).MaxBSTSize
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			IsBST:      true,
			MinVal:     math.MaxInt,
			MaxVal:     math.MinInt,
			MaxBSTHead: nil,
			MaxBSTSize: 0,
			Size:       0,
		}
	}

	left := process(root.Left)
	right := process(root.Right)

	minVal := Min(Min(left.MinVal, right.MinVal), root.Val)
	maxVal := Max(Max(left.MaxVal, right.MaxVal), root.Val)
	isBST := true
	if !left.IsBST {
		isBST = false
	}
	if !right.IsBST {
		isBST = false
	}
	if root.Val <= left.MaxVal || root.Val >= right.MinVal {
		isBST = false
	}
	size := left.Size + right.Size + 1
	maxBSTSize := left.MaxBSTSize
	maxBSTHead := left.MaxBSTHead
	if right.MaxBSTSize > left.MaxBSTSize {
		maxBSTSize = right.MaxBSTSize
		maxBSTHead = right.MaxBSTHead
	}
	if isBST {
		maxBSTSize = size
		maxBSTHead = root
	}

	return &Info{
		IsBST:      isBST,
		MinVal:     minVal,
		MaxVal:     maxVal,
		MaxBSTHead: maxBSTHead,
		MaxBSTSize: maxBSTSize,
		Size:       size,
	}

}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
