package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func wordBreak(s string, wordDict []string) bool {
	// return wordBreakNum(s, wordDict) != 0
	set := make(map[string]struct{})
	for _, item := range wordDict {
		set[item] = struct{}{}
	}

	return wordBreakNumDP(s, set) != 0
}

// 获取能分解的方法数(贴纸问题)
func wordBreakNum(s string, wordDict []string) int {
	return process(s, wordDict, 0)
}

// s[index...]这一段字符，能够被分解的方法数，返回
func process(s string, wordDict []string, index int) int {
	if index == len(s) {
		return 1
	}
	// index 没到最后
	// index...end
	ways := 0
	for end := index; end < len(s); end++ {
		pre := s[index : end+1]
		if Containar(pre, wordDict) {
			ways += process(s, wordDict, end+1)
		}
	}
	return ways
}

func Containar(w string, s []string) bool {
	for _, item := range s {
		if item == w {
			return true
		}
	}
	return false
}

func wordBreakNumDP(s string, wordDict map[string]struct{}) int {
	N := len(s)
	dp := make([]int, N+1)
	dp[N] = 1
	for index := N - 1; index >= 0; index-- {
		ways := 0
		for end := index; end < len(s); end++ {
			// hash表可以用前缀树优化
			pre := s[index : end+1]
			if _, ok := wordDict[pre]; ok {
				ways += dp[end+1]
			}
		}
		dp[index] = ways
	}
	return dp[0]
}

type Node struct {
	End   bool
	Nexts []*Node
}

// 前缀树优化版本
func wordBreakNumDPTrie(s string, wordDict []string) int {
	// 单词建前缀树
	root := &Node{
		End:   false,
		Nexts: make([]*Node, 26),
	}
	for _, item := range wordDict {
		node := root
		for i := 0; i < len(item); i++ {
			index := item[i] - 'a'
			if node.Nexts[index] == nil {
				node.Nexts[index] = &Node{Nexts: make([]*Node, 26)}
			}
			node = node.Nexts[index]
		}
		node.End = true
	}

	N := len(s)
	dp := make([]int, N+1)
	dp[N] = 1
	for index := N - 1; index >= 0; index-- {
		cur := root
		for end := index; end < len(s); end++ {
			cur = cur.Nexts[s[end]-'a']
			if cur == nil {
				break
			}
			if cur.End {
				dp[index] += dp[end+1]
			}
		}
	}
	return dp[0]
}
