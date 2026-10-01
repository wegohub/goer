package leetcode

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func isValidBST(root *TreeNode) bool {
	info := process(root)
	if info == nil {
		return true
	}
	return info.isBST
}

type Info struct {
	isBST  bool
	maxVal int
	minVal int
}

func process(root *TreeNode) *Info {
	if root == nil {
		return nil
	}

	left := process(root.Left)
	right := process(root.Right)
	isBST := true
	if left != nil && !left.isBST {
		isBST = false
	}
	if right != nil && !right.isBST {
		isBST = false
	}
	if left != nil && root.Val <= left.maxVal {
		isBST = false
	}
	if right != nil && root.Val >= right.minVal {
		isBST = false
	}

	maxVal := root.Val
	minVal := root.Val
	if left != nil && right == nil {
		maxVal = int(math.Max(float64(left.maxVal), float64(maxVal)))
		minVal = int(math.Min(float64(left.minVal), float64(minVal)))
	}
	if left == nil && right != nil {
		maxVal = int(math.Max(float64(right.maxVal), float64(maxVal)))
		minVal = int(math.Min(float64(right.minVal), float64(minVal)))
	}
	if left != nil && right != nil {
		maxVal = int(math.Max(float64(maxVal), math.Max(float64(left.maxVal), float64(right.maxVal))))
		minVal = int(math.Min(float64(minVal), math.Min(float64(left.minVal), float64(right.minVal))))
	}

	return &Info{
		isBST:  isBST,
		maxVal: maxVal,
		minVal: minVal,
	}
}

func isValidBST2(root *TreeNode) bool {
	info := processBST(root)
	if info == nil {
		return true
	}
	return info.IsBST
}

type InfoBST struct {
	IsBST bool
	Min   int
	Max   int
}

func processBST(root *TreeNode) *InfoBST {
	if root == nil {
		return nil
	}
	left := processBST(root.Left)
	right := processBST(root.Right)

	minVal := root.Val
	if left != nil && left.Min < minVal {
		minVal = left.Min
	}
	if right != nil && right.Min < minVal {
		minVal = right.Min
	}

	maxVal := root.Val
	if left != nil && left.Max > maxVal {
		maxVal = left.Max
	}
	if right != nil && right.Max > maxVal {
		maxVal = right.Max
	}

	isBST := true
	if left != nil && (!left.IsBST || root.Val <= left.Max) {
		isBST = false
	}
	if right != nil && (!right.IsBST || root.Val >= right.Min) {
		isBST = false
	}

	return &InfoBST{
		IsBST: isBST,
		Max:   maxVal,
		Min:   minVal,
	}

}
