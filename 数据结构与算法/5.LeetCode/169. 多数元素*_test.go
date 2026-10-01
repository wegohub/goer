package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(majorityElement([]int{1, 2, 3, 4, 5, 6}))
	return "Hello World!", nil
}

// 一次删掉两个数，剩下的就是n/2的数（打靶问题）
func majorityElement(nums []int) int {
	// 靶子
	cand := 0
	// 血量
	HP := 0

	for i := 0; i < len(nums); i++ {
		// 没有靶子
		if HP == 0 {
			HP = 1
			cand = nums[i]
			// 跟靶子的值相同，攒血量
		} else if nums[i] == cand {
			HP++
			// 跟靶子的值不相同，掉血量
		} else {
			HP--
		}
	}
	// 最后靶子上的数字就是答案，题目给的测试用例保证一定有n/2的数，如果没有需遍历判断最后靶子上的数是否大于n/2
	return cand
}
