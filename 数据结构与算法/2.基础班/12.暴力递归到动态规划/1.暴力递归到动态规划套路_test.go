package class12

import (
	"fmt"
	"time"
)

// @timeout: 10
func main(params map[string]interface{}) (interface{}, error) {
	start := time.Now()
	fmt.Println(Fibonacci(30))
	fmt.Println("Fibonacci: ", fmt.Sprintf("%dms", time.Since(start).Milliseconds()))

	start = time.Now()
	fmt.Println(FibonacciDP(30))
	fmt.Println("FibonacciDP: ", fmt.Sprintf("%dms", time.Since(start).Milliseconds()))
	return "Hello World!", nil
}

func Fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func FibonacciDP(n int) int {
	if n <= 2 {
		return 1
	}

	//nMulOne := 1
	//nMulTow := 1
	n1 := 1
	n2 := 1
	// ans := 0
	for i := 3; i <= n; i++ {
		tmp := n1
		n1 = n1 + n2
		n2 = tmp
		//ans = nMulOne + nMulTow
		//nMulTow = nMulOne
		//nMulOne = ans
	}
	return n1
}
