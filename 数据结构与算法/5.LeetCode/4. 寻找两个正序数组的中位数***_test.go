package leetcode

import (
	"fmt"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(findMedianSortedArrays1([]int{1, 2}, []int{3, 4}))
	return "Hello World!", nil
}

func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	size := len(nums1) + len(nums2)
	isEven := size&1 == 0
	if len(nums1) != 0 && len(nums2) != 0 {
		if isEven {
			// 1 2 3 4
			return float64(findKthNum(nums1, nums2, size/2)+findKthNum(nums1, nums2, size/2+1)) / 2
		} else {
			// 1 2 3 4 5
			return float64(findKthNum(nums1, nums2, size/2+1))
		}
	} else if len(nums1) != 0 {
		if isEven {
			return float64(nums1[size/2]+nums1[size/2-1]) / 2
		} else {
			return float64(nums1[size/2])
		}
	} else if len(nums2) != 0 {
		if isEven {
			return float64(nums2[size/2]+nums2[size/2-1]) / 2
		} else {
			return float64(nums2[size/2])
		}
	} else {
		return 0
	}
}

// 两个不等长的数组中找到第几小(kth)的数
// 长数组：long     17
// 短数组：short    10
// 1. kth <= short:  getUpMedian(arr1, 0, kth-1, arr2, 0, kth-1)
// 2. short < kth <= long: 求整体15小
// 1 2 3 4 5 6 7 8 9 10
// 1 2 3 4 [ 5 6 7 8 9 10 11 12 13 14 15 ] 16 17
// long[kth-short-1] >= short[short-1] 返回：long[kth-short-1]
// getUpMedian(arr1, 0, short-1, arr2, kth-short, kth-1)
// 3. long < kth <= (short + long): 求第23小
//        kth-long-1   short-1
// 1 2 3 4 5 [ 6 7 8 9 10 ]
//                          kth-short-1       long-1
// 1 2 3 4 5 6 7 8 9 10 11 12 [ 13 14 15 16 17 ]
// 有个问题 短数组淘汰了5个，长数组淘汰了12个，一共是17个，然后求第5小，最后得到的是第22小，所以再多淘汰两个
// short[kth-long-1] >= long[long-1] 返回：short[kth-long-1]
// long[kth-short-1] >= short[short-1] 返回：long[kth-short-1]
// getUpMedian(arr1, kth-long, short-1, arr2, kth-short, long-1)

func findKthNum1(arr1 []int, arr2 []int, kth int) int {
	long := arr1
	short := arr2
	if len(arr2) > len(arr1) {
		long = arr2
		short = arr1
	}
	l := len(long)
	s := len(short)

	if kth <= s {
		return getUpMedian(short, 0, kth-1, long, 0, kth-1)
	}

	if kth > s && kth <= l {
		if long[kth-s-1] >= short[s-1] {
			return long[kth-s-1]
		}
		return getUpMedian(short, 0, s-1, long, kth-s, kth-1)
	}

	if short[kth-l-1] >= long[l-1] {
		return short[kth-l-1]
	}
	if long[kth-s-1] >= short[s-1] {
		return long[kth-s-1]
	}
	return getUpMedian(short, kth-l, s-1, long, kth-s, l-1)
}

// 等长的两个数组获取上中位数，函数调用保证
// 1. 长度为偶数
// arr1: 1  2  3  4   mid1
// arr2: 1' 2' 3' 4'  mid2
// 2 == 2' 返回：arr1[mid1]
// 2 > 2': 3 4 1' 2' 排除 => R1 = mid1  L2 = mid2+1
// 2 < 2': 3' 4' 1 2 排除 => R2 = mid2  L1 = mid1+1
// 2. 长度为奇数
// arr1: 1  2  3  4  5   mid1
// arr2: 1' 2' 3' 4' 5'  mid2
// 3 == 3' 返回：arr1[mid1]
// 3 > 3': 3 4 5 1' 2' 排除
// 长度不一样，先验证3'和2: arr2[mid2] > arr1[mid1-1] 返回：arr2[mid2]
// 1 2 和 4' 5'继续二分：R1 = mid1-1  L2 = mid2+1
// 3 < 3': 3' 4' 5' 1 2 排除
// arr1[mid1] > arr2[mid2-1] 返回：arr1[mid1]
// 4 5 1‘ 2’继续二分：L1 = mid1+1  R2 = mid2-1
// 上中点：最后返回L1 和 L2较小的一个
func getUpMedian1(arr1 []int, L1 int, R1 int, arr2 []int, L2 int, R2 int) int {
	for L1 < R1 {
		mid1 := (L1 + R1) / 2
		mid2 := (L2 + R2) / 2
		if arr1[mid1] == arr2[mid2] {
			return arr1[mid1]
		}
		isEven := (R1-L1+1)&1 == 0
		if isEven {
			if arr1[mid1] > arr2[mid2] {
				R1 = mid1
				L2 = mid2 + 1
			} else {
				R2 = mid2
				L1 = mid1 + 1
			}
		} else {
			if arr1[mid1] > arr2[mid2] {
				if arr2[mid2] >= arr1[mid1-1] {
					return arr2[mid2]
				}
				R1 = mid1 - 1
				L2 = mid2 + 1
			} else {
				if arr1[mid1] >= arr2[mid2-1] {
					return arr1[mid1]
				}
				L1 = mid1 + 1
				R2 = mid2 - 1
			}
		}
	}
	return Min(arr1[L1], arr2[L2])
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
