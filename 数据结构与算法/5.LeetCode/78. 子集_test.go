package leetcode

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(subsets([]int{1, 2, 3}))
	return "Hello World!", nil
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)
	path := make([]int, 0)
	process(nums, 0, path, &ans)
	return ans
}

func process(nums []int, index int, path []int, ans *[][]int) {
	if index == len(nums) {
		// path拷贝一下
		tmp := make([]int, 0, len(path))
		for _, item := range path {
			tmp = append(tmp, item)
		}
		*ans = append(*ans, tmp)
	} else {
		// 不要index位置的数
		process(nums, index+1, path, ans)
		// 要index位置的数
		path = append(path, nums[index])
		process(nums, index+1, path, ans)
		// 恢复现场
		path = path[0 : len(path)-1]
	}
}
