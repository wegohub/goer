package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func hammingWeight(n int) int {
	bits := 0
	for n != 0 {
		bits++
		// 提取最右侧的1
		rightOne := n & (^n + 1)
		// 将最右侧的1去掉
		n = n ^ rightOne
	}
	return bits
}

func hammingWeight1(n int) int {
	ans := 0
	for n != 0 {
		ans++
		n ^= n & (^n + 1)
	}
	return ans
}

func hammingWeight2(n int) int {
	n = (n & 0x55555555) + ((n >> 1) & 0x55555555)
	n = (n & 0x33333333) + ((n >> 2) & 0x33333333)
	n = (n & 0x0f0f0f0f) + ((n >> 4) & 0x0f0f0f0f)
	n = (n & 0x00ff00ff) + ((n >> 8) & 0x00ff00ff)
	n = (n & 0x0000ffff) + ((n >> 16) & 0x0000ffff)
	return n
}
