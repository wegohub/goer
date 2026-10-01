package class01

import (
	"fmt"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	printBinary(123)
	fmt.Println(^123)
	var num int32 = -2
	printBinary(num)
	printBinary(num >> 1) // 带符号右移，用符号位补
	// printBinaryUint32(uint32(uint32(num) >> 1))
	return "Hello World!", nil
}

func printBinary(num int32) {
	builder := strings.Builder{}

	for i := 31; i >= 0; i-- {
		if (num & (1 << i)) == 0 {
			builder.WriteString("0")
		} else { // 结果是 1 << i的值
			builder.WriteString("1")
		}
	}

	fmt.Println(builder.String())
}

func printBinaryUint32(num uint32) {
	builder := strings.Builder{}

	for i := 31; i >= 0; i-- {
		if (num & (1 << i)) == 0 {
			builder.WriteString("0")
		} else {
			builder.WriteString("1")
		}
	}

	fmt.Println(builder.String())
}
