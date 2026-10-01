package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

func NewTrie() *Trie {
	return &Trie{root: &Node{
		Pass:  0,
		End:   0,
		Nexts: make(map[rune]*Node),
	}}
}

// Node 前缀树节点
type Node struct {
	Pass  int            // 字符经过次数
	End   int            // 字符结尾次数
	Nexts map[rune]*Node // 子字符
}

// Trie 前缀树实现
type Trie struct {
	root *Node
}

// Insert 插入字符
func (obj *Trie) Insert(str string) {
	if str == "" {
		return
	}
	obj.root.Pass++
	node := obj.root
	for _, item := range str {
		if cur, ok := node.Nexts[rune(item)]; ok {
			cur.Pass++
			node = cur
		} else {
			newNode := &Node{Pass: 1, End: 0, Nexts: make(map[rune]*Node)}
			node.Nexts[rune(item)] = newNode
			node = newNode
		}
	}
	node.End++
}

// Search 搜索str加入过几次
func (obj *Trie) Search(str string) int {
	if str == "" {
		return 0
	}
	node := obj.root
	for _, item := range str {
		if cur, ok := node.Nexts[rune(item)]; !ok {
			return 0
		} else {
			node = cur
		}
	}
	return node.End
}

// Prefix 有多少个字符是以str开头的
func (obj *Trie) Prefix(str string) int {
	if str == "" {
		return 0
	}
	node := obj.root
	for _, item := range str {
		if cur, ok := node.Nexts[rune(item)]; !ok {
			return 0
		} else {
			node = cur
		}
	}
	return node.Pass
}

// Delete 删除字符
func (obj *Trie) Delete(str string) {
	if str == "" {
		return
	}

	// 存在str才删除
	if obj.Search(str) > 0 {
		node := obj.root
		node.Pass--
		for _, item := range str {
			cur := node.Nexts[rune(item)]
			cur.Pass--
			if cur.Pass == 0 { // 都没有下级了
				delete(node.Nexts, rune(item)) // 将所有的下级释放掉
				return
			}
			node = cur
		}
		node.End--
	}
}
