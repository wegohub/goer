package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(uniquePaths(7, 4))
	return "Hello World!", nil
}

// m = 7 n = 4
func uniquePaths(m int, n int) int {
	part := n - 1    // 3
	all := m + n - 2 // 9
	o1 := 1
	o2 := 1

	i := part + 1
	j := 1

	// C(9, 3) =  9 * 8 * 7 * 6 * 5 * 4 / 6 * 5 * 4 * 3 * 2 * 1
	for i <= all && j <= all-part {
		fmt.Println(i, j)
		o1 *= i
		o2 *= j
		// 防止溢出
		//g := gcd(o1, o2)
		//o1 /= g
		//o2 /= g
		i++
		j++
	}

	//return o1
	return o1 / o2
}

// 初次调用x和y不能为0， 辗转相除法
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func gcd1(x, y int) int {
	if y == 0 {
		return x
	}
	return gcd1(y, x%y)
}

// 动态规划解
func uniquePathsDP(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	dp := make([][]int, m)
	for index := range dp {
		dp[index] = make([]int, n)
	}
	// 最后一列
	for i := 0; i < m; i++ {
		dp[i][n-1] = 1
	}

	// 最后一行
	for j := 0; j < n; j++ {
		dp[m-1][j] = 1
	}

	// 普遍位置
	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			dp[i][j] = dp[i+1][j] + dp[i][j+1]
		}
	}
	return dp[0][0]
}

// 动态规划空间优化
func uniquePathsDP2(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}

	if m < n {
		dp := make([]int, m)
		for i := 0; i < m; i++ {
			dp[i] = 1
		}
		for j := n - 2; j >= 0; j-- {
			pre := 1
			for i := m - 2; i >= 0; i-- {
				pre = pre + dp[i]
				dp[i] = pre
			}
		}
		return dp[0]
	}

	dp := make([]int, n)
	for j := 0; j < n; j++ {
		dp[j] = 1
	}

	for i := m - 2; i >= 0; i-- {
		pre := 1
		for j := n - 2; j >= 0; j-- {
			pre = pre + dp[j]
			dp[j] = pre
		}
	}
	return dp[0]
}
