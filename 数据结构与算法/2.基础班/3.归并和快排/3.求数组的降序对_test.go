package class03

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func reversePairs(record []int) int {
	return process(record, 0, len(record)-1)
}

func process(arr []int, L, R int) int {
	if L >= R {
		return 0
	}
	M := (L + R) / 2
	return process(arr, L, M) + process(arr, M+1, R) + merge(arr, L, M, R)
}

func merge(arr []int, L, M, R int) int {
	ans := 0
	help := make([]int, R-L+1)
	index := 0
	i := L
	j := M + 1

	for i <= M && j <= R {
		if arr[i] <= arr[j] {
			help[index] = arr[i]
			// 产生逆序对 [0, j-1] 是比arr[i]小的
			ans += (j - 1) - (M + 1) + 1
			i++
		} else {
			help[index] = arr[j]
			j++
		}
		index++
	}

	for i <= M {
		ans += R - (M + 1) + 1
		help[index] = arr[i]
		i++
		index++
	}

	for j <= R {
		help[index] = arr[j]
		j++
		index++
	}
	copy(arr[L:R+1], help)
	return ans
}
