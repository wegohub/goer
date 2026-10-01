package class01

import (
	"container/list"
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(getNum([]int{1, 2, 3, 4, 5}, 3))

	return "Hello World!", nil
}

func getNum(arr []int, num int) int {
	if len(arr) == 0 {
		return 0
	}

	// 最大值更新结构
	qmax := list.New()
	// 最小值更新结构
	qmin := list.New()

	// 窗口的左右边界 [L, R), R表示第一个不达标的位置
	L := 0
	R := 0
	ans := 0

	for L < len(arr) { // L 是开头位置，尝试每一个开头

		// 此时以L开头，R向右扩到违规为止
		for R < len(arr) {
			for qmax.Len() > 0 && arr[qmax.Back().Value.(int)] <= arr[R] {
				qmax.Remove(qmax.Back())
			}
			qmax.PushBack(R)

			for qmin.Len() > 0 && arr[qmin.Back().Value.(int)] >= arr[R] {
				qmin.Remove(qmin.Back())
			}
			qmin.PushBack(R)

			if arr[qmax.Front().Value.(int)]-arr[qmin.Front().Value.(int)] > num {
				break
			}
			R++
		}

		// L ... R 达标，[L, R-1]范围上的每一个子数组都成立 (R-1) - L + 1
		ans += R - L

		if qmax.Front().Value.(int) == L {
			qmax.Remove(qmax.Front())
		}

		if qmin.Front().Value.(int) == L {
			qmin.Remove(qmin.Front())
		}

		L++
	}

	return ans
}
