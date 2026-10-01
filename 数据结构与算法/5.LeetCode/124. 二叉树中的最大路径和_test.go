package leetcode

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Info struct {
	MaxPathSum         int
	MaxPathSumFromHead int
}

func maxPathSum(root *TreeNode) int {
	return process(root).MaxPathSum
}

func process(x *TreeNode) *Info {
	if x == nil {
		return nil
	}
	left := process(x.Left)
	right := process(x.Right)

	p1 := math.MinInt
	if left != nil {
		p1 = left.MaxPathSum
	}

	p2 := math.MinInt
	if right != nil {
		p2 = right.MaxPathSum
	}

	p3 := x.Val

	p4 := math.MinInt
	if left != nil {
		p4 = x.Val + left.MaxPathSumFromHead
	}

	p5 := math.MinInt
	if right != nil {
		p5 = x.Val + right.MaxPathSumFromHead
	}

	p6 := math.MinInt
	if left != nil && right != nil {
		p6 = x.Val + left.MaxPathSumFromHead + right.MaxPathSumFromHead
	}

	maxSum := Max(p1, p2, p3, p4, p5, p6)
	maxSumFormHead := Max(p3, p4, p5)

	return &Info{
		MaxPathSum:         maxSum,
		MaxPathSumFromHead: maxSumFormHead,
	}

}

func Max(item ...int) int {
	max := item[0]
	for i := 1; i < len(item); i++ {
		if item[i] > max {
			max = item[i]
		}
	}
	return max
}
