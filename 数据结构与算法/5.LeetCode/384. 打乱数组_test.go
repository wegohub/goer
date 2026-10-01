package leetcode

import (
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// [0-n-1]上随机选一个位置与n-1的数交换
// Solution represents a solution for shuffling and resetting an array.
type Solution struct {
	origin  []int
	shuffle []int
	N       int
}

// Constructor initializes a new Solution with the given nums.
func Constructor(nums []int) Solution {
	N := len(nums)
	shuffle := make([]int, N)
	copy(shuffle, nums)
	return Solution{
		origin:  nums,
		shuffle: shuffle,
		N:       N,
	}
}

// Reset returns the original array.
func (this *Solution) Reset() []int {
	return this.origin
}

// 0 - n-1 上随机选一个数放到n-1上
// 0 - n -2 上随机选一个数放到n-2上
// ...
// 0 - 0
func (this *Solution) Shuffle() []int {
	for i := this.N - 1; i >= 0; i-- {
		r := rand.Intn(i + 1)
		this.shuffle[i], this.shuffle[r] = this.shuffle[r], this.shuffle[i]
	}
	return this.shuffle
}
