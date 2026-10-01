package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// 逆向双指针(最优解) 时间复杂度O(m+n) 空间复杂度O(1)
func merge3(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1)
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			nums1[index-1] = nums1[m-1]
			m--
		} else {
			nums1[index-1] = nums2[n-1]
			n--
		}
		index--
	}
	for m > 0 {
		nums1[index-1] = nums1[m-1]
		m--
		index--
	}
	for n > 0 {
		nums1[index-1] = nums2[n-1]
		n--
		index--
	}
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	index := len(nums1)
	for m > 0 && n > 0 {
		if nums1[m-1] >= nums2[n-1] {
			index--
			m--
			nums1[index] = nums1[m]
		} else {
			index--
			n--
			nums1[index] = nums2[n]
		}
	}
	for m > 0 {
		m--
		index--
		nums1[index] = nums1[m]
	}

	for n > 0 {
		n--
		index--
		nums1[index] = nums2[n]
	}
}

// 归并排序解法 时间复杂度O(m+n) 空间复杂度O(m+n)
func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	help := make([]int, m+n)
	i := 0
	j := 0
	index := 0
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			help[index] = nums1[i]
			i++
			index++
		} else {
			help[index] = nums2[j]
			j++
			index++
		}
	}

	for i < m {
		help[index] = nums1[i]
		i++
		index++
	}

	for j < n {
		help[index] = nums2[j]
		j++
		index++
	}

	for k := 0; k < len(help); k++ {
		nums1[k] = help[k]
	}
}
