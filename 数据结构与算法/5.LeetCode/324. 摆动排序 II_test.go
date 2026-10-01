package leetcode

import (
	"math/rand"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 时间复杂度：O(N*LogN) 空间复杂度：O(N)的解法
// 满足条件重复的元素不会超过(N+1)/2  [1 2 1]
// 1. 将nums拷贝得到数组arr, 并将arr排序
// 2. 依次将arr中的最后一个和 (N+1)/2个 填入到nums中
func wiggleSort2(nums []int) {
	n := len(nums)
	arr := append([]int{}, nums...)
	sort.Ints(arr)
	x := (n + 1) / 2
	for i, j, k := 0, x-1, n-1; i < n; i += 2 {
		nums[i] = arr[j]
		if i+1 < n {
			nums[i+1] = arr[k]
		}
		j--
		k--
	}
}

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	N := len(nums)
	// 找到中位数
	findIndexNum(nums, 0, N-1, N/2)

	//median := findIndexNum(nums, 0, N-1, N / 2)
	// findIndexNum 已经调好了，所有不用再调整
	//partition(nums, 0, N-1, median)

	// 完美洗牌问题 [3 2 1 4 | 4 7 6 8]
	if (N & 1) == 0 { // 偶数
		shuffle(nums, 0, N-1)
		reverse(nums, 0, N-1)
	} else { // 奇数
		shuffle(nums, 1, N-1)
	}
}

// 在数组[L,R]范围上找到中位数
func findIndexNum(arr []int, L, R, index int) int {
	for L < R {
		pivot := arr[L+int(rand.Float64()*float64(R-L+1))]
		left, right := partition(arr, L, R, pivot)
		if index >= left && index <= right {
			return arr[index]
		} else if index < left {
			R = left - 1
		} else {
			L = right + 1
		}
	}
	return arr[L]
}

// 荷兰国旗问题
func partition(arr []int, L, R, pivot int) (int, int) {
	less := L - 1
	more := R + 1
	index := L
	for index < more {
		if arr[index] < pivot {
			arr[index], arr[less+1] = arr[less+1], arr[index]
			less++
			index++
		} else if arr[index] > pivot {
			arr[index], arr[more-1] = arr[more-1], arr[index]
			more--
		} else {
			index++
		}
	}
	return less + 1, more - 1
}

// shuffle shuffles the elements in the array nums from index l to r.
func shuffle(nums []int, l, r int) {
	for r-l+1 > 0 {
		lenAndOne := r - l + 2
		bloom := 3
		k := 1
		for bloom <= lenAndOne/3 {
			bloom *= 3
			k++
		}
		m := (bloom - 1) / 2
		mid := (l + r) / 2
		rotate(nums, l+m, mid, mid+m)
		cycles(nums, l-1, bloom, k)
		l = l + bloom - 1
	}
}

// cycles performs cyclic permutations on the array nums.
func cycles(nums []int, base, bloom, k int) {
	for i, trigger := 0, 1; i < k; i, trigger = i+1, trigger*3 {
		next := (2 * trigger) % bloom
		cur := next
		record := nums[next+base]
		var tmp int
		nums[next+base] = nums[trigger+base]
		for cur != trigger {
			next = (2 * cur) % bloom
			tmp = nums[next+base]
			nums[next+base] = record
			cur = next
			record = tmp
		}
	}
}

// rotate rotates the elements in the array arr from index l to m and m+1 to r.
func rotate(arr []int, l, m, r int) {
	reverse(arr, l, m)
	reverse(arr, m+1, r)
	reverse(arr, l, r)
}

// reverse reverses the elements in the array arr from index l to r.
func reverse(arr []int, l, r int) {
	for l < r {
		swap(arr, l, r)
		l++
		r--
	}
}

// swap swaps the elements at indices i and j in the array nums.
func swap(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
