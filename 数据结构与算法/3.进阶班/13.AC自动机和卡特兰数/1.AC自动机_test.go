package class14

import (
	"container/list"
	"fmt"
	"unicode/utf8"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ac := NewACAutomation()
	ac.Insert("愿意")
	ac.Insert("不愿意")
	ac.Build()

	contains := ac.ContainWords("我是真的不愿意去吃饭")
	for _, item := range contains {
		fmt.Println(fmt.Sprintf("敏感词: %s, 开始位置: %d, 结束位置: %d", item.Value, item.Start, item.End))
	}
	return "Hello World!", nil
}

// 实例化AC自动机
func NewACAutomation() *ACAutomation {
	return &ACAutomation{root: NewNode()}
}

// 实例化实例化AC自动机节点
func NewNode() *Node {
	return &Node{
		IsEnd: false,
		End:   "",
		Fail:  nil,
		Nexts: make(map[rune]*Node),
	}
}

// Node AC自动机节点
type Node struct {
	// 是否是某个字符串结尾
	IsEnd bool
	// IsEnd为true是的某个字符
	End string
	// Fail指针
	Fail  *Node
	Nexts map[rune]*Node
}

// ACAutomation AC自动机实现
type ACAutomation struct {
	root *Node
}

// Insert 插入词
func (obj *ACAutomation) Insert(str string) {
	if str == "" {
		return
	}
	node := obj.root
	for _, item := range str {
		// 有item方向上的路就跳到item方向上的路继续走
		if cur, ok := node.Nexts[rune(item)]; ok {
			node = cur
		} else {
			// 没有item方向上的路就新建一条item方向上的路
			newNode := NewNode()
			node.Nexts[rune(item)] = newNode
			node = newNode
		}
	}
	// 结束的节点设置上这条路的单词
	node.End = str
	node.IsEnd = true
}

// Build 构建Fail指针
// 宽地优先遍历
func (obj *ACAutomation) Build() {
	queue := list.New()
	// 将根节点加入队列
	queue.PushBack(obj.root)
	// 队列不为空就继续
	for queue.Len() > 0 {
		// 弹出一个父级节点，设置子级的fail指针
		cur := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())

		// 遍历当前节点所有的路
		for char, node := range cur.Nexts {
			// 现将当前子级节点的fail指针设置为root
			node.Fail = obj.root
			// 当前父级节点的fail指针
			curFail := cur.Fail
			for curFail != nil {
				// 当前父级节点的fail指针指向的节点有char方向上的路
				// 将当前子级的fail指针指向它
				if failNode, ok := curFail.Nexts[char]; ok {
					node.Fail = failNode
					break
				}
				// 否则蹦到下一个节点
				curFail = curFail.Fail
			}
			// 将下级节点入队
			queue.PushBack(node)
		}
	}
}

type Word struct {
	Value string // 敏感词
	Start int    // 开始位置
	End   int    // 结束位置
}

// ContainWords 判断文本中是否包含关键词
func (obj *ACAutomation) ContainWords(content string) []*Word {
	ans := make([]*Word, 0)
	if content == "" {
		return ans
	}
	cur := obj.root

	// 遍历文本的每一个字符
	index := 0
	for _, item := range content {
		// 如果当前字符在这条路上没配出来，就随着fail方向走向下条路径
		// 如果当前cur节点，没有path的路，就通过fail，跳到别的前缀上去
		for cur.Nexts[item] == nil && cur != obj.root {
			cur = cur.Fail
		}

		// 1. 现在来到的路径，是可以继续匹配的
		// 2. 现在来到的节点，已经是头了
		if cur.Nexts[item] != nil {
			cur = cur.Nexts[item]
		} else {
			cur = obj.root
		}
		// 从当前位置出发，沿着fail指针收集敏感词
		follow := cur
		for follow != obj.root {
			if follow.IsEnd {
				ans = append(ans, &Word{
					Value: follow.End,
					Start: index - utf8.RuneCountInString(follow.End) + 1,
					End:   index,
				})
			}
			follow = follow.Fail
		}
		index++
	}

	return ans
}
