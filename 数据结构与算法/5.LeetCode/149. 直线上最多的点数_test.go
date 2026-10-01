package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func maxPoints(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	if len(points) <= 2 {
		return len(points)
	}

	ans := 0
	// 这条线一定经过i点，i往后有少的个共斜率
	for i := 0; i < len(points); i++ {
		// 分子 分母 次数
		mp := make(map[int]map[int]int)
		// 相同的点
		samePosition := 1
		// 共x轴
		sameX := 0
		// 共y轴
		sameY := 0
		// 哪个斜率压中的点最多，把最多的点的数量，赋值给line
		line := 0
		for j := i + 1; j < len(points); j++ {
			x := points[j][0] - points[i][0]
			y := points[j][1] - points[i][1]
			if x == 0 && y == 0 {
				samePosition++
			} else if x == 0 {
				sameX++
			} else if y == 0 {
				sameY++
			} else { // 有斜率的点
				gcd := Gcd(x, y)
				x /= gcd
				y /= gcd
				// 确保key存在
				if _, ok := mp[x]; !ok {
					mp[x] = make(map[int]int)
				}
				if _, ok := mp[x][y]; !ok {
					mp[x][y] = 0
				}
				// 设置值
				mp[x][y] = mp[x][y] + 1
				line = Max(mp[x][y], line)
			}
		}
		ans = Max(ans, Max(Max(sameX, sameY), line)+samePosition)
	}
	return ans
}

func Gcd(x, y int) int {
	if y != 0 {
		x, y = y, x%y
	}
	return x
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
