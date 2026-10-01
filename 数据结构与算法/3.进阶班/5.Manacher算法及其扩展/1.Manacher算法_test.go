package class05

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	s := "abbaa"
	fmt.Println(Manacher(s))
	return "Hello World!", nil
}

func Manacher(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 将原始字符串处理成manacher串, "12132" -> "#1#2#1#3#2#"
	str := manacherString(s)
	// 回文半径的大小(结果就是pArr-1), 如果记录回文直径就是 pArr / 2
	pArr := make([]int, len(str))
	// 最长回文半径
	C := -1
	// 最长回文半径对应的右边界+1(第一个失败的位置)
	R := -1
	// 结果
	max := math.MinInt

	for i := 0; i < len(str); i++ {

		// i在R外，至少能扩1个字符 i >= R
		pArr[i] = 1
		// i在R内
		if i < R {
			// i'在L...R内 和 i'在L...R外。 2 * C - i 就是i的对称点i'
			pArr[i] = int(math.Min(float64(pArr[2*C-i]), float64(R-i)))
		}

		// 向左右扩，扩的位置不能越界
		// i在R外 和 i'压在L或者R上；兼容i'在L...R内 和 i'在L...R外两种情况，进循环就会出来
		for i+pArr[i] < len(str) && i-pArr[i] > -1 {
			if str[i+pArr[i]] == str[i-pArr[i]] {
				pArr[i]++
			} else {
				break
			}
		}

		// 更新最长回文半径
		if i+pArr[i] > R {
			R = i + pArr[i]
			C = i
		}
		// 更新答案
		max = int(math.Max(float64(max), float64(pArr[i])))
	}

	return max - 1
}

func manacherString(s string) string {
	sb := strings.Builder{}
	sb.WriteString("#")
	for _, item := range s {
		sb.WriteString(string(item) + "#")
	}
	return sb.String()
}
