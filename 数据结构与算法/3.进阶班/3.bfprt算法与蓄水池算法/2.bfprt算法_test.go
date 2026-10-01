package class03

import (
	"fmt"
	"math"
	"math/rand"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr1 := []int{3, 2, 1, 5, 6, 4}
	ans1 := process(arr1, 0, len(arr1)-1, 1)
	arr2 := []int{6, 3, 1, 7, 4, 8, 9, 3, 4}
	ans2 := bfprt(arr2, 0, len(arr2)-1, 3)
	fmt.Println(ans1, ans2)
	return "Hello World!", nil
}

// arr[L..R]  如果排序的话，位于index位置的数，是什么，返回
func bfprt(arr []int, L, R, index int) int {
	if L == R { // L == R == index
		return arr[L]
	}

	// 不止一个数 L + [0, R-L]
	pivot := medianOfMedians(arr, L, R)

	// 返回 == 区域的左右边界
	rang := partition(arr, L, R, pivot)

	if index >= rang[0] && index <= rang[1] {
		return arr[index]
	} else if index < rang[0] {
		return bfprt(arr, L, rang[0]-1, index)
	} else {
		return bfprt(arr, rang[1]+1, R, index)
	}
}

// arr[L...R]  五个数一组
// 每个小组内部排序
// 每个小组中位数领出来，组成marr
// marr中的中位数，返回
func medianOfMedians(arr []int, L, R int) int {
	size := R - L + 1
	offset := 1
	if size%5 == 0 {
		offset = 0
	}
	mArr := make([]int, size/5+offset)
	for team := 0; team < len(mArr); team++ {
		teamStart := L + team*5
		// 每5个数一组，求中位数放到mArr中
		// L ... L + 4
		// L + 5 ... L + 9
		// L +1 0....L + 14
		mArr[team] = getMedian(arr, teamStart, int(math.Min(float64(R), float64(teamStart+4))))
	}

	// marr中，找到中位数
	// marr(0, marr.len - 1,  mArr.length / 2 )
	return bfprt(mArr, 0, len(mArr)-1, len(mArr)/2)
}

func getMedian(arr []int, L, R int) int {
	// 插入排序
	insertSort(arr, L, R)
	// return arr[(L + R) / 2]
	return arr[L+(R-L)/2]
}

func insertSort(arr []int, L, R int) {
	for i := L; i < R; i++ {
		for j := i + 1; j > L && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
}

// 求 arr 中第k小的数
// process(arr, 0, N-1, K-1)
// arr[L...R]范围上，如果排序的话，找到index的数
// index [L...R]
func process(arr []int, L, R, index int) int {
	if L == R { // L == R == index
		return arr[L]
	}

	// 不止一个数 L + [0, R-L]
	pivot := arr[L+int(rand.Float64()*float64(R-L+1))]

	// 返回 == 区域的左右边界
	rang := partition(arr, L, R, pivot)

	if index >= rang[0] && index <= rang[1] {
		return arr[index]
	} else if index < rang[0] {
		return process(arr, L, rang[0]-1, index)
	} else {
		return process(arr, rang[1]+1, R, index)
	}
}

func partition(arr []int, L, R, pivot int) [2]int {
	// 小于区域
	less := L - 1
	// 大于区域
	more := R + 1
	// 当前位置的数
	index := L

	for index < more {
		if arr[index] < pivot {
			// 当前位置的数和小于区域的前一个交换
			arr[index], arr[less+1] = arr[less+1], arr[index]
			less++
			index++
		} else if arr[index] == pivot {
			index++
		} else {
			// 大于区域的前一个数与当前数交换, 大于区域左扩，index不动
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		}
	}
	return [2]int{less + 1, more - 1}
}
