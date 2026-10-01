package leetcode

import "sort"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	// 第一个数选了i位置的数，
	for i := 0; i < len(nums)-2; i++ {
		// 小加速
		if nums[i] > 0 {
			break
		}
		// 去重
		if i == 0 || nums[i-1] != nums[i] {
			nexts := twoSum(nums, i+1, -nums[i])
			for _, next := range nexts {
				ans = append(ans, []int{nums[i], next[0], next[1]})
			}
		}
	}

	return ans
}

// 两数之和
func twoSum(nums []int, begin int, target int) [][]int {
	L := begin
	R := len(nums) - 1
	ans := make([][]int, 0)
	for L < R {
		if nums[L]+nums[R] > target {
			R--
		} else if nums[L]+nums[R] < target {
			L++
		} else {
			// 去重
			if L == begin || nums[L-1] != nums[L] {
				ans = append(ans, []int{nums[L], nums[R]})
			}
			L++
		}
	}
	return ans
}
