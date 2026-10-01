package class11

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

// @timeout: 5
func main(params map[string]interface{}) (interface{}, error) {
	ch := NewConsistentHash(3, nil)
	ch.Add("192.168.0.1", "192.168.0.2", "192.168.0.3")

	fmt.Println(ch.Get("key1"))
	fmt.Println(ch.Get("key2"))
	fmt.Println(ch.Get("key3"))
	return "Hello World!", nil
}

// HashFunc 定义哈希函数
type HashFunc func(data []byte) uint32

// ConsistentHash 一致性哈希结构体
type ConsistentHash struct {
	hashFunc HashFunc       // 哈希函数
	replicas int            // 虚拟节点倍数
	keys     []int          // 哈希环
	hashMap  map[int]string // 虚拟节点与真实节点的映射
}

// NewConsistentHash 创建一致性哈希实例
func NewConsistentHash(replicas int, fn HashFunc) *ConsistentHash {
	c := &ConsistentHash{
		replicas: replicas,
		hashFunc: fn,
		hashMap:  make(map[int]string),
	}
	if c.hashFunc == nil {
		c.hashFunc = crc32.ChecksumIEEE
	}
	return c
}

// Add 添加真实节点
func (c *ConsistentHash) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < c.replicas; i++ {
			hash := int(c.hashFunc([]byte(strconv.Itoa(i) + node)))
			c.keys = append(c.keys, hash)
			c.hashMap[hash] = node
		}
	}
	sort.Ints(c.keys)
}

// Get 获取节点
func (c *ConsistentHash) Get(key string) string {
	if len(c.keys) == 0 {
		return ""
	}

	hash := int(c.hashFunc([]byte(key)))
	// 二分查找合适的虚拟节点
	idx := sort.Search(len(c.keys), func(i int) bool {
		return c.keys[i] >= hash
	})

	// 如果 idx == len(c.keys)，则应选择 c.keys[0]，因为这是哈希环
	if idx == len(c.keys) {
		idx = 0
	}

	return c.hashMap[c.keys[idx]]
}
