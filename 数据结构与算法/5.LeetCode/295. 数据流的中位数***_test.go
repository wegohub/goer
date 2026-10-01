package leetcode

import (
	"container/heap"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

type MedianFinder struct {
	minh *MinH
	maxh *MaxH
}

func Constructor() MedianFinder {
	return MedianFinder{
		minh: &MinH{},
		maxh: &MaxH{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	// 第一个数入大根堆
	if this.maxh.Len() == 0 {
		heap.Push(this.maxh, num)
		return
	}

	// 如果新进来的数 <= 大根堆的堆顶，入大根堆, 否则入小根堆
	if num <= this.maxh.Peek() {
		heap.Push(this.maxh, num)
	} else {
		heap.Push(this.minh, num)
	}

	// 大根堆超长
	if this.maxh.Len()-this.minh.Len() > 1 {
		top := heap.Pop(this.maxh)
		heap.Push(this.minh, top)
	}

	// 小根堆超长
	if this.minh.Len()-this.maxh.Len() > 1 {
		top := heap.Pop(this.minh)
		heap.Push(this.maxh, top)
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxh.Len() == 0 {
		return 0
	}
	// 偶数个
	if (this.maxh.Len()+this.minh.Len())%2 == 0 { // this.maxh.Len() == this.minh.Len()
		return float64(this.maxh.Peek()+this.minh.Peek()) / float64(2)
	}

	// 奇数个，返回较大的堆顶
	if this.maxh.Len() > this.minh.Len() {
		return float64(this.maxh.Peek())
	}

	return float64(this.minh.Peek())
}

// 大根堆
type MaxH []int

func (cls MaxH) Len() int {
	return len(cls)
}

func (cls MaxH) Less(i, j int) bool {
	return cls[i] > cls[j]
}

func (cls MaxH) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (cls MaxH) Peek() int {
	return cls[0]
}

func (obj *MaxH) Push(v interface{}) {
	*obj = append(*obj, v.(int))
}

func (obj *MaxH) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

// 小根堆
type MinH []int

func (cls MinH) Len() int {
	return len(cls)
}

func (cls MinH) Less(i, j int) bool {
	return cls[i] < cls[j]
}

func (cls MinH) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (cls MinH) Peek() int {
	return cls[0]
}

func (obj *MinH) Push(v interface{}) {
	*obj = append(*obj, v.(int))
}

func (obj *MinH) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}
