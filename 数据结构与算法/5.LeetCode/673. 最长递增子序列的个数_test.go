package leetcode

import (
	"math"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// findNumberOfLIS2 找到数组中最长递增子序列的数量
func findNumberOfLIS(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	// dp 是一个列表，每个元素是一个 TreeMap，存储以某个长度结尾的 LIS 的数量
	dp := []map[int]int{}

	for i := 0; i < len(nums); i++ {
		L := 0
		R := len(dp) - 1
		find := -1
		for L <= R {
			mid := (L + R) / 2
			if minKey(dp[mid]) >= nums[i] {
				find = mid
				R = mid - 1
			} else {
				L = mid + 1
			}
		}
		if find == -1 {
			if len(dp) == 0 {
				dp = append(dp, map[int]int{nums[i]: 1})
			} else {
				dp = append(dp, map[int]int{})
				index := len(dp) - 1
				cur := dp[index]
				size := dp[index-1][minKey(dp[index-1])]
				if ceilingKey(dp[index-1], nums[i]) != math.MinInt32 {
					size -= dp[index-1][ceilingKey(dp[index-1], nums[i])]
				}
				cur[nums[i]] = size
			}
		} else {
			newAdd := 1
			if find > 0 {
				pre := dp[find-1]
				newAdd = pre[minKey(pre)]
				if ceilingKey(pre, nums[i]) != math.MinInt32 {
					newAdd -= pre[ceilingKey(pre, nums[i])]
				}
			}
			cur := dp[find]
			if cur[minKey(cur)] == nums[i] {
				cur[nums[i]] += newAdd
			} else {
				preNum := cur[minKey(cur)]
				cur[nums[i]] = newAdd + preNum
			}
		}
	}
	return dp[len(dp)-1][minKey(dp[len(dp)-1])]
}

// minKey 返回 map 中的最小键
func minKey(m map[int]int) int {
	min := math.MaxInt32
	for k := range m {
		if k < min {
			min = k
		}
	}
	return min
}

// ceilingKey 返回大于等于给定值的最小键
func ceilingKey(m map[int]int, key int) int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		if k >= key {
			return k
		}
	}
	return math.MinInt32
}
