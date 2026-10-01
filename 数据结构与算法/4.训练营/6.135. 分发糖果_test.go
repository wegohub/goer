package train

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func candy(arr []int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}
	if len(arr) == 1 {
		return 1
	}
	n := len(arr)

	// 初始化每个孩子至少得到一个糖果
	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	// 从左到右遍历，确保右边评分更高的孩子得到更多糖果
	for i := 1; i < n; i++ {
		if arr[i] > arr[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	// 从右到左遍历，确保左边评分更高的孩子得到更多糖果
	for i := n - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			candies[i] = max(candies[i], candies[i+1]+1)
		}
	}

	// 计算总糖果数
	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
