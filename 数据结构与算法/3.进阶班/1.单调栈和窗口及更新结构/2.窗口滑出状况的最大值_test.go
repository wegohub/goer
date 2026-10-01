package class01

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(windowMax([]int{4, 3, 5, 4, 3, 3, 6, 7}, 3))

	return "Hello World!", nil
}

func windowMax(arr []int, W int) []int {
	if len(arr) == 0 || W < 1 || len(arr) < W {
		return nil
	}

	// 放最大值的双端队列
	qMax := list.New()

	// 结果的数量 len(arr) - W + 1
	ans := make([]int, len(arr)-W+1)

	// ans结果下标
	index := 0

	for R := 0; R < len(arr); R++ {
		for qMax.Len() > 0 && arr[qMax.Back().Value.(int)] <= arr[R] {
			qMax.Remove(qMax.Back())
		}
		qMax.PushBack(R)

		// R - W 是过期下标
		if qMax.Front().Value.(int) == R-W {
			qMax.Remove(qMax.Front())
		}

		// 收集答案
		if R >= W-1 {
			ans[index] = arr[qMax.Front().Value.(int)]
			index++
		}

	}

	return ans
}
