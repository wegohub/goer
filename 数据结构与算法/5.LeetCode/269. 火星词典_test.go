package leetcode

import "strings"

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// alienOrder 根据给定的单词列表推断出外星语言的字典序
func alienOrder(words []string) string {
	if words == nil || len(words) == 0 {
		return ""
	}

	// 入度表
	N := len(words)
	indegree := make(map[rune]int)
	for i := 0; i < N; i++ {
		for _, c := range words[i] {
			indegree[c] = 0
		}
	}

	graph := make(map[rune]map[rune]bool)
	for i := 0; i < N-1; i++ {
		cur := []rune(words[i])
		nex := []rune(words[i+1])
		minLen := min(len(cur), len(nex)) // 直接使用 len 函数
		j := 0
		for ; j < minLen; j++ {
			if cur[j] != nex[j] {
				if _, ok := graph[cur[j]]; !ok {
					graph[cur[j]] = make(map[rune]bool)
				}
				if !graph[cur[j]][nex[j]] {
					graph[cur[j]][nex[j]] = true
					indegree[nex[j]]++
				}
				break
			}
		}
		if j < len(cur) && j == len(nex) {
			return ""
		}
	}

	var ans strings.Builder
	q := []rune{}
	for key := range indegree {
		if indegree[key] == 0 {
			q = append(q, key)
		}
	}

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		ans.WriteRune(cur)
		if nexts, ok := graph[cur]; ok {
			for next := range nexts {
				indegree[next]--
				if indegree[next] == 0 {
					q = append(q, next)
				}
			}
		}
	}

	if ans.Len() == len(indegree) {
		return ans.String()
	}
	return ""
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
