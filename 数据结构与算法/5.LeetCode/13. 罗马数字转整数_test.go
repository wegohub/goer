package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func romanToInt(s string) int {
	nums := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			nums[i] = 1000
		case 'D':
			nums[i] = 500
		case 'C':
			nums[i] = 100
		case 'L':
			nums[i] = 50
		case 'X':
			nums[i] = 10
		case 'V':
			nums[i] = 5
		case 'I':
			nums[i] = 1
		}
	}

	sum := 0
	// 最后一个位置不填
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			sum -= nums[i]
		} else {
			sum += nums[i]
		}
	}

	return sum + nums[len(nums)-1]
}
