package class05

import (
	"fmt"
	"math"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{101, 2, 202, 41, 302}
	RadixSort(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

// RadixSort 基数排序
func RadixSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	// 基数定为10, 处理十进制的数
	radix := 10

	// 数组中的最大值
	max := math.MinInt
	for _, item := range arr {
		if item > max {
			max = item
		}
	}
	// 最大值有多少位
	bits := 0
	for max > 0 {
		bits++
		max = max / 10
	}

	// 准备一个辅助数组
	help := make([]int, n)

	for d := 1; d <= bits; d++ {
		// 准备基数长度的词频数组
		count := make([]int, radix)
		for _, item := range arr {
			// 个位、十位、百位上的数
			bitNum := (item / int(math.Pow10(d-1))) % 10
			count[bitNum] += 1
		}
		// 将count加工厂前缀和数组
		for i := 1; i < len(count); i++ {
			count[i] = count[i] + count[i-1]
		}
		// 处理arr，从后向前遍历以确保稳定性
		for i := n - 1; i >= 0; i-- {
			bitNum := (arr[i] / int(math.Pow10(d-1))) % 10
			help[count[bitNum]-1] = arr[i]
			count[bitNum]--
		}
		// 将help刷回arr
		for index, item := range help {
			arr[index] = item
		}

	}
}
