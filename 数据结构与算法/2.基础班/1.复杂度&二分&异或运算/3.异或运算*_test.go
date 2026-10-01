package class01

import (
	"fmt"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	a := 1
	b := 2
	Swap(&a, &a)
	fmt.Println("a: ", a, ", b: ", b)

	arr := []int{1, 2}
	SwapArr(arr, 0, 1)
	fmt.Println(arr)

	fmt.Println("===============")
	fmt.Println(oddNumber([]int{1, 1, 2, 3, 3, 4, 4, 5, 6, 5, 6}))

	fmt.Println("==================")
	fmt.Println(oddNumber2([]int{1, 1, 2, 2, 3, 3, 3, 4, 4, 6, 5, 6}))

	fmt.Println("==================")
	fmt.Println(oneCount(6))

	printBinary((^6 + 1))
	printBinary(6 & (^6 + 1))

	return "Hello World!", nil
}

// Swap 交换a和b的值(a != b)
func Swap(a *int, b *int) {
	if a == b {
		return
	}
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

// SwapArr 交换数组中的两个下标(a != b)，因为 N ^ N = 0
func SwapArr(arr []int, a, b int) {
	if a == b {
		return
	}
	arr[a] = arr[a] ^ arr[b]
	arr[b] = arr[a] ^ arr[b]
	arr[a] = arr[a] ^ arr[b]
}

// 一个数组中有一种数出现奇数次，其他数都出现偶数次，怎么找到并打印这种数？
func oddNumber(arr []int) int {
	eor := 0
	for _, item := range arr {
		eor = eor ^ item
	}
	return eor
}

// 一个数组中有两种种数出现奇数次，其他数都出现偶数次，怎么找到并打印这种数？
func oddNumber2(arr []int) [2]int {
	eor := 0
	for _, item := range arr {
		eor = eor ^ item
	}
	// 提取最右侧的1, eor 是 a ^ b 的结果，所以 最右侧肯定是 1 ^ 0 || 0 ^ 1
	rightOne := eor & (^eor + 1)

	eor2 := 0
	for _, item := range arr {
		if item&rightOne != 0 {
			eor2 = eor2 ^ item
		}
	}

	return [2]int{eor2, eor2 ^ eor}
}

// 给你一个数n，计算出二进制中有多少个1？
func oneCount(n int) int {
	count := 0

	for n != 0 {
		rightOne := n & (^n + 1)
		count++
		n = n ^ rightOne
	}

	return count
}

func printBinary(n int32) {
	builder := strings.Builder{}
	for i := 31; i >= 0; i-- {
		if n&(1<<i) != 0 {
			builder.WriteString("1")
		} else {
			builder.WriteString("0")
		}
	}

	fmt.Println(builder.String())

}
