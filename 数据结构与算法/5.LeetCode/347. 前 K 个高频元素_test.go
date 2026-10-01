package leetcode

import "container/heap"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// Node represents a node in the frequency map.
type Node struct {
	num   int
	count int
}

// CountComparator is a type for comparing nodes based on their count.
type CountComparator []Node

func (c CountComparator) Len() int { return len(c) }
func (c CountComparator) Less(i, j int) bool {
	return c[i].count < c[j].count
}
func (c CountComparator) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Push adds a node to the heap.
func (c *CountComparator) Push(x interface{}) {
	*c = append(*c, x.(Node))
}

// Pop removes and returns the node with the minimum count from the heap.
func (c *CountComparator) Pop() interface{} {
	old := *c
	n := len(old)
	x := old[n-1]
	*c = old[0 : n-1]
	return x
}

// 堆+记账本方法： 时间复杂度O(N*LogK)
// 排序时间复杂度是：O(N*LogN)
// topKFrequent returns the top k frequent elements from the array nums.
func topKFrequent(nums []int, k int) []int {
	// 统计词频
	mapNode := make(map[int]*Node)
	for _, num := range nums {
		if node, found := mapNode[num]; !found {
			mapNode[num] = &Node{num: num, count: 1}
		} else {
			node.count++
		}
	}

	// 门槛堆 将复杂度从LogN降低到LogK
	heapNodes := &CountComparator{}
	heap.Init(heapNodes)

	// 门槛堆的作用，如果你都干不掉此时最小的元素，那么你一定不行
	for _, node := range mapNode {
		// 堆未满 或者 堆满且次数大于堆顶 入堆
		if heapNodes.Len() < k || (heapNodes.Len() == k && node.count > (*heapNodes)[0].count) {
			heap.Push(heapNodes, *node)
		}
		// 弹出当前的最小值
		if heapNodes.Len() > k {
			heap.Pop(heapNodes)
		}
	}

	// 最好收集答案
	ans := make([]int, k)
	index := 0
	for heapNodes.Len() > 0 {
		ans[index] = heap.Pop(heapNodes).(Node).num
		index++
	}

	return ans
}
