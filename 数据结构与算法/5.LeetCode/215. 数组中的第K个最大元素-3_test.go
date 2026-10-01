package leetcode

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{3, 2, 1, 5, 6, 4}
	ans := findKthLargest(arr, 6)
	fmt.Println(ans)
	return "Hello World!", nil
}

func findKthLargest(nums []int, k int) int {
	return bfprt(nums, 0, len(nums)-1, len(nums)-k)
}

func bfprt(nums []int, L, R, index int) int {
	if L == R {
		return nums[L]
	}
	// [0, R-L]
	poivt := ChosenSon(nums, L, R)

	rang := partition(nums, L, R, poivt)

	if index >= rang[0] && index <= rang[1] {
		return nums[index]
	} else if index < rang[0] {
		return bfprt(nums, L, rang[0]-1, index)
	} else {
		return bfprt(nums, rang[1]+1, R, index)
	}
}

// [L, R]上选出天选之子
func ChosenSon(nums []int, L, R int) int {
	size := R - L + 1
	offset := 1
	if size%5 == 0 {
		offset = 0
	}
	marr := make([]int, size/5+offset)
	for team := 0; team < len(marr); team++ {
		teamStart := L + team*5
		marr[team] = getMiddleNum(nums, teamStart, int(math.Min(float64(R), float64(teamStart+4))))
	}

	return bfprt(marr, 0, len(marr)-1, len(marr)/2)
}

// [L, R]上取中位数
func getMiddleNum(nums []int, L, R int) int {
	for i := L; i < R; i++ {
		for j := i + 1; j > L && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums[L+(R-L)/2]
}

func partition(nums []int, L, R, poivt int) [2]int {
	less := L - 1
	more := R + 1
	index := L
	for index < more {
		if nums[index] < poivt {
			nums[index], nums[less+1] = nums[less+1], nums[index]
			less++
			index++
		} else if nums[index] == poivt {
			index++
		} else {
			nums[index], nums[more-1] = nums[more-1], nums[index]
			more--
		}
	}
	return [2]int{less + 1, more - 1}
}
