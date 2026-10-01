package class04

import (
	"container/heap"
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 堆测试
	h := &Heap{}
	heap.Init(h)
	heap.Push(h, 4)
	heap.Push(h, 3)
	heap.Push(h, 5)
	heap.Push(h, 1)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h).(int))
	}

	arr := []int{5, 2, 1, 8, 67, 9, 5}
	KSort(arr, 3)
	fmt.Println(arr)

	return "Hello World!", nil
}

func KSort(arr []int, k int) {
	n := len(arr)
	if n < 2 {
		return
	}

	h := &Heap{}
	heap.Init(h)

	// 将k个数先入小根堆
	for i := 0; i <= int(math.Min(float64(k), float64(n-1))); i++ {
		heap.Push(h, arr[i])
	}

	index := 0
	for i := k + 1; i < n; i++ {
		arr[index] = heap.Pop(h).(int)
		heap.Push(h, arr[i])
		index++
	}

	for h.Len() > 0 {
		arr[index] = heap.Pop(h).(int)
		index++
	}

}

type Heap []int

func (cls Heap) Len() int {
	return len(cls)
}

func (cls Heap) Less(i, j int) bool {
	return cls[i] < cls[j]
}

func (cls Heap) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (obj *Heap) Push(v interface{}) {
	*obj = append(*obj, v.(int))
}

func (obj *Heap) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}
