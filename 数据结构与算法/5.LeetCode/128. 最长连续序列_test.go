package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))
	return "Hello World!", nil
}

// 最优解
func longestConsecutive2(nums []int) int {
	mp := make(map[int]int)
	max := 0
	for _, num := range nums {
		if _, ok := mp[num]; !ok {
			mp[num] = 1
			pre := mp[num-1]
			pos := mp[num+1]
			all := pre + pos + 1
			mp[num-pre] = all
			mp[num+pos] = all
			max = int(math.Max(float64(max), float64(all)))
		}
	}
	return max
}

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	set := NewUnionSet(nums)
	for _, num := range nums {
		set.Union(num, num-1)
		set.Union(num, num+1)
	}
	return set.maxSize
}

func NewUnionSet(arr []int) *UnionSet {
	parentMap := make(map[int]int)
	sizeMap := make(map[int]int)
	for _, item := range arr {
		parentMap[item] = item
		sizeMap[item] = 1
	}
	return &UnionSet{
		parentMap: parentMap,
		sizeMap:   sizeMap,
		help:      make([]int, len(arr)),
		maxSize:   1,
	}
}

type UnionSet struct {
	parentMap map[int]int // 节点对应的父节点
	sizeMap   map[int]int // 父节点大小
	help      []int
	maxSize   int // 最大区域节点个数
}

func (obj *UnionSet) find(cur int) int {
	index := 0
	for obj.parentMap[cur] != cur {
		obj.help[index] = cur
		index++
		cur = obj.parentMap[cur]
	}

	// 扁平化
	for index--; index >= 0; index-- {
		obj.parentMap[obj.help[index]] = cur
	}

	return cur
}

func (obj *UnionSet) IsSamesite(a, b int) bool {
	return obj.find(a) == obj.find(b)
}

func (obj *UnionSet) Union(a, b int) {
	// a 和 b都要存在
	if _, ok := obj.parentMap[a]; !ok {
		return
	}
	if _, ok := obj.parentMap[b]; !ok {
		return
	}
	aHead := obj.find(a)
	bHead := obj.find(b)
	if aHead != bHead {
		big := aHead
		small := bHead
		if obj.sizeMap[bHead] > obj.sizeMap[aHead] {
			big = bHead
			small = aHead
		}
		// 小挂大
		obj.parentMap[small] = big
		obj.sizeMap[big] += obj.sizeMap[small]
		delete(obj.sizeMap, small)
		// 更新max
		obj.maxSize = int(math.Max(float64(obj.maxSize), float64(obj.sizeMap[big])))
	}
}

func (obj *UnionSet) MaxSize() int {
	return obj.maxSize
}
