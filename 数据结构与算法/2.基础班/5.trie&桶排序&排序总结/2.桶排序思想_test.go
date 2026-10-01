package class05

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	arr := []int{3, 1, 6, 4, 1, 7, 8, 6}
	CountSort(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

// CountSort 计数排序，样本量很小的情况下适用
func CountSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	// 先找出arr中的最大值
	max := math.MinInt
	for _, item := range arr {
		max = int(math.Max(float64(item), float64(max)))
	}

	// 准备max+1长度的桶, index是值，value是词频
	bucket := make([]int, max+1)

	// 将arr入桶
	for _, item := range arr {
		bucket[item] += 1
	}

	// 将桶倒出
	index := 0
	for val, count := range bucket {
		for count > 0 {
			arr[index] = val
			count--
			index++
		}
	}
}
