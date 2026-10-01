package class11

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func win(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return int(math.Max(
		float64(f(arr, 0, len(arr)-1)),
		float64(g(arr, 0, len(arr)-1)),
	))
}

func f(arr []int, L, R int) int {
	if L == R {
		return arr[L]
	}

	return int(math.Max(
		float64(arr[L]+g(arr, L+1, R)),
		float64(arr[R]+g(arr, L, R-1)),
	))
}

func g(arr []int, L, R int) int {
	if L == R {
		return 0
	}

	return int(math.Min(
		float64(f(arr, L+1, R)),
		float64(f(arr, L, R-1)),
	))
}
