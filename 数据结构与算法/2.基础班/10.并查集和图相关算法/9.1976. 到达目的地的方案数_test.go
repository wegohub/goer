package class10

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func NewRNode(value int) *RNode {
	return &RNode{
		Value: value,
	}
}

type RNode struct {
	Value int
	In    int
	Out   int
	Next  []*RNode
	Edge  []*Edge
}

func NewEdge(weight int, from *RNode, to *RNode) *Edge {
	return &Edge{
		Weight: weight,
		From:   from,
		To:     to,
	}
}

type Edge struct {
	Weight int
	From   *RNode
	To     *RNode
}

type Graph struct {
	Nodes map[int]*RNode
	Edges map[*Edge]struct{}
}

func NewGraph(matrix [][]int) *Graph {
	nodes := make(map[int]*RNode)
	edges := make(map[*Edge]struct{})
	for _, item := range matrix {
		from := item[0]
		to := item[1]
		weight := item[2]
		if _, ok := nodes[from]; !ok {
			nodes[from] = NewRNode(from)
		}
		if _, ok := nodes[to]; !ok {
			nodes[to] = NewRNode(to)
		}
		fromNode := nodes[from]
		toNode := nodes[to]

		edge := NewEdge(weight, fromNode, toNode)

		fromNode.Next = append(fromNode.Next, toNode)
		fromNode.Edge = append(fromNode.Edge, edge)
		fromNode.Out++
		toNode.In++
		edges[edge] = struct{}{}
	}
	return &Graph{
		Nodes: nodes,
		Edges: edges,
	}
}

func countPaths(n int, roads [][]int) int {
	if len(roads) == 0 {
		return 0
	}
	if len(roads) == 1 {
		return 1
	}
	graph := NewGraph(roads)
	// 出发点
	from := graph.Nodes[0]
	// 目标点
	target := graph.Nodes[n-1]
	// 距离表
	distanceMap := make(map[*RNode]int)
	distanceMap[from] = 0
	// 已选择过的点集合
	selectedSet := make(map[*RNode]struct{})
	// 从原点出发距离最近的点
	minNode := getMinDistanceNode(distanceMap, selectedSet)

	// 0 - n-1 key: 距离 value：次数
	ans := make(map[int]int64)

	for minNode != nil {
		// 从原点出发到该点的距离
		distance := distanceMap[minNode]
		for _, edge := range minNode.Edge {

			toNode := edge.To
			if toNodeDistance, ok := distanceMap[toNode]; ok {
				distanceMap[toNode] = Min(toNodeDistance, distance+edge.Weight)
			} else {
				distanceMap[toNode] = distance + edge.Weight
			}
			// 收集一下答案
			if toNode == target || edge.From == target {
				//fmt.Println(edge.From.Value, edge.To.Value, "========")
				if _, ok := ans[distanceMap[target]]; ok {
					ans[distanceMap[target]] += 1
				} else {
					ans[distanceMap[target]] = 1
				}
			}
		}
		selectedSet[minNode] = struct{}{}
		minNode = getMinDistanceNode(distanceMap, selectedSet)
	}
	//fmt.Println(ans)
	if len(ans) == 0 {
		return 0
	}
	minKeys := make([]int, 0, len(ans))
	for index := range ans {
		minKeys = append(minKeys, index)
	}
	minKey := minKeys[0]
	for i := 1; i < len(minKeys); i++ {
		if minKeys[i] < minKey {
			minKey = minKeys[i]
		}
	}

	m := int64(math.Pow(10, 9)) + 7
	return int(ans[minKey] % m)
}

func getMinDistanceNode(distanceMap map[*RNode]int, selectedSet map[*RNode]struct{}) *RNode {
	var minNode *RNode
	minDistance := math.MaxInt
	for node, distance := range distanceMap {
		if _, ok := selectedSet[node]; !ok {
			if distance < minDistance {
				minNode = node
				minDistance = distance
			}
		}
	}
	return minNode
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
