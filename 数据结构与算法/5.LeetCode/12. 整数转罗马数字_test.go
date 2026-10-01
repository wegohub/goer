package leetcode

import (
	"fmt"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	fmt.Println(intToRoman(3749))
	return "Hello World!", nil
}

func intToRoman(num int) string {
	mp := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 0 - 9
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 10 - 90
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 100 - 900
		{"", "M", "MM", "MMM"}, // 1000 - 3000
	}

	sb := strings.Builder{}
	sb.WriteString(mp[3][num/1000%10])
	sb.WriteString(mp[2][num/100%10])
	sb.WriteString(mp[1][num/10%10])
	sb.WriteString(mp[0][num%10])

	return sb.String()
}
