package class12

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	ans1 := getHorseMethods(2, 3, 3)
	ans2 := getHorseMethodsDP(2, 3, 3)
	fmt.Println(ans1, ans2)

	return "Hello World!", nil
}

func getHorseMethods(x, y, k int) int {
	return processGetHorseMethods(x, y, 0, 0, k)
}

func processGetHorseMethods(targetX, targetY, currentX, currentY int, rest int) int {
	// 越界的情况
	if currentX < 0 || currentX > 9 || currentY < 0 || currentY > 8 {
		return 0
	}
	// 蹦完了
	if rest == 0 {
		if currentX == targetX && currentY == targetY {
			return 1
		}
		return 0
	}

	ans := processGetHorseMethods(targetX, targetY, currentX-2, currentY+1, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX-1, currentY+2, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX+1, currentY+2, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX+2, currentY+1, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX+2, currentY-1, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX+1, currentY-2, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX-2, currentY-1, rest-1)
	ans += processGetHorseMethods(targetX, targetY, currentX-1, currentY-2, rest-1)

	return ans
}

func getHorseMethodsDP(x, y, k int) int {
	dp := make([][][]int, 10)
	for i := 0; i < 10; i++ {
		tmp := make([][]int, 9)
		for j := 0; j < 9; j++ {
			tmp[j] = make([]int, k+1)
		}
		dp[i] = tmp
	}
	dp[x][y][0] = 1

	for rest := 1; rest <= k; rest++ {
		for currentX := 0; currentX < 10; currentX++ {
			for currentY := 0; currentY < 9; currentY++ {
				ans := getValue(currentX-2, currentY+1, rest-1, dp)
				ans += getValue(currentX-1, currentY+2, rest-1, dp)
				ans += getValue(currentX+1, currentY+2, rest-1, dp)
				ans += getValue(currentX+2, currentY+1, rest-1, dp)
				ans += getValue(currentX+2, currentY-1, rest-1, dp)
				ans += getValue(currentX+1, currentY-2, rest-1, dp)
				ans += getValue(currentX-2, currentY-1, rest-1, dp)
				ans += getValue(currentX-1, currentY-2, rest-1, dp)
				dp[currentX][currentY][rest] = ans
			}
		}
	}
	return dp[0][0][k]
}

func getValue(x, y, k int, dp [][][]int) int {
	if x < 0 || x > 9 || y < 0 || y > 8 {
		return 0
	}
	return dp[x][y][k]
}
