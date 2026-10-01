package class12

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{1, 5, 10, 50, 100}
	aim := 100
	//ans1 := Ways([]int{1,5,10,50,100}, 100)
	ans2 := WaysCache(arr, aim)
	ans3 := WaysDP(arr, aim)
	ans4 := WaysDP2(arr, aim)
	//fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
	fmt.Println(ans4)

	return "Hello World!", nil
}

func Ways(arr []int, aim int) int {
	return process(arr, 0, aim)
}

func process(arr []int, index int, rest int) int {
	if index == len(arr) { // 没有货币可以选择了
		if rest == 0 {
			return 1
		}
		return 0
	}

	ways := 0
	for zhang := 0; arr[index]*zhang <= rest; zhang++ {
		ways += process(arr, index+1, rest-(arr[index]*zhang))
	}

	return ways
}

func WaysDP(arr []int, aim int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest < aim+1; rest++ {
			ways := 0
			for zhang := 0; arr[index]*zhang <= rest; zhang++ {
				ways += dp[index+1][rest-(arr[index]*zhang)]
			}
			dp[index][rest] = ways
		}
	}

	return dp[0][aim]
}

func WaysDP2(arr []int, aim int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, aim+1)
	}
	dp[n][0] = 1
	for index := n - 1; index >= 0; index-- {
		for rest := 0; rest < aim+1; rest++ {
			dp[index][rest] = dp[index+1][rest]
			if rest-arr[index] >= 0 {
				dp[index][rest] += dp[index][rest-arr[index]]
			}
		}
	}

	return dp[0][aim]
}

func WaysCache(arr []int, aim int) int {
	n := len(arr)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		tmp := make([]int, aim+1)
		for j := 0; j < aim+1; j++ {
			tmp[j] = -1
		}
		dp[i] = tmp
	}
	return processCache(arr, 0, aim, dp)
}

func processCache(arr []int, index int, rest int, dp [][]int) int {
	if dp[index][rest] != -1 {
		return dp[index][rest]
	}
	if index == len(arr) { // 没有货币可以选择了
		if rest == 0 {
			dp[index][rest] = 1
			return 1
		}
		dp[index][rest] = 0
		return 0
	}

	ways := 0
	for zhang := 0; arr[index]*zhang <= rest; zhang++ {
		ways += processCache(arr, index+1, rest-(arr[index]*zhang), dp)
	}

	dp[index][rest] = ways
	return ways
}
