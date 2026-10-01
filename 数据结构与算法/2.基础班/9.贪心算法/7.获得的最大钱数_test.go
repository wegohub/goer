package class09

import (
	"container/heap"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	costs := []int{1, 1, 4, 5, 1}
	profixs := []int{3, 1, 3, 2, 3}
	K := 2
	M := 1

	ans := maxProfits(costs, profixs, K, M)
	fmt.Println(ans)

	return "Hello World!", nil
}

type Program struct {
	cost   int // 花费
	profit int // 利润
}

// ConstMinHeap 花费小根堆
type ConstMinHeap []*Program

func (cls ConstMinHeap) Len() int {
	return len(cls)
}

func (cls ConstMinHeap) Less(i, j int) bool {
	return cls[i].cost < cls[j].cost
}

func (cls ConstMinHeap) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (cls ConstMinHeap) Peek() int {
	return cls[0].cost
}

func (obj *ConstMinHeap) Push(v interface{}) {
	*obj = append(*obj, v.(*Program))
}

func (obj *ConstMinHeap) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

// ProfitMaxHeap 收益大根堆
type ProfitMaxHeap []*Program

func (cls ProfitMaxHeap) Len() int {
	return len(cls)
}

func (cls ProfitMaxHeap) Less(i, j int) bool {
	return cls[i].profit > cls[j].profit
}

func (cls ProfitMaxHeap) Swap(i, j int) {
	cls[i], cls[j] = cls[j], cls[i]
}

func (cls ProfitMaxHeap) Peek() int {
	return cls[0].profit
}

func (obj *ProfitMaxHeap) Push(v interface{}) {
	*obj = append(*obj, v.(*Program))
}

func (obj *ProfitMaxHeap) Pop() interface{} {
	old := *obj
	n := len(old)
	v := old[n-1]
	*obj = old[0 : n-1]
	return v
}

func maxProfits(costs []int, profits []int, K, M int) int {
	if len(costs) != len(profits) || K == 0 {
		return 0
	}
	// 所有的项目
	programs := make([]*Program, len(costs))
	for i := 0; i < len(costs); i++ {
		programs[i] = &Program{
			cost:   costs[i],
			profit: profits[i],
		}
	}

	// 项目入花费小根堆
	costHeap := &ConstMinHeap{}
	//*costHeap = append(*costHeap, programs...)
	// heap.Init(costHeap)
	for _, program := range programs {
		heap.Push(costHeap, program)
	}

	// 利润大根堆
	profixHeap := &ProfitMaxHeap{}

	// 做K个项目
	for k := 1; k <= K; k++ {
		// 解锁一批项目
		for costHeap.Len() > 0 && costHeap.Peek() <= M {
			program := heap.Pop(costHeap)
			heap.Push(profixHeap, program)
		}

		// 弹出利润堆顶项目来做
		if profixHeap.Len() > 0 {
			program := heap.Pop(profixHeap).(*Program)
			M += program.profit
		} else { // 没有项目可做
			return M
		}
	}

	return M
}
