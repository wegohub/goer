package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// intersect finds the intersection of two arrays, including duplicates.
func intersect(nums1, nums2 []int) []int {
	map1 := make(map[int]int)
	for _, num := range nums1 {
		map1[num]++
	}

	map2 := make(map[int]int)
	for _, num := range nums2 {
		map2[num]++
	}

	var list []int
	for key := range map1 {
		if count, found := map2[key]; found {
			n := min(map1[key], count)
			for i := 0; i < n; i++ {
				list = append(list, key)
			}
		}
	}

	return list
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
