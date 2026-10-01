package leetcode

import "container/heap"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// Node represents a node in the matrix with a value and its coordinates.
type Node struct {
	value int
	row   int
	col   int
}

// NodeHeap is a min-heap of Nodes.
type NodeHeap []Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push adds a node to the heap.
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

// Pop removes and returns the node with the minimum value from the heap.
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// kthSmallest1 finds the k-th smallest element in the matrix using a heap.
// 小根堆依次放下方和右方的数
// 弹出一个，放入两个，并不重复放
// 弹出k个就是答案
func kthSmallest1(matrix [][]int, k int) int {
	N := len(matrix)
	M := len(matrix[0])
	heapNodes := &NodeHeap{}
	heap.Init(heapNodes)
	set := make([][]bool, N)
	for i := range set {
		set[i] = make([]bool, M)
	}
	heap.Push(heapNodes, Node{matrix[0][0], 0, 0})
	set[0][0] = true
	count := 0
	var ans Node
	for heapNodes.Len() > 0 {
		ans = heap.Pop(heapNodes).(Node)
		count++
		if count == k {
			break
		}
		row := ans.row
		col := ans.col
		if row+1 < N && !set[row+1][col] {
			heap.Push(heapNodes, Node{matrix[row+1][col], row + 1, col})
			set[row+1][col] = true
		}
		if col+1 < M && !set[row][col+1] {
			heap.Push(heapNodes, Node{matrix[row][col+1], row, col + 1})
			set[row][col+1] = true
		}
	}
	return ans.value
}
