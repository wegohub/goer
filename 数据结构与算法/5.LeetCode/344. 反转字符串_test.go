package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func reverseString(s []byte) {
	L := 0
	R := len(s) - 1
	for L < R {
		s[L], s[R] = s[R], s[L]
		L++
		R--
	}
}
