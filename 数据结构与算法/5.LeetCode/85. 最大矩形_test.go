package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	maxArea := 0
	height := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				height[j] = 0
			} else {
				height[j] += 1
			}
		}
		maxArea = Max(maxRecFromBottom(height), maxArea)
	}
	return maxArea
}

func maxRecFromBottom(height []int) int {
	if height == nil || len(height) == 0 {
		return 0
	}
	maxArea := 0
	stack := []int{}
	for i := 0; i < len(height); i++ {
		for len(stack) > 0 && height[i] <= height[stack[len(stack)-1]] {
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			k := -1
			if len(stack) > 0 {
				k = stack[len(stack)-1]
			}
			curArea := (i - k - 1) * height[j]
			maxArea = Max(maxArea, curArea)
		}
		stack = append(stack, i)
	}
	for len(stack) > 0 {
		j := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k := -1
		if len(stack) > 0 {
			k = stack[len(stack)-1]
		}
		curArea := (len(height) - k - 1) * height[j]
		maxArea = Max(maxArea, curArea)
	}
	return maxArea
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
