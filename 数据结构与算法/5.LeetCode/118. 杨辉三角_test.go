package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(generate(3))
	return "Hello World!", nil
}

func generate(numRows int) [][]int {
	ans := make([][]int, 0)
	var lastRow []int
	for i := 1; i <= numRows; i++ {
		var tmp []int
		tmp = append(tmp, 1)
		if i == 1 {
			ans = append(ans, tmp)
			continue
		}
		for j := 1; j < len(lastRow); j++ {
			tmp = append(tmp, lastRow[j-1]+lastRow[j])
		}
		tmp = append(tmp, 1)
		lastRow = tmp
		ans = append(ans, tmp)
	}

	return ans
}

func generate2(numRows int) [][]int {
	ans := make([][]int, numRows)
	ans[0] = []int{1}
	for i := 1; i < numRows; i++ {
		ans[i] = append(ans[i], 1)
		for j := 1; j <= i-1; j++ {
			ans[i] = append(ans[i], ans[i-1][j]+ans[i-1][j-1])
		}
		ans[i] = append(ans[i], 1)
	}
	return ans
}
