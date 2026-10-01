package class10

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	matrix := [][]int{
		{1, 0, 1}, // 点变化
		{2, 0, 2},
		{100, 1, 4},
		{50, 1, 2},
		{8, 2, 3},
		{10, 3, 4},
	}
	graph := createGraph(matrix)
	from := graph.Nodes[0]

	distanceMap := Dijkstra(from)
	for node, distance := range distanceMap {
		fmt.Println(fmt.Sprintf("From %d to %d distance is %d", from.Value, node.Value, distance))
	}

	fmt.Println("======================= 堆优化 =====================")
	graph2 := createGraph(matrix)
	from2 := graph2.Nodes[0]
	size := len(graph2.Nodes)
	distanceMap1 := Dijkstra2(from2, size)
	for node, distance := range distanceMap1 {
		fmt.Println(fmt.Sprintf("From %d to %d distance is %d", from2.Value, node.Value, distance))
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

func Dijkstra(from *Node) map[*Node]int {
	// 准备一个距离表
	distanceMap := make(map[*Node]int)
	distanceMap[from] = 0
	//准备一个已经算过距离的集合，算过距离的点就跳过了
	selectedSet := make(map[*Node]struct{})

	// 获取当前距离表中最近的节点
	minNode := getMinDistanceAndUnSelectedNode(distanceMap, selectedSet)
	for minNode != nil {
		distance := distanceMap[minNode]
		// 从这个点解锁所有的直接边
		for _, edge := range minNode.Edge {
			toNode := edge.To
			// 如果toNode已经计算过距离，那么就是求已存在的距离和当前点到toNode的距离的最小值
			if toDistance, ok := distanceMap[toNode]; ok {
				distanceMap[toNode] = int(math.Min(
					float64(toDistance),
					float64(distance+edge.Weight),
				))
			} else { // toNode没有计算过距离就是当前点到toNode的距离
				distanceMap[toNode] = distance + edge.Weight
			}
		}
		// 标记minNode已经计算过距离了，后续的流程都不会再碰这个点了
		selectedSet[minNode] = struct{}{}
		minNode = getMinDistanceAndUnSelectedNode(distanceMap, selectedSet)
	}

	return distanceMap
}

func getMinDistanceAndUnSelectedNode(distanceMap map[*Node]int, seleselectedSet map[*Node]struct{}) *Node {
	var minNode *Node
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		if _, ok := seleselectedSet[node]; !ok {
			if distance < minDistance {
				minNode = node
				minDistance = distance
			}
		}
	}
	return minNode
}

// ################################### Dijkstra算法用堆改写优化 #####################################

func NewHeap(cap int) *Heap {
	return &Heap{
		nodes:       make([]*Node, cap),
		indexMap:    make(map[*Node]int),
		distanceMap: make(map[*Node]int),
		size:        0,
	}
}

type Heap struct {
	nodes       []*Node       // 堆结果
	indexMap    map[*Node]int // 节点在堆上的位置，-1表示曾经存在后来弹出了
	distanceMap map[*Node]int // 从源节点到该节点的最小距离
	size        int           // 堆上有多少个节点, 表示新增节点的插入位置
}

// IsEmpty 堆是否是空的
func (obj *Heap) IsEmpty() bool {
	return obj.size == 0
}

// IsEntered 某个节点是否进来过，当前在堆上的+已经弹出的
func (obj *Heap) IsEntered(node *Node) bool {
	if _, ok := obj.indexMap[node]; ok {
		return true
	}
	return false
}

// InHeap 判断一个节点是否在堆上
func (obj *Heap) InHeap(node *Node) bool {
	if index, ok := obj.indexMap[node]; ok && index != -1 {
		return true
	}
	return false
}

// AddOrUpdateOrIgnore 节点node，现在从源节点到它的最小距离是distance
// 在该方法中要判断是否需要更新
func (obj *Heap) AddOrUpdateOrIgnore(node *Node, distance int) {
	// 如果这个节点在堆上，就需要更新当前的距离和老距离的最小值
	if obj.InHeap(node) {
		obj.distanceMap[node] = int(math.Min(
			float64(obj.distanceMap[node]),
			float64(distance),
		))
		// 距离可能会变小了，向上做heapInsert
		obj.heapInsert(obj.indexMap[node])
	}
	// 如果这个节点压根就没进过堆
	if !obj.IsEntered(node) {
		obj.nodes[obj.size] = node
		obj.indexMap[node] = obj.size
		obj.distanceMap[node] = distance
		// 向上做heapInsert调整
		obj.heapInsert(obj.size)
		obj.size++
	}
	// 已经弹出的节点就ignore了
}

// Pop 返回堆顶元素和源节点到它的距离（堆排序的问题）
func (obj *Heap) Pop() (*Node, int) {
	node := obj.nodes[0]
	distance := obj.distanceMap[node]

	obj.indexMap[node] = -1
	delete(obj.distanceMap, node)
	obj.swap(0, obj.size-1)
	obj.nodes[obj.size-1] = nil
	// 从上往下调整堆
	obj.size--
	obj.heapify(0)
	return node, distance
}

func (obj *Heap) heapInsert(index int) {
	for obj.distanceMap[obj.nodes[index]] < obj.distanceMap[obj.nodes[(index-1)/2]] {
		obj.swap(index, (index-1)/2)
		index = (index - 1) / 2
	}
}

func (obj *Heap) heapify(index int) {
	left := 2*index + 1
	for left < obj.size {
		smallest := left
		if left+1 < obj.size && obj.distanceMap[obj.nodes[left+1]] < obj.distanceMap[obj.nodes[left]] {
			smallest = left + 1
		}
		if obj.distanceMap[obj.nodes[index]] < obj.distanceMap[obj.nodes[smallest]] {
			smallest = index
		}
		if smallest == index {
			break
		}

		obj.swap(index, smallest)
		index = smallest
		left = 2*index + 1
	}
}

func (obj *Heap) swap(i, j int) {
	obj.indexMap[obj.nodes[i]] = j
	obj.indexMap[obj.nodes[j]] = i
	obj.nodes[i], obj.nodes[j] = obj.nodes[j], obj.nodes[i]
}

func Dijkstra2(from *Node, cap int) map[*Node]int {
	heap := NewHeap(cap)
	heap.AddOrUpdateOrIgnore(from, 0)
	distanceMap := make(map[*Node]int)
	for !heap.IsEmpty() {
		curNode, distance := heap.Pop()
		for _, edge := range curNode.Edge {
			heap.AddOrUpdateOrIgnore(edge.To, edge.Weight+distance)
		}
		distanceMap[curNode] = distance
	}
	return distanceMap
}
