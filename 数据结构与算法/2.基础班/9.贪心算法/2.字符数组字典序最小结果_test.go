package class09

import (
	"fmt"
	"sort"
	"strings"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	strs := []string{"cba", "abd", "dca", "efg", "xyz", "acd"}
	ans := minDict(strs)
	fmt.Println(ans)
	rightAns := right(strs)
	fmt.Println(rightAns)
	fmt.Println(ans == rightAns)
	return "Hello World!", nil
}

// 贪心算法 O(N*LogN)
func minDict(strs []string) string {
	// a.b <= b.a a放前面
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] <= strs[j]+strs[i]
	})

	ans := strings.Builder{}
	for _, str := range strs {
		ans.WriteString(str)
	}
	return ans.String()
}

// 对数器 O(n!)
func right(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	// len >= 2
	ans := make([][]string, 0)
	process(strs, 0, &ans)
	minDict := strings.Join(ans[0], "")
	for i := 1; i < len(ans); i++ {
		cur := strings.Join(ans[i], "")
		if cur < minDict {
			minDict = cur
		}
	}
	return minDict
}

func process(strs []string, index int, ans *[][]string) {
	if index == len(strs) {
		tmp := make([]string, len(strs))
		copy(tmp, strs)
		*ans = append(*ans, tmp)
		return
	}

	for i := index; i < len(strs); i++ {
		// i 位置的字符跟 index位置的字符交换，跑后续的流程
		strs[index], strs[i] = strs[i], strs[index]
		process(strs, index+1, ans)
		// 恢复现场
		strs[index], strs[i] = strs[i], strs[index]
	}
}
