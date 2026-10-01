package class10

import (
	"container/list"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
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

func sortedTopology(graph *Graph) []*Node {
	// 某一个点剩余的入度
	inMap := make(map[*Node]int)
	// 入度为0的队列
	zeroInQueue := list.New()

	// 初始化入度和队列
	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			zeroInQueue.PushBack(node)
		}
	}

	ans := make([]*Node, 0)
	for zeroInQueue.Len() > 0 {
		cur := zeroInQueue.Front().Value.(*Node)
		zeroInQueue.Remove(zeroInQueue.Front())
		ans = append(ans, cur)

		// cur的直接领居入度-1
		for _, next := range cur.Nexts {
			inMap[next] = inMap[next] - 1
			// 入度减为0，加入队列
			if inMap[next] == 0 {
				zeroInQueue.PushBack(next)
			}
		}
	}

	return ans
}
