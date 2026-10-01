package leetcode

import "container/list"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func canCompleteCircuit(gas []int, cost []int) int {
	if len(gas) == 0 || len(cost) == 0 {
		return -1
	}
	N := len(gas)
	// 将gas和cost处理成能耗数组,并将长度增加为原来的2倍
	arr := make([]int, 2*N)
	for i := 0; i < N; i++ {
		arr[i] = gas[i] - cost[i]
		arr[i+N] = arr[i]
	}

	// 将能耗数组arr转换为前缀和数组
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}

	// 准备一个长度为N的最小值窗口
	minWin := list.New()
	for i := 0; i < N; i++ {
		for minWin.Len() != 0 && arr[minWin.Back().Value.(int)] >= arr[i] {
			minWin.Remove(minWin.Back())
		}
		minWin.PushBack(i)
	}

	// 迭代出每个点出发是否达标数组
	ans := make([]bool, N)

	R := N
	for L := 0; L < N; L++ {
		// 刚出窗口的值
		pre := 0
		if L > 0 {
			pre = arr[L-1]
		}
		// 如果当前窗口中最小值 - 刚出窗口的值 > 0 则满足条件
		if arr[minWin.Front().Value.(int)]-pre >= 0 {
			// 这里可以直接return的, 这样做可以收集所有满足条件的点
			ans[L] = true
		}

		// R 位置的数入窗口
		for minWin.Len() != 0 && arr[minWin.Back().Value.(int)] >= arr[R] {
			minWin.Remove(minWin.Back())
		}
		minWin.PushBack(R)

		// L 位置的数出窗口
		if minWin.Front().Value.(int) == L {
			minWin.Remove(minWin.Front())
		}

		R++
	}

	for i := 0; i < len(ans); i++ {
		if ans[i] {
			return i
		}
	}
	return -1
}

func canCompleteCircuit1(gas []int, cost []int) int {
	N := len(gas)
	// 能耗数组
	energy := make([]int, 2*N)
	for i := 0; i < N; i++ {
		val := gas[i] - cost[i]
		energy[i] = val
		energy[i+N] = val
	}
	// 前缀和
	energySum := make([]int, 2*N)
	energySum[0] = energy[0]
	for i := 1; i < len(energySum); i++ {
		energySum[i] = energy[i] + energySum[i-1]
	}
	// 准备一个大小为N的窗口
	winSize := N
	L := 0
	R := 0
	queue := list.New()
	for R < winSize {
		// R位置的数要进窗口
		for queue.Len() > 0 && energySum[queue.Back().Value.(int)] >= energySum[R] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(R)
		R++
	}

	// 窗口内的最小值 - 出窗口的值 = 瓶颈
	for R < len(energySum) {
		for queue.Len() > 0 && energySum[queue.Back().Value.(int)] >= energySum[R] {
			queue.Remove(queue.Back())
		}
		queue.PushBack(R)

		// L 位置的数出窗口
		//if queue.Front().Value.(int) == L {
		//	queue.Remove(queue.Front())
		//}

		// 窗口内的最小值
		minVal := energySum[queue.Front().Value.(int)]

		// 出窗口的值
		outVal := 0
		if L > 0 {
			outVal = energySum[L-1]
		}

		if minVal-outVal >= 0 {
			return L
		}

		L++
		R++
	}

	return -1
}
