package class02

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{1, 2, 3, 4, 5}
	L := 1
	R := 3
	fmt.Println("前缀和实现: ", prefixSumArr(arr, L, R))

	fmt.Println("矩阵实现: ", tableSumArr(arr, L, R))

	return "Hello World!", nil
}

// table实现
func tableSumArr(arr []int, L, R int) int {
	n := len(arr)
	if n == 0 || L > R {
		return 0
	}

	table := make([][]int, n)
	for index := range table {
		table[index] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		table[row][row] = arr[row]
		for col := row + 1; col < n; col++ {
			table[row][col] = table[row][col-1] + arr[col]
		}
	}

	return table[L][R]
}

// 前缀和数组
func prefixSumArr(arr []int, L, R int) int {
	n := len(arr)
	if n == 0 || L > R {
		return 0
	}
	// 构建前缀和数组
	help := make([]int, n)
	pre := 0
	for i := 0; i < n; i++ {
		pre += arr[i]
		help[i] = pre
	}
	// 返回结果
	if L == 0 {
		return help[R]
	}
	return help[R] - help[L-1]
}
