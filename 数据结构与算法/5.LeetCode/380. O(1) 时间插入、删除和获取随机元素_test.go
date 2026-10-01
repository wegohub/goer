package leetcode

import (
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 删除的时候用最后一条记录填删除的洞，保证index连续
// RandomizedSet represents a set that supports insertion, removal, and getting a random element.
type RandomizedSet struct {
	valueIndexMap map[int]int
	indexValueMap map[int]int
	size          int
}

// Constructor initializes a new RandomizedSet.
func Constructor() RandomizedSet {
	return RandomizedSet{
		valueIndexMap: make(map[int]int),
		indexValueMap: make(map[int]int),
		size:          0,
	}
}

// Insert adds a value to the set.
func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valueIndexMap[val]; !exists {
		this.valueIndexMap[val] = this.size
		this.indexValueMap[this.size] = val
		this.size++
		return true
	}
	return false
}

// Remove removes a value from the set.
func (this *RandomizedSet) Remove(val int) bool {
	if idx, exists := this.valueIndexMap[val]; exists {
		this.size--
		lastKey := this.indexValueMap[this.size]
		this.valueIndexMap[lastKey] = idx
		this.indexValueMap[idx] = lastKey
		delete(this.valueIndexMap, val)
		delete(this.indexValueMap, this.size)
		return true
	}
	return false
}

// GetRandom returns a random element from the set.
func (this *RandomizedSet) GetRandom() int {
	if this.size == 0 {
		return -1
	}
	randomIndex := int(rand.Float64() * float64(this.size))
	return this.indexValueMap[randomIndex]
}
