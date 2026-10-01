package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	heapSort(nums)

	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			return true
		}
	}

	return false
}

func heapSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	heapSize := len(arr)

	// O(n) 建堆
	for i := len(arr) - 1; i >= 0; i-- {
		heapify(arr, i, heapSize)
	}

	// 堆排序
	heapSize--
	arr[0], arr[heapSize] = arr[heapSize], arr[0]
	for heapSize > 0 {
		heapify(arr, 0, heapSize)
		heapSize--
		arr[0], arr[heapSize] = arr[heapSize], arr[0]
	}
}

func heapify(arr []int, index, heapSize int) {
	left := 2*index + 1
	for left < heapSize {
		largest := left
		if left+1 < heapSize && arr[left+1] > arr[left] {
			largest = left + 1
		}
		if arr[index] > arr[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		arr[largest], arr[index] = arr[index], arr[largest]
		index = largest
		left = 2*index + 1
	}
}
