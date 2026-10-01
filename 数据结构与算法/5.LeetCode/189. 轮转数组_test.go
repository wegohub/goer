package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func rotate(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	// 得到第几个元素
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 方法一：使用额外的数组
// 这是最直观的方法，创建一个与原数组相同大小的新数组，然后将元素按轮转后的位置放入新数组中。
func rotate1(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n // 处理k大于数组长度的情况
	rotated := make([]int, n)
	for i := 0; i < n; i++ {
		rotated[(i+k)%n] = nums[i]
	}
	copy(nums, rotated)
}

// 方法二：反转数组
// 这种方法不需要额外的空间，通过三次反转数组来实现轮转。首先反转整个数组，然后分别反转前 k 个元素和后 n-k 个元素。
func rotate2(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
	reverse2(nums, 0, n-1)
	reverse2(nums, 0, k-1)
	reverse2(nums, k, n-1)
}

func reverse2(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// 方法三：循环替换
// 这种方法通过直接将每个元素放到它轮转后的位置，使用一个变量来记录已经处理过的元素数量，直到所有元素都被处理。
func rotate3(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
	count := 0
	for start := 0; count < n; start++ {
		current := start
		prev := nums[start]
		for {
			next := (current + k) % n
			nums[next], prev = prev, nums[next]
			current = next
			count++
			if start == current {
				break
			}
		}
	}
}

// 方法四：使用环状替换
// 这种方法与方法三类似，但是通过一个额外的变量来避免重复替换。
func rotate4(nums []int, k int) {
	n := len(nums)
	if n == 0 {
		return
	}
	k %= n
	for start, count := 0, 0; count < n; start++ {
		current, prev := start, nums[start]
		for ok := true; ok; ok = current != start {
			next := (current + k) % n
			nums[next], prev, current = prev, nums[next], next
			count++
		}
	}
}
