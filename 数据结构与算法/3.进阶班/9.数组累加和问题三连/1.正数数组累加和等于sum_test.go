package class10

import "math"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func getMaxLength(arr []int, sum int) int {
	if len(arr) == 0 || sum <= 0 {
		return 0
	}

	// [0, 0]
	L := 0
	R := 0
	winSum := arr[0]
	maxLen := 0
	for R < len(arr) {
		if winSum == sum { // L++
			maxLen = int(math.Max(float64(maxLen), float64(R-L+1)))
			winSum -= arr[L]
			L++
		} else if winSum < sum { // R++
			R++
			if R == len(arr) {
				break
			}
			winSum += arr[R]
		} else { // L++
			winSum -= arr[L]
			L++
		}
	}

	return maxLen
}
