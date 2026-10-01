package class10

import (
	"container/heap"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 图的输入
	matrix := [][]int{
		{1, 0, 1}, // 点变化
		{2, 0, 2},
		{100, 1, 4},
		{50, 1, 2},
		{8, 2, 3},
		{10, 3, 4},
	}
	graph := createGraph(matrix)
	ans, sum := Kruskal(graph)
	fmt.Println(ans)
	fmt.Println(sum)
	for _, item := range ans {
		fmt.Println(fmt.Sprintf("From %d to %d", item.From.Value, item.To.Value))
	}

	fmt.Println("================ P算法 ==================")
	graph1 := createGraph(matrix)
	ansP, sumP := PrimMST(graph1)
	fmt.Println(ansP)
	fmt.Println(sumP)
	for _, item := range ansP {
		fmt.Println(fmt.Sprintf("From %d to %d", item.From.Value, item.To.Value))
	}

	return "Hello World!", nil
}

func createGraph(matrix [][]int) *Graph {
	graph := NewGraph()
	for _, item := range matrix {
		weight := item[0]
		from := item[1]
		to := item[2]

		// 向图中添加点集
		if graph.Nodes[from] == nil {
			graph.Nodes[from] = NewNode(from)
		}
		if graph.Nodes[to] == nil {
			graph.Nodes[to] = NewNode(to)
		}

		formNode := graph.Nodes[from]
		toNode := graph.Nodes[to]
		edge := NewEdge(weight, graph.Nodes[from], graph.Nodes[to])

		formNode.Out++
		formNode.Nexts = append(formNode.Nexts, toNode)
		formNode.Edge = append(formNode.Edge, edge)
		toNode.In++

		// 向图中添加变集
		graph.Edges[edge] = struct{}{}
	}

	return graph
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

// ############################################### K算法 #########################################

// Heap 放所有边的小根堆
type Heap []*Edge

func (cls Heap) Len() int {
	return len(cls)
}

func (cls Heap) Less(i, j int) bool {
	return cls[i].Weight < cls[j].Weight
}

func (cls Heap) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (obj *Heap) Push(v interface{}) {
	*obj = append(*obj, v.(*Edge))
}

func (obj *Heap) Pop() interface{} {
	old := *obj
	n := len(*obj)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

func NewUionSet(nodes []*Node) *UnionSet {
	parents := make(map[*Node]*Node)
	sizeMap := make(map[*Node]int)
	for _, node := range nodes {
		parents[node] = node
		sizeMap[node] = 1
	}
	return &UnionSet{
		parents: parents,
		sizeMap: sizeMap,
	}
}

type UnionSet struct {
	parents map[*Node]*Node
	sizeMap map[*Node]int
}

func (obj *UnionSet) find(cur *Node) *Node {
	stack := make([]*Node, 0)
	index := 0
	for cur != obj.parents[cur] {
		stack = append(stack, cur)
		index++
		cur = obj.parents[cur]
	}

	// 扁平化
	for index--; index >= 0; index-- {
		obj.parents[stack[index]] = cur
	}

	return cur
}

func (obj *UnionSet) IsSameSet(a, b *Node) bool {
	if obj.parents[a] == nil || obj.parents[b] == nil {
		return false
	}
	return obj.find(a) == obj.find(b)
}

func (obj *UnionSet) Union(a, b *Node) {
	if obj.parents[a] == nil || obj.parents[b] == nil {
		return
	}
	aHead := obj.find(a)
	bHead := obj.find(b)
	if aHead != bHead {
		big := aHead
		small := bHead
		if obj.sizeMap[aHead] < obj.sizeMap[bHead] {
			big = bHead
			small = aHead
		}
		obj.parents[small] = big
		obj.sizeMap[big] += obj.sizeMap[small]
		delete(obj.sizeMap, small)
	}
}

// graph 简单一点是有向图, 需要准备的数据结构 小根堆、并查集、图
func Kruskal(graph *Graph) ([]*Edge, int) {
	edgeHeap := &Heap{}
	heap.Init(edgeHeap)
	for edge := range graph.Edges {
		heap.Push(edgeHeap, edge)
	}

	nodes := make([]*Node, 0)
	for _, node := range graph.Nodes {
		nodes = append(nodes, node)
	}
	set := NewUionSet(nodes)

	ans := make([]*Edge, 0)
	sum := 0
	for edgeHeap.Len() > 0 {
		edge := heap.Pop(edgeHeap).(*Edge)
		// fmt.Println(edge.Weight)
		if !set.IsSameSet(edge.From, edge.To) {
			set.Union(edge.From, edge.To)
			ans = append(ans, edge)
			sum += edge.Weight
		}
	}

	return ans, sum
}

// ############################################### P算法 #########################################

// 指定任意一个出发点
// 解锁一个点，点的直接边被解锁，从解锁边里面中选值最小的点，
// 如果是新点，再由2得到的点去解锁一批边
// 最小边数
// 点 -> 边 -> 点 -> 边

func PrimMST(graph *Graph) ([]*Edge, int) {
	// 准备一个小根堆，按照边权重最小值组织
	edgeHeap := &Heap{}
	heap.Init(edgeHeap)

	// 解锁点集合
	nodeSet := make(map[*Node]struct{})
	// 解锁边集合(已经进入的边不要重复考虑)
	edgeSet := make(map[*Edge]struct{})
	// 要了哪些边，结果集合
	result := make([]*Edge, 0)
	sum := 0

	// 随便挑一个点，防森林
	for _, node := range graph.Nodes {

		// 原出发点，第一个节点
		if _, ok := nodeSet[node]; !ok {
			nodeSet[node] = struct{}{}
			for _, edge := range node.Edge { // 由一个点解锁直接连接的边
				if _, ok1 := edgeSet[edge]; !ok1 {
					edgeSet[edge] = struct{}{}
					heap.Push(edgeHeap, edge)
				}
			}

			for edgeHeap.Len() > 0 {
				curEdge := heap.Pop(edgeHeap).(*Edge) // 弹出权重最小的边
				toNode := curEdge.To
				if _, ok2 := nodeSet[toNode]; !ok2 { // 一个新节点
					nodeSet[toNode] = struct{}{}
					result = append(result, curEdge)
					sum += curEdge.Weight
					for _, nextEdge := range toNode.Edge { // 再由这个新节点出发解锁一批边
						if _, ok3 := edgeSet[nextEdge]; !ok3 {
							edgeSet[nextEdge] = struct{}{}
							heap.Push(edgeHeap, nextEdge)
						}
					}
				}

			}

		}
	}

	return result, sum

}
