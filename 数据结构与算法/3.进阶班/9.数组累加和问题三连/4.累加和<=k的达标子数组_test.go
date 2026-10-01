package class10

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxLength(arr []int, k int) int {
	if len(arr) == 0 {
		return 0
	}

	N := len(arr)

	// 以i位置开头，累加和最小能扩到那个位置(动态规划)
	minSums := make([]int, N)
	minSumEnds := make([]int, N)

	minSums[N-1] = arr[N-1]
	minSumEnds[N-1] = N - 1

	for i := N - 2; i >= 0; i-- {
		if minSums[i+1] <= 0 {
			minSums[i] = arr[i] + minSums[i+1]
			minSumEnds[i] = minSumEnds[i+1]
		} else {
			minSums[i] = arr[i]
			minSumEnds[i] = i
		}
	}

	// (i...)(...)(...)(end...x) 第一个不达标的位置
	end := 0
	// i............sum(end...x)
	sum := 0
	// 结果
	maxLen := 0

	for i := 0; i < N; i++ {
		for end < N && sum+minSums[end] <= k {
			sum += minSums[end]
			end = minSumEnds[end] + 1
		}
		// [i.........](end...x)
		// i..................x 越界
		maxLen = int(math.Max(float64(maxLen), float64(end-i)))

		// 缩窗口
		if end > i { // 窗口内还有数
			sum -= arr[i]
		} else { // 窗口内没有数了，i==end，即将i++,所以让end跟着一起走，换一个开头尝试
			end = i + 1
		}
	}

	return maxLen
}
