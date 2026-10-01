package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 最优解
func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}
	N := len(height)
	L := 1
	leftMax := height[0]
	R := N - 2
	rightMax := height[N-1]
	ans := 0
	// 谁小结算谁  左边最大值(i-1) i 右边最大值(i+1)
	for L <= R {
		if leftMax <= rightMax {
			ans += max(0, leftMax-height[L])
			leftMax = max(height[L], leftMax)
			L++
		} else {
			ans += max(0, rightMax-height[R])
			rightMax = max(height[R], rightMax)
			R--
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 方法二， 每一次只结算i位置，生成 left 和 right高度辅助数组
