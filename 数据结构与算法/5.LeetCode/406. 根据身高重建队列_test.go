package leetcode

import "sort"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 1. 按照身高从大到小排序，身高一样的按照次数从小到达排序
// 2. 排序后的数组按照次数作为索引插入结果切片中
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		// 身高一样
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	result := make([][]int, len(people))
	for _, item := range people {
		// 这里导致了该算法是O(N^2), 用有序表优化到O(LogN)
		result = append(result[:item[1]], append([][]int{item}, result[item[1]:]...)...)
	}
	return result[:len(people)]
}
