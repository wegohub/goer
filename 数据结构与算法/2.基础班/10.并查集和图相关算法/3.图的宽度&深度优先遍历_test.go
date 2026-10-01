package class10

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	node1 := NewNode(1)
	node2 := NewNode(2)
	set := make(map[*Node]struct{})
	set[node1] = struct{}{}
	fmt.Println(set[node1])
	v, ok := set[node2]
	fmt.Println(v, ok)
	return "Hello World!", nil
}

// Node 节点描述
type Node struct {
	Value int
	In    int     // 入度(指向它的边条数)
	Out   int     // 出度(由它出发的边条数)，Nexts的size
	Nexts []*Node // 直接邻居，从它出发的点
	Edge  []*Edge // 直接边集，从它出发的边
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		In:    0,
		Out:   0,
		Nexts: make([]*Node, 0),
		Edge:  make([]*Edge, 0),
	}
}

// Edge 边描述
type Edge struct {
	Weight int   // 权重
	From   *Node //
	To     *Node
}

// NewEdge 图的推荐结构
func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

// Graph 图描述(点集和边集的描述)
type Graph struct {
	Nodes map[int]*Node      // 点的值到点的映射
	Edges map[*Edge]struct{} // 边集合
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edges: make(map[*Edge]struct{}),
	}
}

// BFS 图的宽度优先遍历
func BFS(node *Node) {
	if node == nil {
		return
	}

	queue := list.New()
	set := make(map[*Node]struct{})

	queue.PushBack(node)
	set[node] = struct{}{}

	for queue.Len() > 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		fmt.Println(cur.Value)

		for _, item := range cur.Nexts {
			if _, ok := set[item]; !ok {
				queue.PushBack(item)
				set[item] = struct{}{}
			}
		}
	}
}

// DFS 图的深度优先遍历
func DFS(node *Node) {
	if node == nil {
		return
	}
	stack := list.New() // 从栈底到栈顶表示的是遍历的path
	set := make(map[*Node]struct{})
	stack.PushFront(node)
	set[node] = struct{}{}
	fmt.Println(node.Value) // 入栈的时候打印
	for stack.Len() > 0 {
		cur := stack.Front().Value.(*Node)
		stack.Remove(stack.Front())
		for _, next := range cur.Nexts {
			if _, ok := set[next]; !ok {
				stack.PushFront(cur)
				stack.PushFront(next)
				set[next] = struct{}{}
				fmt.Println(next.Value)
				break // 去搞next这条支路
			} else {
				// cur 和 next节点存在引用关系
			}
		}
	}

}

// DFS 图的深度优先遍历，返回出现环路的节点列表
func DFS1(node *Node) ([]*Node, bool) {
	if node == nil {
		return nil, false
	}
	stack := list.New() // 从栈底到栈顶表示的是遍历的path
	set := make(map[*Node]struct{})
	cycleNodes := make([]*Node, 0)
	inStack := make(map[*Node]bool)

	stack.PushFront(node)
	set[node] = struct{}{}
	inStack[node] = true
	fmt.Println(node.Value) // 入栈的时候打印

	for stack.Len() > 0 {
		cur := stack.Front().Value.(*Node)
		stack.Remove(stack.Front())
		inStack[cur] = false

		for _, next := range cur.Nexts {
			if _, ok := set[next]; !ok {
				stack.PushFront(cur)
				stack.PushFront(next)
				set[next] = struct{}{}
				inStack[next] = true
				fmt.Println(next.Value)
				break // 去搞next这条支路
			} else if inStack[next] {
				// 如果next已经在栈中，说明存在环路
				cycleNodes = append(cycleNodes, next)
				return cycleNodes, true
			}
		}
	}

	return nil, false
}
