package base

import "fmt"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	// test1()
	test10()
	return "Hello World!", nil
}

func test1() {
	s := make([]int, 8, 16)
	// 输出零值
	fmt.Println(s[7])
	// Error: reflect: slice index out of range
	fmt.Println(s[8])
}

func test2() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Println(len(s1), cap(s1), s1)
	// 这次还没有扩容
	s1 = append(s1, 1, 2)
	s1[0] = 111
	// 超出cap=4，发生扩容 slice header会重新赋值
	s1 = append(s1, 3)
	s1[0] = 222
	fmt.Println(s, s1, len(s1), cap(s1))
}

// 虽然切片是引用传递，但是在方法调用时，传递的会是一个新的 slice header.
// 因此在局部方法 changeSlice 中，虽然对 s1 进行了 append 操作，但这这会在局部方法中这个独立的 slice header 中生效，不会影响到原方法 Test_slice 当中的 s 和 s1 的长度和容量.
func test10() {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("s1 = %p, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	changeSlice(s1)
	fmt.Printf("s1 = %p, len=%d, cap=%d \n", s1, len(s1), cap(s1))

	fmt.Println("输出2: ", s1[2])
}

func changeSlice(s1 []int) {
	s1 = append(s1, 10)
	fmt.Printf("s1 = %p, len=%d, cap=%d \n", s1, len(s1), cap(s1))
	fmt.Println("输出1: ", s1[2])
}
