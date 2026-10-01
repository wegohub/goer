package class01

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	arr := []int{5, 21, 4, 6, 1, 3, 0}
	insertSort(arr)
	fmt.Println(arr)
	return "Hello World!", nil
}

func insertSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	// 0 - 0 上有序
	// 0 - 1 上有序
	// ...
	// 0 ～ n-1 上有序
	for i := 0; i < n-1; i++ {
		for j := i + 1; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	// 或者
	// for end := 1; end < n; end++ {
	//    for pre := end-1; pre >= 0 && nums[pre] > nums[pre+1];pre--{
	//        nums[pre], nums[pre+1] = nums[pre+1], nums[pre]
	//    }
	//}
}
