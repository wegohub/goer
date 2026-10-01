package class15

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	tree := NewIndexTree(10)
	tree.Add(2, 2)
	tree.Add(3, 1)
	fmt.Println(tree.Sum(3))
	tree.Add(2, 4)
	fmt.Println(tree.Sum(3))
	return "Hello World!", nil
}

func NewIndexTree(size int) *IndexTree {
	return &IndexTree{
		tree: make([]int, size+1), // 为什么长度要加1？因为0位置弃而不用
		size: size,
	}
}

// 一维IndexTree
type IndexTree struct {
	tree []int // 数据
	size int   // 大小
}

// Sum 返回[0-index]的累加和
func (obj *IndexTree) Sum(index int) int {
	ans := 0
	for index > 0 {
		ans += obj.tree[index]
		// 抹掉最右侧的1
		index -= index & -index
	}
	return ans
}

// Add  index位置的数，想加上d，还有哪些位置也要都加d
func (obj *IndexTree) Add(index int, d int) {
	for index <= obj.size {
		obj.tree[index] += d
		// 加上最右侧的1
		index += index & -index
	}
}

// 二维IndexTree 任何左上角和右下角坐标得出该区域的累加和
type IndexTreeTwo struct {
}
