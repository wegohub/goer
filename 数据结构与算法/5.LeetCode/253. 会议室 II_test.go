package leetcode

import (
	"container/heap"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type EndHeap [][]int

func (cls EndHeap) Len() int {
	return len(cls)
}

func (cls EndHeap) Less(i, j int) bool {
	return cls[i][1] < cls[j][1]
}

func (cls EndHeap) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (cls EndHeap) Peek() []int {
	return cls[0]
}

func (obj *EndHeap) Push(v interface{}) {
	*obj = append(*obj, v.([]int))
}

func (obj *EndHeap) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

func minMeetingRooms(intervals [][]int) int {
	// 先按照开始时间从早到晚排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// 结束时间小根堆
	h := &EndHeap{}
	heap.Init(h)
	//for _, item := range intervals {
	//	heap.Push(h, item)
	//}
	minCount := 0
	for i := 0; i < len(intervals); i++ {
		// 结束时间小于开始时间
		for h.Len() > 0 && h.Peek()[1] <= intervals[i][0] {
			heap.Pop(h)
		}
		heap.Push(h, intervals[i])
		minCount = Max(minCount, h.Len())
	}
	return minCount
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minMeetingRooms2(intervals [][]int) int {
	// 先按照开始时间从早到晚排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[0][i] < intervals[0][j]
	})
	timeLine := 0
	result := 0
	for i := 0; i < len(intervals); i++ {
		if timeLine <= intervals[i][0] {
			result++
			timeLine = intervals[i][1]
		}
	}
	return result
}
