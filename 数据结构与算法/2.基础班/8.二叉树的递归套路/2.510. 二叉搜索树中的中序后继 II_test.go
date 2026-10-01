package class08

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func inorderSuccessor(node *Node) *Node {
	if node == nil {
		return nil
	}

	if node.Right != nil { // 有右树, 后继节点在右树上的最左节点
		node = node.Right
		for node.Left != nil {
			node = node.Left
		}
		return node
	} else { // 第一次出现当前节点是父节点的左孩子
		for node.Parent != nil && node != node.Parent.Left {
			node = node.Parent
		}

		if node.Parent == nil { // 该节点是整棵树上的最后一个节点
			return nil
		}
		// 返回当前节点的父级节点
		return node.Parent
	}
}
