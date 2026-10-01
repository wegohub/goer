package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {

	return "Hello World!", nil
}

func singleNumber(nums []int) int {
	eor := 0
	for _, item := range nums {
		eor = eor ^ item
	}
	return eor
}
