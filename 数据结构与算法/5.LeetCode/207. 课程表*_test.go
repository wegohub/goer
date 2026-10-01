package leetcode

import "container/list"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// Node 点描述
type Node struct {
	Value int     // 节点值
	In    int     // 入度
	Out   int     // 出度
	Nexts []*Node // 直接领居
	Edges []*Edge // 直接边
}

// Edge 边描述
type Edge struct {
	Weight int   // 权重
	From   *Node // 起始点
	To     *Node // 终点
}

// Graph 图描述
type Graph struct {
	Nodes map[int]*Node      // 点集(下标对于节点)
	Edge  map[*Edge]struct{} // 边集
}

func NewNode(value int) *Node {
	return &Node{
		Value: value,
		In:    0,
		Out:   0,
		Nexts: make([]*Node, 0),
		Edges: make([]*Edge, 0),
	}
}

func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[int]*Node),
		Edge:  make(map[*Edge]struct{}),
	}
}

func createGraph(prerequisites [][]int) *Graph {
	graph := NewGraph()

	for _, item := range prerequisites {
		from := item[1]
		to := item[0]
		if graph.Nodes[from] == nil {
			graph.Nodes[from] = NewNode(from)
		}
		if graph.Nodes[to] == nil {
			graph.Nodes[to] = NewNode(to)
		}

		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		edge := NewEdge(0, fromNode, toNode)

		fromNode.Out++
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Edges = append(fromNode.Edges, edge)
		toNode.In++

		graph.Edge[edge] = struct{}{}
	}

	return graph
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if len(prerequisites) == 0 {
		return true
	}

	graph := createGraph(prerequisites)

	queue := list.New()
	// 入度表
	inMap := make(map[*Node]int)

	for _, node := range graph.Nodes {
		inMap[node] = node.In
		if node.In == 0 {
			queue.PushBack(node)
		}
	}

	ans := 0
	for queue.Len() > 0 {
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())
		ans++

		for _, next := range cur.Nexts {
			inMap[next] = inMap[next] - 1
			if inMap[next] == 0 {
				queue.PushBack(next)
			}
		}
	}

	return ans == len(inMap)
}
