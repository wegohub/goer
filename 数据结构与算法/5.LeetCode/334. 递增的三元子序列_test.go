package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 最优解(用arr做ends数组)
func increasingTriplet1(arr []int) bool {
	right := 0
	ans := 1
	for i := 1; i < len(arr); i++ {
		L := 0
		R := right
		for L <= R {
			M := (L + R) / 2
			if arr[M] >= arr[i] {
				R = M - 1
			} else {
				L = M + 1
			}
		}
		right = Max(right, L)
		arr[L] = arr[i]
		ans = Max(ans, L+1)
	}
	return ans >= 3
}

func increasingTriplet(arr []int) bool {
	ends := make([]int, len(arr))
	ends[0] = arr[0]
	right := 0
	ans := 1
	for i := 1; i < len(arr); i++ {
		L := 0
		R := right
		for L <= R {
			M := (L + R) / 2
			if ends[M] >= arr[i] {
				R = M - 1
			} else {
				L = M + 1
			}
		}
		right = Max(right, L)
		ends[L] = arr[i]
		ans = Max(ans, L+1)
	}
	return ans >= 3
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
