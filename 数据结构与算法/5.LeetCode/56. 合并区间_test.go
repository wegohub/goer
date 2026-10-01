package leetcode

import (
	"math"
	"sort"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func merge(intervals [][]int) [][]int {
	// 第一个位置排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 结果
	ans := make([][]int, 0)

	// 每一组的开始和结尾
	start := intervals[0][0]
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		// 下一组的开始 > 上一组的结尾，结算答案
		if intervals[i][0] > end {
			ans = append(ans, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][1]
		} else {
			// 推高当前组的结尾位置
			end = int(math.Max(float64(end), float64(intervals[i][1])))
		}
	}
	// 最后一组加入答案
	ans = append(ans, []int{start, end})

	return ans
}
