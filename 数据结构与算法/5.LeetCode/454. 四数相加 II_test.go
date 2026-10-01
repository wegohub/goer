package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 分治思想
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// key: 累加和 value: -sum的种树
	mp := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			sum := nums1[i] + nums2[j]
			if v, ok := mp[sum]; ok {
				mp[sum] = v + 1
			} else {
				mp[sum] = 1
			}
		}
	}

	ans := 0
	for i := 0; i < len(nums3); i++ {
		for j := 0; j < len(nums4); j++ {
			sum := nums3[i] + nums4[j]
			if v, ok := mp[-sum]; ok {
				ans += v
			}
		}
	}

	return ans
}

// 这个只是coding更优雅
func fourSumCount2(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}
