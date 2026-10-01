package class12

import (
	"fmt"
	"math"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	str := "babac"
	arr := []string{"ba", "c", "abcd"}
	ans1 := minStickers(arr, str)
	fmt.Println(ans1)
	return "Hello World!", nil
}

// 如果任务不可能，则返回 -1
func minStickers(stickers []string, target string) int {
	n := len(stickers)
	// 统计贴纸词频
	stickersMap := make([][]int, n)
	for index, sticker := range stickers {
		pin := make([]int, 26)
		for _, item := range sticker {
			pin[item-'a'] += 1
		}
		stickersMap[index] = pin
	}

	dp := make(map[string]int) // 剩余字符串的长度
	dp[""] = 0                 // 空字符返回0张贴纸

	return process(stickersMap, target, dp)
}

func process(stickers [][]int, rest string, dp map[string]int) int {
	if v, ok := dp[rest]; ok {
		return v
	}

	// 递归调用过程
	ans := math.MaxInt // 搞定rest的最好贴纸数量
	n := len(stickers) // 贴纸的种数
	restMap := make([]int, 26)
	for _, item := range rest {
		restMap[item-'a'] += 1
	}

	// stickers 搞定 restMap
	// 枚举第一张贴纸是谁
	for i := 0; i < n; i++ {
		// 贴纸中不含有rest中的第一个字符
		// 没有这一行代码回跑不完
		if stickers[i][rest[0]-'a'] == 0 {
			continue
		}

		sb := strings.Builder{}

		// 枚举所有剩余字符
		for j := 0; j < 26; j++ {
			// 当前位置的剩余字符>0
			if restMap[j] > 0 {
				// i号位置的贴纸能搞定多少个rest字符，将剩下的字符加入sb
				for k := 0; k < int(math.Max(0, float64(restMap[j]-stickers[i][j]))); k++ {
					sb.WriteString(string(rune('a' + j)))
				}
			}
		}

		// 剩余字符调后续的过程
		s := sb.String()
		tmp := process(stickers, s, dp)
		if tmp != -1 {
			ans = int(math.Min(float64(ans), float64(1+tmp)))
		}
	}

	if ans == math.MaxInt {
		ans = -1
	}

	dp[rest] = ans
	return ans
}
