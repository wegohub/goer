package leetcode

import "container/list"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxSlidingWindow(arr []int, W int) []int {
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

// 数组代替双端队列
func maxSlidingWindow1(nums []int, W int) []int {
	if len(nums) == 0 || W > len(nums) {
		return []int{}
	}
	N := len(nums)
	ans := make([]int, 0, N-W+1)
	dueue := make([]int, 0)
	R := 0
	for R < N {
		// R 位置的数入窗口
		for len(dueue) > 0 && nums[dueue[len(dueue)-1]] <= nums[R] {
			dueue = dueue[0 : len(dueue)-1]
		}
		dueue = append(dueue, R)

		// L位置的数出窗口
		if dueue[0] == R-W {
			dueue = dueue[1:]
		}

		// 收集答案
		if R >= W-1 {
			ans = append(ans, nums[dueue[0]])
		}

		R++
	}
	return ans
}
