package class08

import (
	. "backend/utils/algo"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func IsFBT(root *TreeNode) bool {
	info := process(root)
	return int(math.Pow(2, float64(info.height)))-1 == info.nodeNum
}

type Info struct {
	height  int // 高度
	nodeNum int // 节点数量
}

func process(root *TreeNode) *Info {
	if root == nil {
		return &Info{
			height:  0,
			nodeNum: 0,
		}
	}

	left := process(root.Left)
	right := process(root.Right)

	height := Max(left.height, right.height) + 1
	nodeNum := left.nodeNum + right.nodeNum + 1

	return &Info{
		height:  height,
		nodeNum: nodeNum,
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
