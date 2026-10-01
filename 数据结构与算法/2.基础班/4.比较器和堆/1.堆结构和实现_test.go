package class04

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	// 测试堆排序
	fmt.Println("测试堆排序开始")
	arr := []int{6, 3, 2, 6, 1, 7, 8, 5}
	HeapSort(arr)
	fmt.Println(arr)
	fmt.Println("测试堆排序结束")
	fmt.Println("======================")
	h1 := NewHeap(20, []int{45, 89, 32, 54})
	h1.Push(6)
	h1.Push(5)
	h1.Push(4)
	h1.Push(3)
	h1.Pop()
	h1.Push(2)
	h1.Push(1)
	for !h1.IsEmpty() {
		fmt.Println(h1.Pop())
	}

	return "Hello World!", nil
}

func NewHeap(maxSize int, init ...[]int) *Heap {
	data := make([]int, maxSize)
	size := 0
	// init 用于实现 O(n) 的建堆方法
	if len(init) > 0 && len(init[0]) > 0 {
		initData := init[0]
		if len(initData) > maxSize {
			panic("init more max size")
		}
		// 建堆 O(n)
		for i := len(initData) - 1; i >= 0; i-- {
			heapify(initData, i, len(initData))
		}
		// 拷贝
		for index, item := range initData {
			data[index] = item
			size++
		}

	}
	return &Heap{
		data:    data,
		maxSize: maxSize,
		size:    size,
	}
}

type Heap struct {
	data    []int
	maxSize int
	size    int
}

func (obj *Heap) Size() int {
	return obj.size
}

func (obj *Heap) IsEmpty() bool {
	return obj.size == 0
}

func (obj *Heap) Push(val int) {
	if obj.size == obj.maxSize {
		panic("heap is full")
	}
	obj.data[obj.size] = val
	obj.heapInsert()
	obj.size++
}

func (obj *Heap) Pop() int {
	if obj.IsEmpty() {
		panic("heap is empty")
	}
	ans := obj.data[0]
	obj.size--
	obj.data[0], obj.data[obj.size] = obj.data[obj.size], obj.data[0]
	obj.heapify()
	return ans
}

// 插入 从obj.size位置开始往上调整堆
func (obj *Heap) heapInsert() {
	index := obj.size
	parent := (index - 1) / 2
	for obj.data[index] < obj.data[parent] {
		obj.data[index], obj.data[parent] = obj.data[parent], obj.data[index]
		index = parent
		parent = (index - 1) / 2
	}
}

// 下沉 从0位置开始向下调整堆
func (obj *Heap) heapify() {
	index := 0
	left := 2*index + 1
	for left < obj.size {
		// 左右两个孩子谁小
		largest := left
		if left+1 < obj.size && obj.data[left+1] < obj.data[left] {
			largest = left + 1
		}

		// 跟父节点比较
		if obj.data[index] < obj.data[largest] {
			largest = index
		}

		// 如果此时父节点最小就不用下沉
		if largest == index {
			break
		}

		// 交换下沉
		obj.data[index], obj.data[largest] = obj.data[largest], obj.data[index]
		index = largest
		left = 2*index + 1
	}
}

// ============================================================================

// 在数组中从index出发向下调整堆
func heapify(arr []int, index int, n int) {
	left := 2*index + 1
	for left < n {
		// 孩子谁小
		largest := left
		if left+1 < n && arr[left+1] < arr[left] {
			largest = left + 1
		}
		// 跟父节点比谁小
		if arr[index] < arr[largest] {
			largest = index
		}
		// 父节点最小不用下沉了
		if largest == index {
			break
		}

		// 交换下沉
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*index + 1
	}
}

func heapify2(arr []int, index int, heapSize int) {
	left := 2*index + 1
	for left < heapSize {
		// 孩子谁大
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		// 跟父节点比谁大
		if arr[index] > arr[largest] {
			largest = index
		}
		// 父节点最大不用下沉了
		if largest == index {
			break
		}

		// 交换下沉
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*index + 1
	}
}

// 堆排序, 时间复杂度 O(NlogN), 额外空间复杂度是O(1) 的方法
//
//	从小到大排序要用大根堆
//
// 从大到小用小根堆
func HeapSort(arr []int) {
	n := len(arr)
	if len(arr) < 2 {
		return
	}

	// 建堆
	for i := n - 1; i >= 0; i-- {
		heapify2(arr, i, n)
	}

	// 排序过程
	heapSize := n
	heapSize--
	arr[0], arr[heapSize] = arr[heapSize], arr[0]
	for heapSize > 0 {
		heapify2(arr, 0, heapSize)
		heapSize--
		arr[0], arr[heapSize] = arr[heapSize], arr[0]
	}

}
