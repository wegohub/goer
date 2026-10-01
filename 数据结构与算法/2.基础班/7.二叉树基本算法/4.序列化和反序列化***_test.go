package class07

import (
	. "backend/utils/algo"
	"container/list"
	"fmt"
	"strconv"
	"testing"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func pre(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	pre(root.Left)
	pre(root.Right)
}

func preMorris(root *TreeNode) {

	cur := root
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				fmt.Println(cur.Val)
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		} else {
			fmt.Println(cur.Val)
		}

		cur = cur.Right
	}

}

func preIteration(root *TreeNode) {
	stack := list.New()
	stack.PushBack(root)

	for stack.Len() > 0 {
		cur := stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())

		fmt.Println(cur.Val)

		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
	}

}

func Test_Pre(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	pre(root)
	fmt.Println("=========")
	preMorris(root)
	fmt.Println("=========")
	preIteration(root)
}

func in(root *TreeNode) {
	if root == nil {
		return
	}
	in(root.Left)
	fmt.Println(root.Val)
	in(root.Right)
}

func inMorris(root *TreeNode) {
	cur := root
	var mostRight *TreeNode
	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
			}
		}
		fmt.Println(cur.Val)
		cur = cur.Right
	}
}

func inIteration(root *TreeNode) {
	stack := list.New()
	for stack.Len() > 0 || root != nil {
		if root != nil {
			stack.PushBack(root)
			root = root.Left
		} else {
			root = stack.Back().Value.(*TreeNode)
			stack.Remove(stack.Back())
			fmt.Println(root.Val)
			root = root.Right
		}
	}
}

func Test_In(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	in(root)
	fmt.Println("=========")
	inMorris(root)
	fmt.Println("=========")
	inIteration(root)
}

func post(root *TreeNode) {
	if root == nil {
		return
	}
	post(root.Left)
	post(root.Right)
	fmt.Println(root.Val)
}

func postMorris(root *TreeNode) {
	cur := root
	var mostRight *TreeNode

	for cur != nil {
		mostRight = cur.Left
		if mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != cur {
				mostRight = mostRight.Right
			}
			if mostRight.Right == nil {
				mostRight.Right = cur
				cur = cur.Left
				continue
			} else {
				mostRight.Right = nil
				// 逆序打印左树的右边界
				printRight(cur.Left)
			}
		}
		cur = cur.Right
	}

	// 逆序打印整棵树的右边界
	printRight(root)
}

func printRight(root *TreeNode) {
	newHead := reverse(root)
	cur := newHead
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Right
	}
	reverse(newHead)
}

func reverse(root *TreeNode) *TreeNode {
	var pre1 *TreeNode
	for root != nil {
		next := root.Right
		root.Right = pre1
		pre1 = root
		root = next
	}
	return pre1
}

func postIteration(root *TreeNode) {
	stack := list.New()
	ans := list.New()

	stack.PushBack(root)
	for stack.Len() > 0 {
		cur := stack.Back().Value.(*TreeNode)
		stack.Remove(stack.Back())
		ans.PushBack(cur)

		if cur.Left != nil {
			stack.PushBack(cur.Left)
		}
		if cur.Right != nil {
			stack.PushBack(cur.Right)
		}
	}

	for ans.Len() > 0 {
		cur := ans.Back().Value.(*TreeNode)
		ans.Remove(ans.Back())
		fmt.Println(cur.Val)
	}
}

func Test_Post(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	post(root)
	fmt.Println("=========")
	postMorris(root)
	fmt.Println("===========")
	postIteration(root)
}

func preEncode(root *TreeNode, ans *[]string) {
	if root == nil {
		*ans = append(*ans, "null")
		return
	}
	*ans = append(*ans, strconv.Itoa(root.Val))
	preEncode(root.Left, ans)
	preEncode(root.Right, ans)
}

func preDecode(ans []string) *TreeNode {
	root, _ := preDecodeProcess(ans, 0)
	return root
}

// 返回头和消费到的位置
func preDecodeProcess(ans []string, index int) (*TreeNode, int) {
	// 这里递归不会跑不完
	if ans[index] == "null" {
		return nil, index + 1
	}
	v, _ := strconv.Atoi(ans[index])
	root := &TreeNode{Val: v}
	next := index + 1
	root.Left, next = preDecodeProcess(ans, next)
	root.Right, next = preDecodeProcess(ans, next)
	return root, next
}

func Test_preEncode(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	ans := make([]string, 0)
	preEncode(root, &ans)
	fmt.Println(ans)
	fmt.Println("================")
	tree := preDecode(ans)
	PrintBinaryTree(tree)
}

func postEncode(root *TreeNode, ans *[]string) {
	if root == nil {
		*ans = append(*ans, "null")
		return
	}
	postEncode(root.Left, ans)
	postEncode(root.Right, ans)
	*ans = append(*ans, strconv.Itoa(root.Val))
}

func postDecode(ans []string) *TreeNode {
	// 左右头 -> 头右左
	stack := list.New()
	for _, item := range ans {
		stack.PushBack(item)
	}
	return postDecodeProcess(stack)
}

func postDecodeProcess(stack *list.List) *TreeNode {
	cur := stack.Back().Value.(string)
	stack.Remove(stack.Back())
	if cur == "null" {
		return nil
	}
	v, _ := strconv.Atoi(cur)
	root := &TreeNode{Val: v}
	root.Right = postDecodeProcess(stack)
	root.Left = postDecodeProcess(stack)
	return root
}

func Test_postDecode(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	ans := make([]string, 0)
	postEncode(root, &ans)
	fmt.Println(ans)
	fmt.Println("================")
	tree := postDecode(ans)
	PrintBinaryTree(tree)
}

func layerEncode(root *TreeNode, ans *[]string) {
	queue := list.New()
	queue.PushBack(root)
	*ans = append(*ans, strconv.Itoa(root.Val))
	for queue.Len() > 0 {
		cur := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		if cur.Left != nil {
			queue.PushBack(cur.Left)
			*ans = append(*ans, strconv.Itoa(cur.Left.Val))
		} else {
			*ans = append(*ans, "null")
		}

		if cur.Right != nil {
			queue.PushBack(cur.Right)
			*ans = append(*ans, strconv.Itoa(cur.Right.Val))
		} else {
			*ans = append(*ans, "null")
		}
	}
}

func layerDecode(ans []string) *TreeNode {
	ansQ := list.New()
	for _, item := range ans {
		ansQ.PushBack(item)
	}
	queue := list.New()
	root := genTreeNode(ansQ)
	if root == nil {
		return nil
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		cur := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		cur.Left = genTreeNode(ansQ)
		if cur.Left != nil {
			queue.PushBack(cur.Left)
		}

		cur.Right = genTreeNode(ansQ)
		if cur.Right != nil {
			queue.PushBack(cur.Right)
		}
	}
	return root
}

func genTreeNode(ans *list.List) *TreeNode {
	if ans.Len() == 0 {
		return nil
	}
	cur := ans.Front().Value.(string)
	ans.Remove(ans.Front())
	if cur == "null" {
		return nil
	}
	v, _ := strconv.Atoi(cur)
	return &TreeNode{
		Val: v,
	}
}

func Test_layerEncode(t *testing.T) {
	root := GenTreeWithWidthMarshal(GenTreeWidthMarshal(3, 100, 0.5))
	PrintBinaryTree(root)
	ans := make([]string, 0)
	layerEncode(root, &ans)
	fmt.Println(ans)
	fmt.Println("================")
	tree := layerDecode(ans)
	PrintBinaryTree(tree)
}
