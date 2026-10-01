package class06

import (
	. "backend/utils/algo"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	marshal := GenTreeWidthMarshal(10, 100, 0.5)
	root := GenTreeWithWidthMarshal(marshal)
	fmt.Println(maxDepth(root))
	return "Hello World!", nil
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := maxDepth(root.Left)
	rifhtHeight := maxDepth(root.Right)
	return int(math.Max(float64(leftHeight), float64(rifhtHeight))) + 1
}
