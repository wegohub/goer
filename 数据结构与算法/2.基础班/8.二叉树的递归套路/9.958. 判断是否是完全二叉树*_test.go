package class08

import . "backend/utils/algo"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// isCompleteTree1 使用层序遍历判断是否为完全二叉树
func isCompleteTree1(head *TreeNode) bool {
	if head == nil {
		return true
	}
	queue := []*TreeNode{head}
	// 是否遇到过左右两个孩子不双全的节点
	leaf := false

	for len(queue) > 0 {
		head = queue[0]
		queue = queue[1:]
		left := head.Left
		right := head.Right
		//     必须为叶节点却有孩子                            有右无左
		if (leaf && (left != nil || right != nil)) || (left == nil && right != nil) {
			return false
		}

		if left != nil {
			queue = append(queue, left)
		}
		if right != nil {
			queue = append(queue, right)
		}
		if left == nil || right == nil {
			leaf = true
		}
	}
	return true
}

// Info 定义递归过程中的信息结构
type Info struct {
	IsFull bool
	IsCBT  bool
	Height int
}

// isCompleteTree2 使用递归判断是否为完全二叉树
func isCompleteTree2(head *TreeNode) bool {
	return process(head).IsCBT
}

// process 递归处理节点
func process(x *TreeNode) Info {
	if x == nil {
		return Info{IsFull: true, IsCBT: true, Height: 0}
	}

	leftInfo := process(x.Left)
	rightInfo := process(x.Right)

	height := max(leftInfo.Height, rightInfo.Height) + 1
	isFull := leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height
	isCBT := false

	if leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height { // 左树满 右树满 高度还一样
		isCBT = true
	} else if leftInfo.IsCBT && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 { // 左树是完全二叉树  右树是满二叉树  左树高度比右树高度大一个
		isCBT = true
	} else if leftInfo.IsFull && rightInfo.IsFull && leftInfo.Height == rightInfo.Height+1 { // 左树是满二叉树  右树是满二叉树  左树高度比右树高度大一个
		isCBT = true
	} else if leftInfo.IsFull && rightInfo.IsCBT && leftInfo.Height == rightInfo.Height { // 左树是满二叉树  右树是完全二叉树 左树高度比右树高度一样
		isCBT = true
	}

	return Info{IsFull: isFull, IsCBT: isCBT, Height: height}
}

// max 返回两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
