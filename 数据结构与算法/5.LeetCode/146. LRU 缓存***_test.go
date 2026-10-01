package leetcode

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	return "Hello World!", nil
}

// Node 节点定义
type Node struct {
	Key  int
	Val  int
	Last *Node
	Next *Node
}

// 实例化双向链表
func NewDoubleLink() *DoubleLink {
	return &DoubleLink{}
}

// 双链表结构定义
type DoubleLink struct {
	head *Node
	tail *Node
}

// 往尾巴上加入节点
func (obj *DoubleLink) AddNode(node *Node) {
	if node == nil {
		return
	}
	if obj.head == nil {
		obj.head = node
		obj.tail = node
	} else {
		obj.tail.Next = node
		node.Last = obj.tail
		obj.tail = node
	}
}

// 将某一个节点移动到尾巴上(潜台词：上游保证节点一定存在)
func (obj *DoubleLink) MoveNodeToTail(node *Node) {
	// 是尾巴
	if obj.tail == node {
		return
	}
	// 是头
	if obj.head == node {
		obj.head = obj.head.Next
		obj.head.Last = nil
	} else { // 普遍位置
		node.Last.Next = node.Next
		node.Next.Last = node.Last
	}
	node.Last = obj.tail
	node.Next = nil
	obj.tail.Next = node
	obj.tail = node
}

// 头口节点移除
func (obj *DoubleLink) RemoveHead() *Node {
	if obj.head == nil {
		return nil
	}
	ans := obj.head
	// 只有一个节点
	if obj.head == obj.tail {
		obj.head = nil
		obj.tail = nil
	} else {
		obj.head = obj.head.Next
		obj.head.Last = nil
		ans.Next = nil
	}
	return ans
}

// LRU缓存结构
type LRUCache struct {
	mp   map[int]*Node // 哈希表，key: 值；value：节点
	link *DoubleLink   // 节点组成的双向链表
	cap  int           // 容量
}

// 实例化LRU缓存
func Constructor(capacity int) LRUCache {
	if capacity < 1 {
		panic("capacity should be more than 0")
	}
	return LRUCache{
		mp:   map[int]*Node{},
		link: NewDoubleLink(),
		cap:  capacity,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.mp[key]; ok {
		this.link.MoveNodeToTail(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.mp[key]; ok {
		node.Val = value
		this.link.MoveNodeToTail(node)
		return
	}
	// 淘汰缓存
	if len(this.mp) == this.cap {
		uselessNode := this.link.RemoveHead()
		delete(this.mp, uselessNode.Key)
	}
	node := &Node{Key: key, Val: value}
	this.mp[key] = node
	this.link.AddNode(node)
}
