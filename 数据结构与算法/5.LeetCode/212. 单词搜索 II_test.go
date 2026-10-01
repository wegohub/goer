package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// TrieNode 定义前缀树节点结构
type TrieNode struct {
	nexts [26]*TrieNode
	pass  int
	end   int
}

// NewTrieNode 创建一个新的前缀树节点
func NewTrieNode() *TrieNode {
	return &TrieNode{
		nexts: [26]*TrieNode{},
		pass:  0,
		end:   0,
	}
}

// findWords 主函数，找到所有在board中出现的words
func findWords(board [][]byte, words []string) []string {
	head := NewTrieNode() // 前缀树最顶端的头
	set := make(map[string]bool)
	for _, word := range words {
		if !set[word] {
			fillWord(head, word)
			set[word] = true
		}
	}

	// 答案
	res := []string{}
	// 沿途走过的字符，收集起来，存在path里
	path := []byte{}
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {
			// 枚举在board中的所有位置
			// 每一个位置出发的情况下，答案都收集
			process(board, row, col, &path, head, &res)
		}
	}
	return res
}

// fillWord 将word填充到前缀树中
func fillWord(node *TrieNode, word string) {
	node.pass++
	chs := []byte(word)
	for i := 0; i < len(chs); i++ {
		index := chs[i] - 'a'
		if node.nexts[index] == nil {
			node.nexts[index] = NewTrieNode()
		}
		node = node.nexts[index]
		node.pass++
	}
	node.end++
}

// process 从board[row][col]位置的字符出发，寻找words
func process(board [][]byte, row int, col int, path *[]byte, cur *TrieNode, res *[]string) int {
	cha := board[row][col]
	if cha == 0 { // 这个row col位置是之前走过的位置
		return 0
	}

	index := cha - 'a'
	// 如果没路，或者这条路上最终的字符串之前加入过结果里
	if cur.nexts[index] == nil || cur.nexts[index].pass == 0 {
		return 0
	}

	// 没有走回头路且能登上去
	cur = cur.nexts[index]
	*path = append(*path, cha) // 当前位置的字符加到路径里去
	fix := 0                   // 从row和col位置出发，后续一共搞定了多少答案

	// 当我来到row col位置，如果决定不往后走了。是不是已经搞定了某个字符串了
	if cur.end > 0 {
		*res = append(*res, string(*path))
		cur.end--
		fix++
	}

	// 往上、下、左、右，四个方向尝试
	board[row][col] = 0
	if row > 0 {
		fix += process(board, row-1, col, path, cur, res)
	}
	if row < len(board)-1 {
		fix += process(board, row+1, col, path, cur, res)
	}
	if col > 0 {
		fix += process(board, row, col-1, path, cur, res)
	}
	if col < len(board[0])-1 {
		fix += process(board, row, col+1, path, cur, res)
	}
	board[row][col] = cha
	*path = (*path)[:len(*path)-1]
	cur.pass -= fix
	return fix
}
