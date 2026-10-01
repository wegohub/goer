package class06

import (
	. "backend/utils/algo"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	root := GenTreeWithWidthMarshal([]int{2, 1, 3})
	fmt.Println(isValidBST(root))
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
