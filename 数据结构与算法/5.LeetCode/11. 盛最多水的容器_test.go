package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	return "Hello World!", nil
}

func maxArea(height []int) int {
	L := 0
	R := len(height) - 1

	ans := 0

	// 谁小结算谁, 谁小移动谁
	for L < R {
		ans = max(ans, min(height[L], height[R])*(R-L+1))
		if height[L] > height[R] {
			R--
		} else {
			L++
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
