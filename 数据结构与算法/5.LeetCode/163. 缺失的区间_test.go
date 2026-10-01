package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	var ans [][]int
	for _, num := range nums {
		if num > lower && num-1 >= lower {
			ans = append(ans, []int{lower, num - 1})
		}
		if num == upper {
			return ans
		}
		lower = num + 1
	}
	if lower <= upper {
		ans = append(ans, []int{lower, upper})
	}

	return ans
}
