package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		cand := 0
		// 找到明星，明星只认识自己
		for i := 0; i < n; i++ {
			if knows(cand, i) {
				cand = i
			}
		}
		// 看看明星是不是不认识所有的人
		for i := 0; i < cand; i++ {
			if knows(cand, i) {
				return -1
			}
		}
		// 看看是不是所有人都认识他
		for i := 0; i < n; i++ {
			if !knows(i, cand) {
				return -1
			}
		}
		return cand
	}
}
