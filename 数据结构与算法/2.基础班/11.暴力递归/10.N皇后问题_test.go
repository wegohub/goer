package class11

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// @timeout: 3
func main(params map[string]interface{}) (interface{}, error) {
	N := 4

	start := time.Now()
	ans1 := Num(N)
	fmt.Println(fmt.Sprintf("%dms", time.Since(start).Milliseconds()), "====", ans1)

	start = time.Now()
	ans2 := Num2(N)
	fmt.Println(fmt.Sprintf("%dms", time.Since(start).Milliseconds()), "====", ans2)

	return "Hello World!", nil
}

func Num(n int) int {
	// 记录在哪一行那一列放了皇后
	record := make([]int, n)
	return process(0, record, n)
}

func process(i int, record []int, n int) int {
	if i == n {
		return 1
	}

	ans := 0
	for j := 0; j < n; j++ {
		if isValidate(record, i, j) {
			record[i] = j
			ans += process(i+1, record, n)
		}
	}

	return ans
}

func isValidate(record []int, i, j int) bool {
	for k := 0; k < i; k++ {
		// 共列和共对角线
		if record[k] == j || math.Abs(float64(record[k]-j)) == math.Abs(float64(k-i)) {
			return false
		}
	}
	return true
}

func Num2(n int) int {
	if n < 1 || n > 32 {
		return 0
	}

	var limit int32 = -1
	if limit < 32 {
		limit = (1 << n) - 1
	}

	printBinary(int32(1 << n))
	printBinary(limit)

	return proces2(limit, 0, 0, 0)
}

// limit 总限制，就是N
// colLimit 列的限制
// leftLimit 左对角线的限制
// rightLimit 右对角线的限制
// 用二进制位来保存前面皇后的摆放
func proces2(limit, colLimit, leftLimit, rightLimit int32) int {
	if colLimit == limit {
		return 1
	}

	// colLimit | leftLimit | rightLimit 总限制
	// ^(colLimit | leftLimit | rightLimit) 转换为为1的可以放皇后的位置
	pos := limit & ^(colLimit | leftLimit | rightLimit)

	// 最右侧的1
	var mostRighOne int32 = 0
	ans := 0
	for pos != 0 {
		// 提取出最右侧的1
		mostRighOne = pos & (^pos + 1)
		pos = pos - mostRighOne

		ans += proces2(
			limit,
			colLimit|mostRighOne,        // 加上当前列的限制
			(leftLimit|mostRighOne)<<1,  // 加上当前列的限制 << 1位得到左对角线的限制
			(rightLimit|mostRighOne)>>1, // 加上当前列的限制 >> 1位得到右对角线的限制
		)
	}

	return ans
}

func printBinary(num int32) {
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
