package class02

import (
	"backend/utils/algo"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	head := algo.GenLink(10, 100)
	algo.PrintLink(head)

	arr := algo.GenTreeWidthMarshal(3, 10, 0.5)
	tree := algo.GenTreeWithWidthMarshal(arr)
	algo.PrintBinaryTree(tree)

	fmt.Println(algo.GenArray(10, 100))
	return "", nil
}

// GenArray 生成随机数组 [0,maxLen] 和 [0,maxVal]
func GenArray(maxLen, maxVal int) []int {
	length := int(rand.Float64() * float64(maxLen+1))
	ans := make([]int, length)
	for i := 0; i < length; i++ {
		ans[i] = int(rand.Float64() * float64(maxVal+1))
	}
	return ans
}

// ListNode 单链表结构
type ListNode struct {
	Val  int
	Next *ListNode
}

// GenLink 随机生成单链表
func GenLink(maxLen, maxVal int) *ListNode {
	length := int(rand.Float64() * (float64(maxLen + 1)))
	if length == 0 {
		return nil
	}
	// 把头节点建出来
	head := &ListNode{Val: int(rand.Float64() * (float64(maxVal + 1)))}
	// 建剩下的节点
	cur := head
	for i := 1; i < length; i++ {
		cur.Next = &ListNode{Val: int(rand.Float64() * (float64(maxVal + 1)))}
		cur = cur.Next
	}
	return head
}

// PrintLink 打印单链表
func PrintLink(head *ListNode) {
	builder := strings.Builder{}
	for head != nil {
		builder.WriteString(fmt.Sprintf("%d -> ", head.Val))
		head = head.Next
	}
	builder.WriteString("nil")
	fmt.Println(builder.String())
}

// TreeNode 二叉树节点定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PrintBinaryTree 打印二叉树
func PrintBinaryTree(root *TreeNode) {
	fmt.Println("#################数组转二叉树#################")
	defer func() {
		fmt.Println("##############################################")
	}()
	if root == nil {
		fmt.Println("空树")
		return
	}
	// 获取树的高度
	height := treeHeight(root) + 1
	// 计算最底层的宽度
	width := int(math.Pow(2, float64(height))) - 1
	// 初始化结果数组
	res := make([][]string, height)
	for i := range res {
		res[i] = make([]string, width)
		for j := range res[i] {
			res[i][j] = " "
		}
	}

	// 填充结果数组
	fill(res, root, 0, 0, width-1)

	// 打印结果数组
	for _, row := range res {
		for _, val := range row {
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func fill(res [][]string, node *TreeNode, level, left, right int) {
	mid := (left + right) / 2
	v := "x"
	if node != nil {
		v = strconv.Itoa(node.Val)
	}
	res[level][mid] = v
	if node == nil {
		return
	}
	fill(res, node.Left, level+1, left, mid-1)
	fill(res, node.Right, level+1, mid+1, right)
}

func treeHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftHeight := treeHeight(node.Left)
	rightHeight := treeHeight(node.Right)
	return max(leftHeight, rightHeight) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
