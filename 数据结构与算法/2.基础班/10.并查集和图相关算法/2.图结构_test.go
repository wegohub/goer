package class10

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

func createGraph(matrix [][]int) *Graph {
	graph := NewGraph()
	for _, item := range matrix {

		// 拿出原始数据中边描述
		weight := item[0]
		from := item[1]
		to := item[2]

		// 点集
		if graph.Nodes[from] == nil {
			graph.Nodes[from] = NewNode(from)
		}
		if graph.Nodes[to] == nil {
			graph.Nodes[to] = NewNode(to)
		}

		fromNode := graph.Nodes[from]
		toNode := graph.Nodes[to]

		edge := NewEdge(weight, fromNode, toNode)
		fromNode.Nexts = append(fromNode.Nexts, toNode)
		fromNode.Out++
		fromNode.Edge = append(fromNode.Edge, edge)
		toNode.In++
		graph.Edges[edge] = struct{}{}
	}

	return graph
}
