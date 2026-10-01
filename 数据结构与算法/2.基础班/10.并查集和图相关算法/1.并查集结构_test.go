package class10

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// 并查集map实现
	set := NewUnionSet([]int{1, 2, 3, 4, 5, 6, 7, 8})
	fmt.Println(set.IsSameSet(1, 2))
	set.Union(1, 2)
	fmt.Println(set.IsSameSet(1, 2))

	// 并查集数组实现
	return "Hello World!", nil
}

type Node struct {
	Val int
}

func NewUnionSet(values []int) *UnionSetByMap {
	nodes := make(map[int]*Node)
	parents := make(map[*Node]*Node)
	sizeMap := make(map[*Node]int)

	for _, item := range values {
		node := &Node{Val: item}
		nodes[item] = node
		parents[node] = node
		sizeMap[node] = 1
	}

	return &UnionSetByMap{
		nodes:   nodes,
		parents: parents,
		sizeMap: sizeMap,
	}
}

// UnionSetByMap 并查集数组实现
type UnionSetByMap struct {
	// 节点表 k => Node
	nodes map[int]*Node
	// 父级表
	parents map[*Node]*Node
	// 代表点的记录
	sizeMap map[*Node]int
}

// find 从当前节点找父节点
func (obj *UnionSetByMap) find(cur *Node) *Node {
	stack := make([]*Node, 0)
	index := 0
	for cur != obj.parents[cur] {
		stack = append(stack, cur)
		index++
		cur = obj.parents[cur]
	}

	for index--; index >= 0; index-- {
		obj.parents[stack[index]] = cur
	}
	return cur
}

// IsSameSet a和b是否在同一个集合
func (obj *UnionSetByMap) IsSameSet(a, b int) bool {
	if obj.nodes[a] == nil || obj.nodes[b] == nil {
		return false
	}
	return obj.find(obj.nodes[a]) == obj.find(obj.nodes[b])
}

// Union 合并a和b
func (obj UnionSetByMap) Union(a, b int) {
	if obj.nodes[a] == nil || obj.nodes[b] == nil {
		return
	}
	aHead := obj.find(obj.nodes[a])
	bHead := obj.find(obj.nodes[b])
	if aHead != bHead {
		aSize := obj.sizeMap[aHead]
		bSize := obj.sizeMap[bHead]
		if aSize < bSize {
			obj.parents[aHead] = bHead
			obj.sizeMap[bHead] += obj.sizeMap[aHead]
			delete(obj.sizeMap, aHead)
		} else {
			obj.parents[bHead] = aHead
			obj.sizeMap[aHead] += obj.sizeMap[bHead]
			delete(obj.sizeMap, bHead)
		}
	}
}

func NewUnionSetArr(values []int) *UnionSetArr {
	parent := make([]int, len(values))
	size := make([]int, len(values))
	help := make([]int, len(values))

	for index, _ := range values {
		parent[index] = index
		size[index] = 1
	}

	return &UnionSetArr{
		values: values,
		parent: parent,
		size:   size,
		help:   help,
	}

}

// UnionSetArr 并查集数组实现
type UnionSetArr struct {
	values []int // 值数组
	parent []int // 父节点下标
	size   []int // 代表节点下标
	help   []int // 做栈
}

func (obj *UnionSetArr) IsSameSet(a, b int) bool {
	return obj.parent[a] == obj.parent[b]
}

// Union 用下标合并
func (obj *UnionSetArr) Union(a, b int) {
	aHead := obj.parent[a]
	bHead := obj.parent[b]
	if aHead != bHead {
		if obj.size[aHead] > obj.size[bHead] {
			obj.parent[bHead] = aHead
			obj.size[aHead] += obj.size[bHead]
		} else {
			obj.parent[aHead] = bHead
			obj.size[bHead] += obj.size[aHead]
		}
	}
}

// 从cur位置往上找，找到不能再往上了，返回此时的代表节点
func (obj *UnionSetArr) find(cur int) int {
	height := 0
	for cur != obj.parent[cur] {
		obj.help[height] = cur
		height++
		cur = obj.parent[cur]
	}
	// 此时的cur就是代表节点的下标
	height--
	for height >= 0 {
		height--
		obj.parent[obj.help[height]] = cur
	}
	return cur
}
